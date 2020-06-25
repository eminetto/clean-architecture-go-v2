package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go/api/handler"
	"github.com/eminetto/clean-architecture-go/config"
	"github.com/eminetto/clean-architecture-go/pkg/bookmark"
	"github.com/eminetto/clean-architecture-go/pkg/middleware"
	"github.com/eminetto/clean-architecture-go/pkg/metric"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	r := mux.NewRouter()

	mPool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer mPool.Close()

	bookmarkRepo := bookmark.NewMongoRepository(mPool, config.MONGODB_DATABASE)
	bookmarkService := bookmark.NewService(bookmarkRepo)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)
	//bookmark
	handler.MakeBookmarkHandlers(r, *n, bookmarkService)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
