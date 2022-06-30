package funcs

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "{\"message\":\"Test Succeeded\"}"
	fmt.Fprint(w, res)
}

func Grab(w http.ResponseWriter, r *http.Request) {
	log.Info().Str("url", r.URL.String()).Msg(r.RemoteAddr + ": Grab")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	q, err := URL_args_query(r, "url")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := "{\"message\":\"URL PARAM not found\"}"
		fmt.Fprint(w, res)
		return
	}

	ref := Grabber(q[0])
	res := fmt.Sprintf("{\"message\":\"%s\"}", ref)
	fmt.Fprint(w, res)
}
