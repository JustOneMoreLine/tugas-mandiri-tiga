package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var router *mux.Router
var log *logrus.Logger


func InitializeLogger() {
	log = logrus.New()
	conn, err := tls.Dial("tcp", "8d97ce45-37df-4f05-9c80-d54b6887c2ab-ls.logit.io:18595", &tls.Config{ RootCAs: nil }) 
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "tugas-mandiri-tiga"}))
	log.Hooks.Add(hook)
}

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	router.HandleFunc("/echo", EchoServiceMonitor).Methods("GET")
	router.HandleFunc("/", Ping).Methods("GET")
}

func EchoServiceMonitor(w http.ResponseWriter, r *http.Request) {
	var headers = make(map[string]interface{})
	for name, values := range r.URL.Query() {
		for _, value := range values {
			headers[name] = value
		}
	}

	check := EchoService(w, headers)
	if check == nil {
		ctx := log.WithFields(logrus.Fields{"method": "main.Echo"})
		ctx.Warn("Echo Service Failed!")
	} else {
		ctx := log.WithFields(check)
		ctx.Info("Echo Service Success!")
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	InitializeLogger()
	CreateRouter()
	InitializeRoute()
	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}


