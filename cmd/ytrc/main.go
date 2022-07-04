package main

import (
	"fmt"
	"net/http"
	kvdb "nifri2/ytrc/db"
	funcs "nifri2/ytrc/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	bin := funcs.Find_bin()
	if bin == "not found" {
		panic("yt-dlp not found in $PATH")
	}

	fmt.Printf("Using: %s \n", bin)

	kvdb.Init()
	//kvdb.Set("kvdb.go", time.Now().Unix())

	router := mux.NewRouter()

	router.HandleFunc("/", funcs.Index).Methods("GET")
	router.HandleFunc("/grab", funcs.Grab).Methods("GET")

	log.Info().Str("port", ":4000").Msg("API is running")

	// start fileserver in this folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.ListenAndServe("0.0.0.0:4000", router)

}
