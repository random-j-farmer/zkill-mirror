// Package server contains the http server for the web app
package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/random-j-farmer/zkill-mirror/internal/assets"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
)

// Serve the web application.
func Serve() error {
	if config.CacheTemplates() {
		mustParseTemplates()
	}
	http.HandleFunc("/", makeHandler(regexp.MustCompile("^(/)$"), rootHandler))
	http.HandleFunc("/api/", makeHandler(regexp.MustCompile("^/api/(.*)"), apiHandler))
	fs := &assetfs.AssetFS{Asset: assets.Asset, AssetDir: assets.AssetDir, AssetInfo: assets.AssetInfo, Prefix: ""}
	http.Handle("/static/", http.FileServer(fs))
	return listenAndServe()
}

func listenAndServe() error {
	on := fmt.Sprintf(":%d", config.Port())
	log.Printf("Serving on %s", on)
	return http.ListenAndServe(on, nil)
}

var cachedTemplates *template.Template

var templateFuncs = template.FuncMap{"json": jsonMarshal}

func jsonMarshal(data interface{}) template.HTML {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("json marshall error: %v", err)
	}
	return template.HTML(b)
}

func executeTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) error {
	if strings.HasSuffix(name, ".json") {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	return getTemplate(name).ExecuteTemplate(w, name, data)
}

func getTemplate(name string) *template.Template {
	if config.CacheTemplates() {
		return cachedTemplates
	}
	var body = string(assets.MustAsset("templates/" + name))
	return template.Must(template.New(name).Funcs(templateFuncs).Parse(body))
}

func mustParseTemplates() {
	for _, name := range assets.AssetNames() {
		if strings.HasPrefix(name, "templates/") {
			tname := name[len("templates/"):]
			if config.Verbose() {
				log.Printf("asset: %s template: %s", name, tname)
			}
			var body = string(assets.MustAsset(name))
			cachedTemplates = template.Must(template.New(tname).Funcs(templateFuncs).Parse(body))
		}
	}
}

func logRequestf(r *http.Request, fmtstr string, args ...interface{}) {
	rstr := fmt.Sprintf("%v@%s> ", r.RemoteAddr, r.URL.Path)
	msg := fmt.Sprintf(fmtstr, args...)
	log.Printf("%s %s", rstr, msg)
}

func rootHandler(w http.ResponseWriter, r *http.Request, url string) {
	redir := "/api/activity/1h/"
	if config.Verbose() {
		logRequestf(r, "redirecting to %s", redir)
	}
	http.Redirect(w, r, redir, http.StatusTemporaryRedirect)
}

func makeHandler(validPath *regexp.Regexp, fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if validPath == nil {
			fn(w, r, r.URL.Path)
			return
		}

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			logRequestf(r, "http.NotFound")
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[1])
	}
}
