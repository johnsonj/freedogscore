package main

import (
	"flag"
	"net/http"

	"github.com/GoogleCloudPlatform/cloud-functions-go/nodego"
	"github.com/johnsonj/freedogscore/handlers"
)

func main() {
	flag.Parse()

	upload := handlers.Upload{}
	http.HandleFunc(nodego.HTTPTrigger, upload.Handle)

	nodego.TakeOver()
}
