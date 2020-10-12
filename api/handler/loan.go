package handler

import (
	"fmt"
	"net/http"

	"github.com/eminetto/clean-architecture-go-v2/usecase/book"
	"github.com/eminetto/clean-architecture-go-v2/usecase/user"

	"github.com/eminetto/clean-architecture-go-v2/usecase/loan"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/entity"
	"github.com/gorilla/mux"
)

func borrowBook(bookService book.UseCase, userService user.UseCase, loanService loan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error borrowing book"
		vars := mux.Vars(r)
		bID, err := entity.StringToID(vars["book_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		b, err := bookService.GetBook(bID)
		if err != nil && err != entity.ErrNotFound {
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
		u, err := userService.GetUser(uID)
		if err != nil && err != entity.ErrNotFound {
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

func returnBook(bookService book.UseCase, loanService loan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error returning book"
		vars := mux.Vars(r)
		bID, err := entity.StringToID(vars["book_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		b, err := bookService.GetBook(bID)
		if err != nil && err != entity.ErrNotFound {
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
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

//MakeLoanHandlers make url handlers
func MakeLoanHandlers(r *mux.Router, n negroni.Negroni, bookService book.UseCase, userService user.UseCase, loanService loan.UseCase) {
	r.Handle("/v1/loan/borrow/{book_id}/{user_id}", n.With(
		negroni.Wrap(borrowBook(bookService, userService, loanService)),
	)).Methods("GET", "OPTIONS").Name("borrowBook")

	r.Handle("/v1/loan/return/{book_id}", n.With(
		negroni.Wrap(returnBook(bookService, loanService)),
	)).Methods("GET", "OPTIONS").Name("returnBook")
}
