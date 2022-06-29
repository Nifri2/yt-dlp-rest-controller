package funcs

import (
	"fmt"
	"net/http"

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

func Grabber(url string) string {
	log.Info().Str("url", url).Msg("Grabber")
	ytp := cmd.NewCmd("yt-dlp", "-o", "something.mp4", url)
	infoChan := ytp.Start()
	for {
		select {
		case finalStatus := <-infoChan:
			log.Info().Str("status", finalStatus.Cmd).Msg("Done")
			return "eheh"
		default:
			continue
		}
	}
}
