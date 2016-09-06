package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/blobs"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
	"github.com/random-j-farmer/zkill-mirror/internal/mapdata"
	"github.com/random-j-farmer/zkill-mirror/internal/zkb"
)

const defaultLimit = 100
const maxLimit = 1000

type apiQuery struct {
	KillID        uint64
	CharacterID   uint64
	CorporationID uint64
	AllianceID    uint64
	SystemID      uint64
	RegionID      uint64
	Activity      time.Duration
	Limit         int

	accept string // json or html
}

func apiHandler(w http.ResponseWriter, r *http.Request, url string) {

	if config.Debug() {
		dump, _ := httputil.DumpRequest(r, false)
		logRequestf(r, "request header: %s", dump)
	}

	q, err := unmarshalQuery(url)
	if err != nil {
		apiError(w, r, err)
		return
	}

	// XXX: it's dirty. i think i like it
	if strings.Contains(r.Header.Get("accept"), "html") {
		q.accept = "html"
	} else {
		q.accept = "json"
	}

	switch {
	case q.KillID > 0:
		err = apiByKillID(w, r, q)

	case q.CharacterID > 0:
		err = apiByCharacterID(w, r, q)

	case q.CorporationID > 0:
		err = apiByCorporationID(w, r, q)

	case q.AllianceID > 0:
		err = apiByAllianceID(w, r, q)

	case q.SystemID > 0:
		err = apiBySystemID(w, r, q)

	case q.RegionID > 0:
		err = apiByRegionID(w, r, q)

	case q.Activity > 0:
		err = apiActivity(w, r, q)

	default:
		logRequestf(r, "hmmmm ... maybe you would like all the newest kills?")
		err = apiNewest(w, r, q)
	}

	if err != nil {
		apiError(w, r, err)
	}
}

func apiByKillID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	ref, err := db.ByKillID(q.KillID)
	if err != nil {
		return errors.Wrap(err, "db.ByKillID")
	}

	infos, err := retrieveKillmails([]bobstore.Ref{ref})
	if err != nil {
		return err
	}

	return executeTemplate(w, r, templateName(q, "detail"), infos[0])
}

func apiByCharacterID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByCharacterID(q.CharacterID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByCharacterID")
	}

	logRequestf(r, "apiByCharacterID: %d results", len(refs))
	return apiWriteResponse(w, r, q, refs)
}

func apiByCorporationID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByCorporationID(q.CorporationID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByCorporationID")
	}

	logRequestf(r, "apiByCorporationID: %d results", len(refs))
	return apiWriteResponse(w, r, q, refs)
}

func apiByAllianceID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByAllianceID(q.AllianceID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByAllianceID")
	}

	logRequestf(r, "apiByAllianceID: %d results", len(refs))
	return apiWriteResponse(w, r, q, refs)
}

func apiBySystemID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.BySystemID(q.SystemID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.BySystemID")
	}

	logRequestf(r, "apiBySystemID: %d results", len(refs))
	return apiWriteResponse(w, r, q, refs)
}

func apiByRegionID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByRegionID(q.RegionID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByRegionID")
	}

	logRequestf(r, "apiByRegionID: %d results", len(refs))
	return apiWriteResponse(w, r, q, refs)
}

func apiActivity(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	stats, err := db.Activity(q.Activity)
	if err != nil {
		return errors.Wrap(err, "db.Activity")
	}

	dot := struct {
		Time         string
		SolarSystems []*db.SystemStat
	}{Time: time.Now().Format(db.EveTimeFormat), SolarSystems: stats}

	return executeTemplate(w, r, templateName(q, "activity"), &dot)
}

func templateName(q *apiQuery, n string) string {
	return n + "." + q.accept
}

func apiNewest(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.Newest(q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.Newest")
	}

	logRequestf(r, "apiNewest: %d results", len(refs))
	return apiWriteResponse(w, r, q, refs)
}

func apiWriteResponse(w http.ResponseWriter, r *http.Request, q *apiQuery, refs []bobstore.Ref) error {
	infos, err := retrieveKillmails(refs)
	if err != nil {
		return err
	}

	dot := struct {
		Time      string
		Killmails []*killmailInfo
	}{Time: time.Now().Format(db.EveTimeFormat), Killmails: infos}

	return executeTemplate(w, r, templateName(q, "killmails"), &dot)
}

type killmailInfo struct {
	*zkb.Killmail
	Security        float32
	VictimSummary   string
	AttackerSummary string
}

func retrieveKillmails(refs []bobstore.Ref) ([]*killmailInfo, error) {
	kms := make([]*killmailInfo, len(refs))
	for i, ref := range refs {
		b, err := blobs.DB.Read(ref)
		if err != nil {
			return nil, errors.Wrapf(err, "bobstore.Read %s", ref)
		}

		parsed, err := zkb.Parse(b, ref)
		if err != nil {
			return nil, errors.Wrapf(err, "zkb.Parse %s", ref)
		}

		si := mapdata.SolarSystemByID(parsed.SolarSystemID)
		kms[i] = &killmailInfo{
			Killmail:        parsed,
			Security:        si.Security,
			VictimSummary:   victimSummary(parsed),
			AttackerSummary: attackerSummary(parsed),
		}
	}
	return kms, nil
}

func victimSummary(km *zkb.Killmail) string {
	allOrCorp := km.Victim.AllianceName
	if allOrCorp == "" {
		allOrCorp = km.Victim.CorporationName
	}
	return fmt.Sprintf("%s (%s)", km.Victim.CharacterName, allOrCorp)
}

func attackerSummary(km *zkb.Killmail) string {
	var att *zkb.Attacker
	for _, a := range km.Attackers {
		if a.FinalBlow != 0 {
			att = &a
		}
	}
	if att == nil {
		att = &km.Attackers[0]
	}
	allOrCorp := att.AllianceName
	if allOrCorp == "" {
		allOrCorp = att.CorporationName
	}
	final := fmt.Sprintf("%s (%s)", att.CharacterName, allOrCorp)
	if len(km.Attackers) > 1 {
		return fmt.Sprintf("%s +%d", final, len(km.Attackers))
	}

	return final
}

func apiError(w http.ResponseWriter, r *http.Request, err error) error {
	logRequestf(r, "error: %v", err)
	http.Error(w, fmt.Sprintf("error: %v", err), 500)
	return err
}

func unmarshalQuery(url string) (*apiQuery, error) {
	parts := strings.Split(url, "/")

	// hangle dangling key (no value)
	if len(parts)&1 == 1 {
		if parts[len(parts)-1] == "" {
			parts = parts[:len(parts)-1]
		} else {
			parts = append(parts, "")
		}
	}

	q := &apiQuery{Limit: defaultLimit}

	var err error

	for i := 0; i < len(parts); i = i + 2 {
		switch parts[i] {
		case "killID":
			q.KillID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "characterID":
			q.CharacterID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "corporationID":
			q.CorporationID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "allianceID":
			q.AllianceID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "solarSystemID":
			q.SystemID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "regionID":
			q.RegionID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "activity":
			q.Activity, err = time.ParseDuration(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "time.ParseDuration %s", parts[i])
			}
		case "limit":
			q.Limit, err = strconv.Atoi(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
			if q.Limit > maxLimit {
				log.Printf("limit over maxLimit, using maxLimit=%d", maxLimit)
				q.Limit = maxLimit
			}
		default:
			return nil, fmt.Errorf("invalid query key: %s", parts[i])
		}
	}
	return q, nil
}

func str2id(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}
