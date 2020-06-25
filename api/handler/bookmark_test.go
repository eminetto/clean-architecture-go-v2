package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go/pkg/bookmark/mock"
	"github.com/eminetto/clean-architecture-go/pkg/entity"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkIndex(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookmarkHandlers(r, *n, service)
	path, err := r.GetRoute("bookmarkIndex").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/bookmark", path)
	b := &entity.Bookmark{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	service.EXPECT().
		FindAll().
		Return([]*entity.Bookmark{b}, nil)
	ts := httptest.NewServer(bookmarkIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestBookmarkIndexNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(bookmarkIndex(service))
	defer ts.Close()
	service.EXPECT().
		Search("github").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?name=github")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestBookmarkSearch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	b := &entity.Bookmark{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	service.EXPECT().
		Search("minetto").
		Return([]*entity.Bookmark{b}, nil)
	ts := httptest.NewServer(bookmarkIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=minetto")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestBookmarkAdd(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookmarkHandlers(r, *n, service)
	path, err := r.GetRoute("bookmarkAdd").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/bookmark", path)

	service.EXPECT().
		Store(gomock.Any()).
		Return(entity.NewID(), nil)
	ts := httptest.NewServer(r)
	defer ts.Close()
	payload := fmt.Sprintf(`{
  "name": "Github",
  "description": "Github site",
  "link": "http://github.com",
  "tags": [
    "git",
    "social"
  ]
}`)
	resp, _ := http.Post(ts.URL+"/v1/bookmark", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var b *entity.Bookmark
	json.NewDecoder(resp.Body).Decode(&b)
	assert.Equal(t, "http://github.com", b.Link)
}

func TestBookmarkFind(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookmarkHandlers(r, *n, service)
	path, err := r.GetRoute("bookmarkFind").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/bookmark/{id}", path)
	b := &entity.Bookmark{
		ID:          entity.NewID(),
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	service.EXPECT().
		Find(b.ID).
		Return(b, nil)
	handler := bookmarkFind(service)
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

func TestBookmarkRemove(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookmarkHandlers(r, *n, service)
	path, err := r.GetRoute("bookmarkDelete").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/bookmark/{id}", path)
	b := &entity.Bookmark{
		ID:          entity.NewID(),
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    false,
	}
	service.EXPECT().Delete(b.ID).Return(nil)
	handler := bookmarkDelete(service)
	req, _ := http.NewRequest("DELETE", "/v1/bookmark/"+b.ID.String(), nil)
	r.Handle("/v1/bookmark/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
