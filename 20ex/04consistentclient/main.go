package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/baronwithyou/bcache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	bcache.NewGroup("scores", 2<<10, bcache.GetterFunc(getter))

	addr := "localhost:9999"
	peers := bcache.NewHTTPPool(addr)
	log.Println("bcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}

func getter(key string) ([]byte, error) {
	log.Println("[SlowDB] search key", key)
	if v, ok := db[key]; ok {
		return []byte(v), nil
	}
	return nil, fmt.Errorf("%s not exist", key)
}
