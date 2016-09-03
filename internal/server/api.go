package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/blobs"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
)

const maxReturns = 1000

type apiQuery struct {
	KillID        uint64
	CharacterID   uint64
	CorporationID uint64
	AllianceID    uint64
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
	refs, err := db.ByCharacterID(q.CharacterID, maxReturns)
	if err != nil {
		return errors.Wrap(err, "db.ByCharacterID")
	}

	logRequestf(r, "apiByCharacterID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiByCorporationID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByCorporationID(q.CorporationID, maxReturns)
	if err != nil {
		return errors.Wrap(err, "db.ByCorporationID")
	}

	logRequestf(r, "apiByCorporationID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiByAllianceID(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.ByAllianceID(q.AllianceID, maxReturns)
	if err != nil {
		return errors.Wrap(err, "db.ByAllianceID")
	}

	logRequestf(r, "apiByAllianceID: %d results", len(refs))
	return apiWriteResponse(w, r, refs)
}

func apiNewest(w http.ResponseWriter, r *http.Request, q *apiQuery) error {
	refs, err := db.Newest(maxReturns)
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

	q := &apiQuery{}

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
		default:
			return nil, fmt.Errorf("invalid query key: %s", parts[i])
		}
	}
	return q, nil
}

func str2id(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}
