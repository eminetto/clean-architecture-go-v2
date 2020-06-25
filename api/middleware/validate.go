package middleware

import (
	"context"
	"encoding/json"
	"github.com/eminetto/clean-architecture-go/pkg/entity"
	"net/http"
	"github.com/codegangsta/negroni"
)

//Validate
func Validate(e entity.Validable) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = e.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), "InputParam", e)
		next(w, r.WithContext(ctx))
	}
}
