package funcs

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	kvdb "nifri2/ytrc/db"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/rs/zerolog/log"
)

func Test_func() {
	fmt.Println("test")
}

func URL_args_query(req *http.Request, query string) ([]string, error) {
	q := req.URL.Query()
	if _, ok := q[query]; ok {
		return q[query], nil
	} else {
		return nil, fmt.Errorf("query not found")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "{\"message\":\"Index\"}"
	fmt.Fprint(w, res)
}

func Grabber(url string) string {
	log.Info().Str("url", url).Msg("Grabber")

	// build hash of the url
	hash := sha256.New()
	hash.Write([]byte(url))
	hash_val := fmt.Sprintf("%x", hash.Sum(nil))
	name := "../static/" + hash_val + ".%(ext)s"

	// kvdb set url, unix time
	kvdb.Set(hash_val, time.Now().Unix())

	ytp := cmd.NewCmd("yt-dlp", "-o", name, url)
	infoChan := ytp.Start()

	// Failsafe to prevent infinite loop
	go func() {
		<-time.After(10 * time.Minute)
		ytp.Stop()
	}()

	for {
		select {
		case finalStatus := <-infoChan:
			log.Info().Str("status", finalStatus.Cmd).Msg("Done")
			return kvdb.Serve(hash_val)
		default:
			continue
		}
	}
}

func Find_bin() string {
	var bin string
	var found bool
	var sep string
	var psep string

	env_path := os.Getenv("PATH")
	os_v := runtime.GOOS

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
