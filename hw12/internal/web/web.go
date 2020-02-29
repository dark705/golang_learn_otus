package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dark705/otus/hw12/internal/config"
	"github.com/sirupsen/logrus"
)

func RunServer(conf config.Config, log *logrus.Logger) {
	log.Info("Start HTTP server: ", conf.HttpListen)

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "Hello world!")

	})

	listenAndServeErr := http.ListenAndServe(conf.HttpListen, logRequest(http.DefaultServeMux, *log))
	if listenAndServeErr != nil {
		log.Error("Error on start HTTP server: ", listenAndServeErr)
		os.Exit(2)
	}
}

func logRequest(handler http.Handler, log logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}
