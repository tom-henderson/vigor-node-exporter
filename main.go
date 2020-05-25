package main

import (
	"./Vigor"
	"os"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var username = os.Getenv("VIGOR_USERNAME")
var password = os.Getenv("VIGOR_PASSWORD")
var ip = os.Getenv("VIGOR_IP")

var vigor *Vigor.Vigor

func loginIfError(err error) {
	if err != nil {
		print(err)
		vigor.Login(username, password)
	}
}

func main() {
	var err error
	vigor, err = Vigor.New(ip)
	if err != nil {
		panic(err)
	}
	vigor.Login(username, password)

	vigor.UpdateStatus()
	vigor.FetchStatus()

	go func() {
		for {
			time.Sleep(5 * time.Second)

			loginIfError(vigor.UpdateStatus())
			loginIfError(vigor.FetchStatus())
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9103", nil))
}
