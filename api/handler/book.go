package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/api/presenter"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listBooks(manager book.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading books"
		var data []*book.Book
		var err error
		title := r.URL.Query().Get("title")
		switch {
		case title == "":
			data, err = manager.List()
		default:
			data, err = manager.Search(title)
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

func createBook(manager book.Manager) http.Handler {
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
		b := &book.Book{
			ID:        entity.NewID(),
			Title:     input.Title,
			Author:    input.Author,
			Pages:     input.Pages,
			Quantity:  input.Quantity,
			CreatedAt: time.Now(),
		}
		b.ID, err = manager.Create(b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Book{
			ID:       b.ID,
			Title:    b.Title,
			Author:   b.Author,
			Pages:    b.Pages,
			Quantity: b.Quantity,
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

func getBook(manager book.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading book"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := manager.Get(id)
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

func deleteBook(manager book.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing bookmark"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = manager.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeBookHandlers make url handlers
func MakeBookHandlers(r *mux.Router, n negroni.Negroni, manager book.Manager) {
	r.Handle("/v1/book", n.With(
		negroni.Wrap(listBooks(manager)),
	)).Methods("GET", "OPTIONS").Name("listBooks")

	r.Handle("/v1/book", n.With(
		negroni.Wrap(createBook(manager)),
	)).Methods("POST", "OPTIONS").Name("createBook")

	r.Handle("/v1/book/{id}", n.With(
		negroni.Wrap(getBook(manager)),
	)).Methods("GET", "OPTIONS").Name("getBook")

	r.Handle("/v1/book/{id}", n.With(
		negroni.Wrap(deleteBook(manager)),
	)).Methods("DELETE", "OPTIONS").Name("deleteBook")
}
