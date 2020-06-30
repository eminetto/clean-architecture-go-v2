package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/pkg/password"

	"github.com/eminetto/clean-architecture-go-v2/domain/usecase/loan"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/api/handler"
	"github.com/eminetto/clean-architecture-go-v2/api/middleware"
	"github.com/eminetto/clean-architecture-go-v2/config"
	"github.com/eminetto/clean-architecture-go-v2/pkg/metric"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	bookRepo := book.NewMySQLRepository(db)
	bookManager := book.NewManager(bookRepo)

	userRepo := user.NewMySQLRepoRepository(db)
	userManager := user.NewManager(userRepo, password.NewService())

	loanUseCase := loan.NewUseCase(userManager, bookManager)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)
	//book
	handler.MakeBookHandlers(r, *n, bookManager)

	//user
	handler.MakeUserHandlers(r, *n, userManager)

	//loan
	handler.MakeLoanHandlers(r, *n, bookManager, userManager, loanUseCase)

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
