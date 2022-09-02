package main

import (
	"github.com/serialx/hashring"
	"log"
	"net/http"
	"time"
)

var env, err = LoadConfig(".")

var Servers = []string{
	env.Shard1,
	env.Shard2,
	env.Shard3,
}

// create a hash ring with the shards
var hr = hashring.New(Servers)

func init() {

	if err != nil {
		log.Println(err)
	}
	log.Println("Connecting to database")
	ConnectPostgres(env.Shard1)
	ConnectPostgres(env.Shard2)
	ConnectPostgres(env.Shard3)
	log.Println("Connected to the databases")

}

func main() {

	srv := &http.Server{
		Addr:              ":8088",
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}

}
