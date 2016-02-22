package master

import (
	"html/template"
	"log"
	"net/http"
)

var mux *http.ServeMux
var links []link

type link struct {
	Pattern string
	Name    string
}

type info struct {
	Links   []link
	Content interface{}
}

func init() {
	mux = http.NewServeMux()
	AddContent("/", "Default", nil)
}

func Mux() *http.ServeMux {
	return mux
}

func AddContent(pattern, name string, content interface{}, files ...string) {
	files = append([]string{"master/master.html"}, files...)
	newTmpl := template.Must(template.ParseFiles(files...))
	links = append(links, link{pattern, name})
	mux.HandleFunc(pattern, handler(newTmpl, content))
}

func handler(tmpl *template.Template, content interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "master.html", info{links, content}); err != nil {
			log.Println("Master execute template error:", err)
		}
	}
}
