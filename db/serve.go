package kvdb

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	log.Info().Str("endpoint", "endpoint").Msg("Serveing")
	// write id to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// get value from db
	val, ok := Get(key)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		res := "{\"message\":\"Key not found\"}"
		fmt.Fprint(w, res)
		return
	}
	// write value to response
	val_str := fmt.Sprintf("%d", val)
	res := fmt.Sprintf("{\"message\":\"%s\"}", val_str)
	fmt.Fprint(w, res)

}
