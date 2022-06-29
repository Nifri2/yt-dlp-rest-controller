package funcs

import (
	"fmt"
	"html"
	"net/http"
)

func Grab(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := fmt.Sprintf("{\"message\":\"%s\"}", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, res)
}
