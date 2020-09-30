package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/codegangsta/negroni"
	bmock "github.com/eminetto/clean-architecture-go-v2/domain/usecase/book/mock"
	lmock "github.com/eminetto/clean-architecture-go-v2/domain/usecase/loan/mock"
	umock "github.com/eminetto/clean-architecture-go-v2/domain/usecase/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_borrowBook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	uMock := umock.NewMockUseCase(controller)
	bMock := bmock.NewMockUseCase(controller)
	lMock := lmock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeLoanHandlers(r, *n, bMock, uMock, lMock)
	path, err := r.GetRoute("borrowBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/loan/borrow/{book_id}/{user_id}", path)
	handler := borrowBook(bMock, uMock, lMock)
	r.Handle("/v1/loan/borrow/{book_id}/{user_id}", handler)
	t.Run("book not found", func(t *testing.T) {
		bID := entity.NewID()
		uID := entity.NewID()
		bMock.EXPECT().GetBook(bID).Return(nil, domain.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/loan/borrow/%s/%s", ts.URL, bID.String(), uID.String()))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("user not found", func(t *testing.T) {
		b := entity.NewFixtureBook()
		uID := entity.NewID()
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		uMock.EXPECT().GetUser(uID).Return(nil, domain.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/loan/borrow/%s/%s", ts.URL, b.ID.String(), uID.String()))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		b := entity.NewFixtureBook()
		u := entity.NewFixtureUser()
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		uMock.EXPECT().GetUser(u.ID).Return(u, nil)
		lMock.EXPECT().Borrow(u, b).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/loan/borrow/%s/%s", ts.URL, b.ID.String(), u.ID.String()))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_returnBook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	uMock := umock.NewMockUseCase(controller)
	bMock := bmock.NewMockUseCase(controller)
	lMock := lmock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeLoanHandlers(r, *n, bMock, uMock, lMock)
	path, err := r.GetRoute("returnBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/loan/return/{book_id}", path)
	handler := returnBook(bMock, lMock)
	r.Handle("/v1/loan/return/{book_id}", handler)
	t.Run("book not found", func(t *testing.T) {
		bID := entity.NewID()
		bMock.EXPECT().GetBook(bID).Return(nil, domain.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/loan/return/%s", ts.URL, bID.String()))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		b := entity.NewFixtureBook()
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		lMock.EXPECT().Return(b).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/loan/return/%s", ts.URL, b.ID.String()))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
