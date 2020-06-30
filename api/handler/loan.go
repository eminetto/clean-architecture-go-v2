package handler

import (
	"fmt"
	"net/http"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"

	"github.com/eminetto/clean-architecture-go-v2/domain/usecase/loan"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func borrowBook(bService book.Repository, uService user.Repository, loanService loan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error borrowing book"
		vars := mux.Vars(r)
		bID, err := entity.StringToID(vars["book_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		b, err := bService.Get(bID)
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if b == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		uID, err := entity.StringToID(vars["user_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		u, err := uService.Get(uID)
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if u == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		err = loanService.Borrow(u, b)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func returnBook(bService book.Repository, loanService loan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error returning book"
		vars := mux.Vars(r)
		bID, err := entity.StringToID(vars["book_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		b, err := bService.Get(bID)
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if b == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		err = loanService.Return(b)
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

//MakeLoanHandlers make url handlers
func MakeLoanHandlers(r *mux.Router, n negroni.Negroni, bService book.Repository, uService user.Repository, loanService loan.UseCase) {
	r.Handle("/v1/loan/borrow/{book_id}/{user_id}", n.With(
		negroni.Wrap(borrowBook(bService, uService, loanService)),
	)).Methods("GET", "OPTIONS").Name("borrowBook")

	r.Handle("/v1/loan/return/{book_id}", n.With(
		negroni.Wrap(returnBook(bService, loanService)),
	)).Methods("GET", "OPTIONS").Name("returnBook")
}
