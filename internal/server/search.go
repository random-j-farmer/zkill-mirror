package server

import (
	"fmt"
	"net/http"

	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
)

func searchHandler(w http.ResponseWriter, r *http.Request, _ string) {
	q := r.FormValue("q")
	if len(q) < 3 {
		logRequestf(r, "error: query too short: %s", q)
		http.Error(w, fmt.Sprintf("error: query to short: %s", q), 400)
		return
	}

	logRequestf(r, "searching for: %s", q)
	result, err := db.Search(q)
	if err != nil {
		logRequestf(r, "error: %v", err)
		http.Error(w, fmt.Sprintf("error: %v", err), 500)
		return
	}

	if config.Debug() {
		logRequestf(r, "result: %#v", result)
	}

	executeTemplate(w, r, "search.html", result)
	return
}
