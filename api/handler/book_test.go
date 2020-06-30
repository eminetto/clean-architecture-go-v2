package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_listBooks(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	manager := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, manager)
	path, err := r.GetRoute("listBooks").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/book", path)
	b := book.NewFixtureBook()
	manager.EXPECT().
		List().
		Return([]*book.Book{b}, nil)
	ts := httptest.NewServer(listBooks(manager))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listBooks_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	manager := mock.NewMockManager(controller)
	ts := httptest.NewServer(listBooks(manager))
	defer ts.Close()
	manager.EXPECT().
		Search("book of books").
		Return(nil, domain.ErrNotFound)
	res, err := http.Get(ts.URL + "?title=book+of+books")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listBooks_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	manager := mock.NewMockManager(controller)
	b := book.NewFixtureBook()
	manager.EXPECT().
		Search("ozzy").
		Return([]*book.Book{b}, nil)
	ts := httptest.NewServer(listBooks(manager))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?title=ozzy")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createBook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	manager := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, manager)
	path, err := r.GetRoute("createBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/book", path)

	manager.EXPECT().
		Create(gomock.Any()).
		Return(entity.NewID(), nil)
	h := createBook(manager)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
  "title": "I Am Ozzy",
  "author": "Ozzy Osbourne",
  "pages": 294,
  "quantity":1
}`)
	resp, _ := http.Post(ts.URL+"/v1/book", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var b *book.Book
	json.NewDecoder(resp.Body).Decode(&b)
	assert.Equal(t, "Ozzy Osbourne", b.Author)
}

func Test_getBook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	manager := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, manager)
	path, err := r.GetRoute("getBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/book/{id}", path)
	b := book.NewFixtureBook()
	manager.EXPECT().
		Get(b.ID).
		Return(b, nil)
	handler := getBook(manager)
	r.Handle("/v1/book/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/book/" + b.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *book.Book
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, b.ID, d.ID)
}

func Test_deleteBook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	manager := mock.NewMockManager(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, manager)
	path, err := r.GetRoute("deleteBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/book/{id}", path)
	b := book.NewFixtureBook()
	manager.EXPECT().Delete(b.ID).Return(nil)
	handler := deleteBook(manager)
	req, _ := http.NewRequest("DELETE", "/v1/bookmark/"+b.ID.String(), nil)
	r.Handle("/v1/bookmark/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
