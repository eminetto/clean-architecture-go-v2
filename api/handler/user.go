package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eminetto/clean-architecture-go-v2/usecase/user"

	"github.com/eminetto/clean-architecture-go-v2/api/presenter"

	"github.com/eminetto/clean-architecture-go-v2/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listUsers(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*entity.User
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListUsers()
		default:
			data, err = service.SearchUsers(name)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.User
		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:        d.ID,
				Email:     d.Email,
				FirstName: d.FirstName,
				LastName:  d.LastName,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var input struct {
			Email     string `json:"email"`
			Password  string `json:"password"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateUser(input.Email, input.Password, input.FirstName, input.LastName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.User{
			ID:        id,
			Email:     input.Email,
			FirstName: input.FirstName,
			LastName:  input.LastName,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.User{
			ID:        data.ID,
			Email:     data.Email,
			FirstName: data.FirstName,
			LastName:  data.LastName,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteUser(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(listUsers(service)),
	)).Methods("GET", "OPTIONS").Name("listUsers")

	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(getUser(service)),
	)).Methods("GET", "OPTIONS").Name("getUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(deleteUser(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}
