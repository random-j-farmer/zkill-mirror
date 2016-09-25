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
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/random-j-farmer/zkill-mirror/internal/assets"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
)

// Serve the web application.
func Serve() error {
	if config.CacheTemplates() {
		mustParseTemplates()
	}
	rootPath := dirPath("")
	rootRegex := regexp.MustCompile(fmt.Sprintf("^(%s)$", rootPath))
	apiPath := dirPath("api")
	apiRegex := regexp.MustCompile(fmt.Sprintf("^%s(.*)$", apiPath))
	http.HandleFunc(dirPath(""), makeHandler(rootRegex, rootHandler))
	http.HandleFunc(dirPath("api"), makeHandler(apiRegex, apiHandler))
	http.HandleFunc(simplePath("search"), makeHandler(nil, searchHandler))

	pfs := &prefixedAssetFS{
		Prefix: dirPath(""), // only remove zkm_url_prefix, not the static bit
		fs:     &assetfs.AssetFS{Asset: assets.Asset, AssetDir: assets.AssetDir, AssetInfo: assets.AssetInfo, Prefix: ""},
	}
	http.Handle(dirPath("static"), http.FileServer(pfs))
	return listenAndServe()
}

type prefixedAssetFS struct {
	Prefix string
	fs     *assetfs.AssetFS
}

func (pfs *prefixedAssetFS) Open(name string) (http.File, error) {
	name2 := name
	if strings.HasPrefix(name, pfs.Prefix) {
		name2 = name[len(pfs.Prefix):]
	}
	if config.Debug() {
		log.Printf("prefixedAssetFS: %s => %s", name, name2)
	}
	return pfs.fs.Open(name2)
}

// simplePath is a non-/ terminated path, starts with prefix
func simplePath(s string) string {
	p := path.Join(config.URLPrefix(), s)
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return p
}

// dirPath for subtree ... joins with prefix and ends with /
func dirPath(s string) string {
	p := simplePath(s)
	if !strings.HasSuffix(p, "/") {
		p = p + "/"
	}

	// log.Printf("dirPath: %s => >%s<", s, p)
	return p
}

func listenAndServe() error {
	on := fmt.Sprintf(":%d", config.Port())
	log.Printf("Serving on %s", on)
	return http.ListenAndServe(on, nil)
}

var cachedTemplates = make(map[string]*template.Template)

var templateFuncs = template.FuncMap{
	"json":           jsonMarshal,
	"isk":            formatISK,
	"evenOdd":        evenOrOdd,
	"dirPath":        dirPath,
	"currentEVETime": currentEVETime,
}

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

func currentEVETime() string {
	return time.Now().UTC().Format(db.EveTimeFormat)
}

func executeTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) error {
	if strings.HasSuffix(name, ".json") {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	layout := name
	if strings.HasSuffix(name, ".html") {
		layout = "layout" // execute layout template!
	}

	return getTemplate(name).ExecuteTemplate(w, layout, data)
}

func getTemplate(name string) *template.Template {
	if config.CacheTemplates() {
		return cachedTemplates[name]
	}
	var layoutBody = string(assets.MustAsset("layouts/layout.html"))
	var body = string(assets.MustAsset("templates/" + name))

	tmpl, err := template.New(name).Funcs(templateFuncs).Parse(layoutBody)
	if err != nil {
		log.Printf("error parsing layout.html: %v", err)
	}

	tmpl, err = tmpl.Parse(body)
	if err != nil {
		log.Printf("error parsing %s: %v", name, err)
	}
	return tmpl
}

func mustParseTemplates() {
	var layoutBody = string(assets.MustAsset("layouts/layout.html"))
	for _, name := range assets.AssetNames() {
		if strings.HasPrefix(name, "templates/") {
			tname := name[len("templates/"):]
			if config.Verbose() {
				log.Printf("cached template: %s template: %s", name, tname)
			}
			var body = string(assets.MustAsset(name))
			cachedTemplates[tname] = template.Must(template.Must(template.New(tname).Funcs(templateFuncs).Parse(layoutBody)).Parse(body))
		}
	}
}

func logRequestf(r *http.Request, fmtstr string, args ...interface{}) {
	rstr := fmt.Sprintf("%v@%s> ", r.RemoteAddr, r.URL.Path)
	msg := fmt.Sprintf(fmtstr, args...)
	log.Printf("%s %s", rstr, msg)
}

func rootHandler(w http.ResponseWriter, r *http.Request, url string) {
	redir := dirPath("api") + "activity/1h/"
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
