package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dark705/otus/hw17/internal/helpers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
)

type Server struct {
	c  Config
	l  *logrus.Logger
	ws *http.Server
	ps *http.Server
}

type Config struct {
	HttpListen       string
	PrometheusListen string
}

func NewServer(conf Config, log *logrus.Logger) Server {
	prometheusMiddlewareHandler := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	return Server{
		c:  conf,
		l:  log,
		ws: &http.Server{Addr: conf.HttpListen, Handler: prometheusMiddlewareHandler.Handler("", logRequest(ServeHTTP, log))},
		ps: &http.Server{Addr: conf.PrometheusListen, Handler: promhttp.Handler()},
	}
}

func (s *Server) RunServer() {
	go func() {
		s.l.Infoln("Start HTTP server: ", s.c.HttpListen)
		err := s.ws.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			helpers.FailOnError(err, "Fail start HTTP Server")
		}
	}()

	go func() {
		s.l.Infoln("Start Prometheus Http metrics server: ", s.c.PrometheusListen)
		err := s.ps.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			helpers.FailOnError(err, "Fail start Prometheus Http metrics server")
		}
	}()

}

func (s *Server) Shutdown() {
	s.l.Infoln("Shutdown HTTP server... ")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := s.ws.Shutdown(ctx)
	if err != nil {
		s.l.Errorln("Fail Shutdown HTTP server")
		return
	}
	s.l.Infoln("Success Shutdown HTTP server")

	s.l.Infoln("Shutdown Prometheus metrics server... ")
	err = s.ps.Shutdown(ctx)
	if err != nil {
		s.l.Errorln("Fail Shutdown Prometheus Http metrics server")
		return
	}
	s.l.Infoln("Success shutdown Prometheus Http metrics server")
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Calendar")
	_, _ = w.Write([]byte("Hello world"))
}

//middleware logger
func logRequest(h http.HandlerFunc, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Infoln(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		h(w, r)
	}
}
