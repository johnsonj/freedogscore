package handlers

import (
	"fmt"
	"net/http"
)

type Upload struct{}

func (Upload) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is dog")
}
