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

func borrowBook(bManager book.Manager, uManager user.Manager, loanUseCase loan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error borrowing book"
		vars := mux.Vars(r)
		bID, err := entity.StringToID(vars["book_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		b, err := bManager.Get(bID)
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
		u, err := uManager.Get(uID)
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
		err = loanUseCase.Borrow(u, b)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func returnBook(bManager book.Manager, loanUseCase loan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error returning book"
		vars := mux.Vars(r)
		bID, err := entity.StringToID(vars["book_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		b, err := bManager.Get(bID)
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
		err = loanUseCase.Return(b)
		if err != nil && err != domain.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

//MakeLoanHandlers make url handlers
func MakeLoanHandlers(r *mux.Router, n negroni.Negroni, bManager book.Manager, uManager user.Manager, loanUseCase loan.UseCase) {
	r.Handle("/v1/loan/borrow/{book_id}/{user_id}", n.With(
		negroni.Wrap(borrowBook(bManager, uManager, loanUseCase)),
	)).Methods("GET", "OPTIONS").Name("borrowBook")

	r.Handle("/v1/loan/return/{book_id}", n.With(
		negroni.Wrap(returnBook(bManager, loanUseCase)),
	)).Methods("GET", "OPTIONS").Name("returnBook")
}
