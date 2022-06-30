package kvdb

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func Test_srv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func Serve(key string) string {
	log.Info().Str("endpoint", "endpoint").Msg("Serveing")
	// write id to response

	// get value from db
	_, ok := Get(key)
	if !ok {
		res := "{\"message\":\"Key not found\"}"
		return res
	}

	// write value to response
	val_str := key

	p, err := os.ReadDir("../static")
	if err != nil {
		log.Error().Err(err).Msg("Error reading folder")
		panic("Error reading static folder")
	}

	// if val_str in
	for _, f := range p {
		if strings.Split(f.Name(), ".")[0] == key {
			val_str = f.Name()
		}
	}

	res := fmt.Sprintf(val_str)
	return res

}
