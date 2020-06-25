package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/api/presenter"
	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_UserIndex(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userIndex").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)
	u := user.NewFixtureUser()
	service.EXPECT().
		List().
		Return([]*user.User{u}, nil)
	ts := httptest.NewServer(userIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_UserIndexNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(userIndex(service))
	defer ts.Close()
	service.EXPECT().
		Search("dio").
		Return(nil, domain.ErrNotFound)
	res, err := http.Get(ts.URL + "?name=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_UserSearch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	u := user.NewFixtureUser()
	service.EXPECT().
		Search("ozzy").
		Return([]*user.User{u}, nil)
	ts := httptest.NewServer(userIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=ozzy")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_UserAdd(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userAdd").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)

	service.EXPECT().
		Create(gomock.Any()).
		Return(entity.NewID(), nil)
	h := userAdd(service)

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

func Test_UserFind(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userFind").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	u := user.NewFixtureUser()
	service.EXPECT().
		Get(u.ID).
		Return(u, nil)
	handler := userFind(service)
	r.Handle("/v1/bookmark/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/bookmark/" + b.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *entity.Bookmark
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, b.ID, d.ID)
}

//
//func TestBookmarkRemove(t *testing.T) {
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	service := mock.NewMockUseCase(controller)
//	r := mux.NewRouter()
//	n := negroni.New()
//	MakeBookmarkHandlers(r, *n, service)
//	path, err := r.GetRoute("bookmarkDelete").GetPathTemplate()
//	assert.Nil(t, err)
//	assert.Equal(t, "/v1/bookmark/{id}", path)
//	b := &entity.Bookmark{
//		ID:          entity.NewID(),
//		Name:        "Elton Minetto",
//		Description: "Minetto's page",
//		Link:        "http://www.eltonminetto.net",
//		Tags:        []string{"golang", "php", "linux", "mac"},
//		Favorite:    false,
//	}
//	service.EXPECT().Delete(b.ID).Return(nil)
//	handler := bookmarkDelete(service)
//	req, _ := http.NewRequest("DELETE", "/v1/bookmark/"+b.ID.String(), nil)
//	r.Handle("/v1/bookmark/{id}", handler).Methods("DELETE", "OPTIONS")
//	rr := httptest.NewRecorder()
//	r.ServeHTTP(rr, req)
//	assert.Equal(t, http.StatusOK, rr.Code)
//}
