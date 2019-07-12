package main

import (
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/happypath"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting...")

	go runTest()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
}

func runTest() {
	for {
		happypath.Run()
		time.Sleep(10 * time.Second)
	}
}
