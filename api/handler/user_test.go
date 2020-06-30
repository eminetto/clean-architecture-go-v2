package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/api/presenter"
	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_listUsers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("listUsers").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)
	u := user.NewFixtureUser()
	m.EXPECT().
		List().
		Return([]*user.User{u}, nil)
	ts := httptest.NewServer(listUsers(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listUsers_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockManager(controller)
	ts := httptest.NewServer(listUsers(m))
	defer ts.Close()
	m.EXPECT().
		Search("dio").
		Return(nil, domain.ErrNotFound)
	res, err := http.Get(ts.URL + "?name=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listUsers_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockManager(controller)
	u := user.NewFixtureUser()
	m.EXPECT().
		Search("ozzy").
		Return([]*user.User{u}, nil)
	ts := httptest.NewServer(listUsers(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=ozzy")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("createUser").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)

	m.EXPECT().
		Create(gomock.Any()).
		Return(entity.NewID(), nil)
	h := createUser(m)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"name": "ozzy",
"email": "ozzy@hell.com",
"password": "asasa",
"first_name":"Ozzy",
"last_name":"Osbourne"
}`)
	resp, _ := http.Post(ts.URL+"/v1/user", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.User
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "Ozzy Osbourne", fmt.Sprintf("%s %s", u.FirstName, u.LastName))
}

func Test_getUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("getUser").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	u := user.NewFixtureUser()
	m.EXPECT().
		Get(u.ID).
		Return(u, nil)
	handler := getUser(m)
	r.Handle("/v1/user/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/user/" + u.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.User
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("deleteUser").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	u := user.NewFixtureUser()
	m.EXPECT().Delete(u.ID).Return(nil)
	handler := deleteUser(m)
	req, _ := http.NewRequest("DELETE", "/v1/user/"+u.ID.String(), nil)
	r.Handle("/v1/user/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
