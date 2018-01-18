package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/johnsonj/freedogscore/handlers"
)

var port = flag.Int("port", 8080, "port to bind server")

type Endpoint struct {
	Name     string
	Path     string
	HandleFn func(w http.ResponseWriter, r *http.Request)
}

var endpoints []Endpoint

func init() {
	flag.Parse()

	upload := handlers.Upload{}
	endpoints = []Endpoint{{"Upload", "/upload", upload.Handle}}
}

const index = `
<html>
<head><title>FreeDogSocre: Know your dog</title></head>
<body>
<h1>Endpoints</h1>
<ul>
{{ range . }}
  <li><a href="{{ .Path }}">{{ .Name }}</a></li>
{{ end }}
</ul>
</body>
</html>
`

func main() {
	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint.Path, endpoint.HandleFn)
	}

	t := template.Must(template.New("index").Parse(index))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, endpoints)
	})

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Printf("listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, http.DefaultServeMux))
}
