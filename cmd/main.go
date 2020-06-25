package main

import (
	"fmt"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"
)

//func handleParams() (string, error) {
//	if len(os.Args) < 2 {
//		return "", errors.New("Invalid query")
//	}
//	return os.Args[1], nil
//}

func main() {
	u := &user.User{
		ID: entity.NewID(),
		FirstName: "Elton",
	}
	b := &book.Book{
		ID: entity.NewID(),
		Title:"Book",
	}
	fmt.Println(u)
	fmt.Println(b)
	//metricService, err := metric.NewPrometheusService()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//appMetric := metric.NewCLI("search")
	//appMetric.Started()
	//query, err := handleParams()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//session, err := mgo.Dial(config.MONGODB_HOST)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//defer session.Close()
	//
	//mPool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	//defer mPool.Close()
	//
	//bookmarkRepo := bookmark.NewMongoRepository(mPool, config.MONGODB_DATABASE)
	//bookmarkService := bookmark.NewService(bookmarkRepo)
	//all, err := bookmarkService.Search(query)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if len(all) == 0 {
	//	log.Fatal(entity.ErrNotFound.Error())
	//}
	//for _, j := range all {
	//	fmt.Printf("%s %s %v \n", j.Name, j.Link, j.Tags)
	//}
	//appMetric.Finished()
	//err = metricService.SaveCLI(appMetric)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
