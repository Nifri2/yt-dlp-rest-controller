package kvdb

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var db = sync.Map{}
var termchan = make(chan bool)

func Init() {
	log.Info().Msg("Starting KVDB")
	go Checkerd()
}

func Get(key string) (any, bool) {
	log.Info().Str("key", key).Msg("Get")
	val, ok := db.Load(key)
	return val, ok
}

func Set(key string, val int64) {
	val_str := fmt.Sprintf("%d", val)
	log.Info().Str("key", key).Str("val", val_str).Msg("Set")
	db.Store(key, val)
}

func Delete(key string) {
	log.Info().Str("key", key).Msg("Delete")
	db.Delete(key)
}

func Iter(f func(key string, val string)) {
	db.Range(func(key, val interface{}) bool {
		fmt.Println(key, val)
		f(key.(string), val.(string))
		return true
	})
}

func Checkerd() {
	for {
		select {
		case <-termchan:
			return
		default:
			db.Range(func(key, val interface{}) bool {
				if time.Now().Unix()-val.(int64) > 300 {
					log.Info().Str("key", key.(string)).Msg("Expired")
					Delete(key.(string))
					// iterate over folder

					p, err := os.ReadDir("../static")
					if err != nil {
						log.Error().Err(err).Msg("Error reading folder")
						panic("could not read static folder")
					}
					for _, file := range p {
						if strings.Split(file.Name(), ".")[0] == key.(string) {
							log.Info().Str("file", file.Name()).Msg("Deleting file")
							os.Remove("../static/" + file.Name())
						}
					}

				}
				return true
			})
		}
	}
}
