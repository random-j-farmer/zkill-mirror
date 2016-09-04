package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/blobs"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
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
	Hot           time.Duration
	Limit         int
}

func apiHandler(w http.ResponseWriter, r *http.Request, url string) {
	q, err := unmarshalQuery(url)
	if err != nil {
		apiError(w, r, err)
		return
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

	case q.Hot > 0:
		err = apiHot(w, r, q)

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

	if config.Verbose() {
		logRequestf(r, "killID %d: retrieving %v", q.KillID, ref)
	}

	b, err := blobs.DB.Read(ref)
	if err != nil {
		return errors.Wrap(err, "bobstore.read")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(b)
	if err != nil {
		return errors.Wrap(err, "response.write")
	}

	return nil
}

func apiByCharacterID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByCharacterID(q.CharacterID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByCharacterID")
	}

	logRequestf(r, "apiByCharacterID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiByCorporationID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByCorporationID(q.CorporationID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByCorporationID")
	}

	logRequestf(r, "apiByCorporationID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiByAllianceID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByAllianceID(q.AllianceID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByAllianceID")
	}

	logRequestf(r, "apiByAllianceID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiBySystemID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.BySystemID(q.SystemID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.BySystemID")
	}

	logRequestf(r, "apiBySystemID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiByRegionID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByRegionID(q.RegionID, q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.ByRegionID")
	}

	logRequestf(r, "apiByRegionID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiHot(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	stats, err := db.Hot(q.Hot)
	if err != nil {
		return errors.Wrap(err, "db.Hot")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write([]byte{'['})
	if err != nil {
		return errors.Wrap(err, "response.Write")
	}

	size := len(stats)
	for i, stat := range stats {
		b, err := json.Marshal(stat)
		if err != nil {
			return errors.Wrap(err, "json.Marshal")
		}

		// append a , or ] so we only have to do one write ...
		if i == size-1 {
			b = append(b, ']')
		} else {
			b = append(b, ',')
		}

		_, err = w.Write(b)
		if err != nil {
			return errors.Wrap(err, "response.write")
		}
	}
	if size == 0 {
		_, err := w.Write([]byte{']'})
		if err != nil {
			return errors.Wrap(err, "response.Write")
		}
	}

	return nil
}

func apiNewest(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.Newest(q.Limit)
	if err != nil {
		return errors.Wrap(err, "db.Newest")
	}

	logRequestf(r, "apiNewest: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiWriteResponse(w http.ResponseWriter, r *http.Request, refs []bobstore.Ref) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte{'['})
	if err != nil {
		return errors.Wrap(err, "response.Write")
	}

	size := len(refs)
	for i, ref := range refs {
		b, err := blobs.DB.Read(ref)
		if err != nil {
			return errors.Wrapf(err, "bobstore.Read %s", ref)
		}

		if len(b) == cap(b) {
			logRequestf(r, "warning: at cap for: %s - inefficient append", ref)
		}
		// append a , or ] so we only have to do one write ...
		if i == size-1 {
			b = append(b, ']')
		} else {
			b = append(b, ',')
		}

		_, err = w.Write(b)
		if err != nil {
			return errors.Wrap(err, "response.write")
		}
	}
	if size == 0 {
		_, err := w.Write([]byte{']'})
		if err != nil {
			return errors.Wrap(err, "response.Write")
		}
	}

	return nil
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
		case "systemID":
			q.SystemID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "regionID":
			q.RegionID, err = str2id(parts[i+1])
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.ParseUInt %s", parts[i])
			}
		case "hot":
			q.Hot, err = time.ParseDuration(parts[i+1])
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
