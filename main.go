package main

import (
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

var redisUrl string = os.Getenv("REDIS_URL")

func ping (w http.ResponseWriter, r *http.Request) {
	c, err := redis.DialURL(redisUrl)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	defer c.Close()

	_, err = redis.String(c.Do("PING"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
}

func main () {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":6378", nil)
}
