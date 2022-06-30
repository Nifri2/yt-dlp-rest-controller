package kvdb

import (
	"fmt"
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
				}
				return true
			})
		}
	}
}
