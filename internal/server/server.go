// Package server contains the http server for the web app
package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
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
	http.HandleFunc(dirPath(""), makeHandler(regexp.MustCompile("^(/)$"), rootHandler))
	http.HandleFunc(dirPath("api"), makeHandler(regexp.MustCompile("^/api/(.*)"), apiHandler))
	http.HandleFunc(simplePath("search"), makeHandler(nil, searchHandler))
	fs := &assetfs.AssetFS{Asset: assets.Asset, AssetDir: assets.AssetDir, AssetInfo: assets.AssetInfo, Prefix: ""}
	http.Handle(dirPath("static"), http.FileServer(fs))
	return listenAndServe()
}

// simplePath is a non-/ terminated path, starts with prefix
func simplePath(s string) string {
	p := "/" + path.Join(config.URLPrefix(), s)
	// log.Printf("simplePath: %s => %s", s, p)
	return p
}

// dirPath for subtree ... joins with prefix and ends with /
func dirPath(s string) string {
	p := simplePath(s)
	if strings.HasSuffix(p, "/") {
		p = p + "/"
	}

	// log.Printf("dirPath: %s => %s", s, p)
	return p
}

func listenAndServe() error {
	on := fmt.Sprintf(":%d", config.Port())
	log.Printf("Serving on %s", on)
	return http.ListenAndServe(on, nil)
}

var cachedTemplates = make(map[string]*template.Template)

var templateFuncs = template.FuncMap{"json": jsonMarshal, "isk": formatISK, "evenOdd": evenOrOdd}

func jsonMarshal(data interface{}) template.HTML {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("json marshall error: %v", err)
	}
	return template.HTML(b)
}

func formatISK(data interface{}) string {
	format := "%d.00"
	switch data.(type) {
	case float32:
		format = "%0.2f"
	case float64:
		format = "%0.2f"
	}

	return formatISKString(fmt.Sprintf(format, data))
}

var iskRE = regexp.MustCompile(`\d\d{3}[,.]`)

func formatISKString(s string) string {
	for {
		m := iskRE.FindStringSubmatch(s)
		if m == nil {
			return s
		}
		s = iskRE.ReplaceAllStringFunc(s, func(m string) string {
			return m[0:1] + "," + m[1:]
		})
	}
}

func evenOrOdd(data interface{}) string {
	if data.(int)&1 == 0 {
		return "even"
	}
	return "odd"
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
		return cachedTemplates[name]
	}
	var body = string(assets.MustAsset("templates/" + name))
	return template.Must(template.New(name).Funcs(templateFuncs).Parse(body))
}

func mustParseTemplates() {
	for _, name := range assets.AssetNames() {
		if strings.HasPrefix(name, "templates/") {
			tname := name[len("templates/"):]
			if config.Verbose() {
				log.Printf("cached template: %s template: %s", name, tname)
			}
			var body = string(assets.MustAsset(name))
			cachedTemplates[tname] = template.Must(template.New(tname).Funcs(templateFuncs).Parse(body))
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
