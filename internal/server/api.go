package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/jq"
	"github.com/random-j-farmer/zkill-mirror/internal/blobs"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
)

type apiQuery struct {
	KillID uint64
}

func apiHandler(w http.ResponseWriter, r *http.Request, url string) {
	q := unmarshalQuery(url)

	if q.KillID > 0 {
		ref, err := db.ByKillID(q.KillID)
		if err != nil {
			apiError(w, r, errors.Wrap(err, "db.ByKillID"))
			return
		}

		if config.Verbose() {
			logRequestf(r, "killID %d: retrieving %v", q.KillID, ref)
		}

		b, err := blobs.DB.Read(ref)
		if err != nil {
			apiError(w, r, errors.Wrap(err, "bobstore.read"))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err = w.Write(b)
		if err != nil {
			logRequestf(r, "error writing response: %v", err)
		}
		return
	}

	newest, err := db.Newest(1000)
	if err != nil {
		apiError(w, r, errors.Wrap(err, "db.Newest"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = fmt.Fprint(w, "[")
	if err != nil {
		apiError(w, r, err)
		return
	}

	lastIdx := len(newest) - 1
	for i, ref := range newest {
		logRequestf(r, "newest refs: %v", ref)

		b, err2 := blobs.DB.Read(ref)
		if err2 != nil {
			apiError(w, r, errors.Wrap(err2, "bobstore.read"))
			return
		}

		_, err2 = fmt.Fprintf(w, "%s", b)
		if err2 != nil {
			apiError(w, r, err2)
			return
		}

		if i < lastIdx {
			_, err2 = fmt.Fprint(w, ",")
			if err2 != nil {
				apiError(w, r, err2)
				return
			}
		}
	}

	_, err = fmt.Fprint(w, "]")
	if err != nil {
		apiError(w, r, err)
		return
	}
}

func apiError(w http.ResponseWriter, r *http.Request, err error) {
	logRequestf(r, "error: %v", err)
	http.Error(w, fmt.Sprintf("error: %v", err), 500)
}

func unmarshalQuery(url string) *apiQuery {
	parts := strings.Split(url, "/")
	m := make(map[string]interface{})
	ll := len(parts) - len(parts)&1
	for i := 0; i < ll; i = i + 2 {
		m[parts[i]] = parts[i+1]
	}

	q := jq.New(m)

	return &apiQuery{
		KillID: uint64(q.UInt64("killID")),
	}
}
