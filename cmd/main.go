package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/eminetto/clean-architecture-go-v2/infrastructure/repository"
	"github.com/eminetto/clean-architecture-go-v2/usecase/book"

	"github.com/eminetto/clean-architecture-go-v2/config"
	_ "github.com/go-sql-driver/mysql"

	"github.com/eminetto/clean-architecture-go-v2/pkg/metric"
)

func handleParams() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Invalid query")
	}
	return os.Args[1], nil
}

func main() {
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	appMetric := metric.NewCLI("search")
	appMetric.Started()
	query, err := handleParams()
	if err != nil {
		log.Fatal(err.Error())
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	repo := repository.NewBookMySQL(db)
	service := book.NewService(repo)
	all, err := service.SearchBooks(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range all {
		fmt.Printf("%s %s \n", j.Title, j.Author)
	}
	appMetric.Finished()
	err = metricService.SaveCLI(appMetric)
	if err != nil {
		log.Fatal(err)
	}
}
