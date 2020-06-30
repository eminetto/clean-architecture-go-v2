package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/api/presenter"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listUsers(manager user.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*user.User
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = manager.List()
		default:
			data, err = manager.Search(name)
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

func createUser(manager user.Manager) http.Handler {
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
		//TODO: validate data ;)
		u := &user.User{
			ID:        entity.NewID(),
			Email:     input.Email,
			Password:  input.Password,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			CreatedAt: time.Now(),
		}
		u.ID, err = manager.Create(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.User{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
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

func getUser(manager user.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := manager.Get(id)
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

func deleteUser(manager user.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
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

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, manager user.Manager) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(listUsers(manager)),
	)).Methods("GET", "OPTIONS").Name("listUsers")

	r.Handle("/v1/user", n.With(
		negroni.Wrap(createUser(manager)),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(getUser(manager)),
	)).Methods("GET", "OPTIONS").Name("getUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(deleteUser(manager)),
	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}
