package main

import (
	"fmt"
	"net/http"
	funcs "nifri2/ytrc/handlers"
	"os"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func find_bin() string {
	var bin string
	var found bool
	var sep string
	var psep string

	env_path := os.Getenv("PATH")

	os_v := runtime.GOOS
	fmt.Println("PATH:", env_path)

	switch os_v {
	case "windows":
		log.Info().Msg("OS: Windows")
		sep = ";"
		psep = "\\"
	case "darwin":
		log.Info().Msg("OS: Mac")
		sep = ":"
		psep = "/"
	case "linux":
		log.Info().Msg("OS: Linux")
		sep = ":"
		psep = "/"
	default:
		log.Warn().Msg("OS: Unknown")
		sep = ":"
		psep = "/"
	}

	for _, path := range strings.Split(env_path, sep) {
		log.Info().Str("path", path).Msg("Checking path")

		p, err := os.ReadDir(path)
		if err != nil {
			log.Error().Err(err).Msg("Error reading path")
			continue
		}

		for _, file := range p {
			switch file.Name() {
			case "yt-dlp":
				bin = path + psep + file.Name()
				log.Info().Str("bin", bin).Msg("Found yt-dlp")
				found = true
			case "yt-dlp.exe":
				bin = path + psep + file.Name()
				log.Info().Str("bin", bin).Msg("Found yt-dlp.exe")
				found = true
			case "yt-dlp_macos":
				bin = path + psep + file.Name()
				log.Info().Str("bin", bin).Msg("Found yt-dlp_macos")
				found = true
			case "yt-dlp_linux":
				bin = path + psep + file.Name()
				log.Info().Str("bin", bin).Msg("Found yt-dlp_linux")
				found = true
			}

			if found {
				return bin
			}
		}

	}

	if !found {
		log.Error().Str("file", "yt-dlp").Msg("Not found, please install yt-dlp")
	}

	bin = "not found"
	return bin
}

func main() {
	bin := find_bin()
	if bin == "not found" {
		panic("yt-dlp not found in $PATH")
	}

	fmt.Printf("Using: %s \n", bin)

	router := mux.NewRouter()

	router.HandleFunc("/grab", funcs.Grab).Methods("GET")

	log.Info().Str("port", ":4000").Msg("API is running")
	http.ListenAndServe("0.0.0.0:4000", router)

}
