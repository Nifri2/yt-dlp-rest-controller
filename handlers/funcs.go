package funcs

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	kvdb "nifri2/ytrc/db"
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
	hash_val := fmt.Sprintf("%x ", hash.Sum(nil))
	name := hash_val + ".%(ext)s"

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
			return "Done"
		default:
			continue
		}
	}
}
