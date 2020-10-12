package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eminetto/clean-architecture-go-v2/domain/usecase/book"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/api/presenter"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/gorilla/mux"
)

func listBooks(service book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading books"
		var data []*entity.Book
		var err error
		title := r.URL.Query().Get("title")
		switch {
		case title == "":
			data, err = service.ListBooks()
		default:
			data, err = service.SearchBooks(title)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Book
		for _, d := range data {
			toJ = append(toJ, &presenter.Book{
				ID:       d.ID,
				Title:    d.Title,
				Author:   d.Author,
				Pages:    d.Pages,
				Quantity: d.Quantity,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createBook(service book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding book"
		var input struct {
			Title    string `json:"title"`
			Author   string `json:"author"`
			Pages    int    `json:"pages"`
			Quantity int    `json:"quantity"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateBook(input.Title, input.Author, input.Pages, input.Quantity)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Book{
			ID:       id,
			Title:    input.Title,
			Author:   input.Author,
			Pages:    input.Pages,
			Quantity: input.Quantity,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getBook(service book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading book"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetBook(id)
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Book{
			ID:       data.ID,
			Title:    data.Title,
			Author:   data.Author,
			Pages:    data.Pages,
			Quantity: data.Quantity,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteBook(service book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing bookmark"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteBook(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeBookHandlers make url handlers
func MakeBookHandlers(r *mux.Router, n negroni.Negroni, service book.UseCase) {
	r.Handle("/v1/book", n.With(
		negroni.Wrap(listBooks(service)),
	)).Methods("GET", "OPTIONS").Name("listBooks")

	r.Handle("/v1/book", n.With(
		negroni.Wrap(createBook(service)),
	)).Methods("POST", "OPTIONS").Name("createBook")

	r.Handle("/v1/book/{id}", n.With(
		negroni.Wrap(getBook(service)),
	)).Methods("GET", "OPTIONS").Name("getBook")

	r.Handle("/v1/book/{id}", n.With(
		negroni.Wrap(deleteBook(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteBook")
}
