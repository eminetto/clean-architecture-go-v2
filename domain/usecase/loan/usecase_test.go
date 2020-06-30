package loan

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"
	bmock "github.com/eminetto/clean-architecture-go-v2/domain/entity/book/mock"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"
	umock "github.com/eminetto/clean-architecture-go-v2/domain/entity/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Borrow(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	uMock := umock.NewMockManager(controller)
	bMock := bmock.NewMockManager(controller)
	service := NewUseCase(uMock, bMock)
	t.Run("user not found", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		uMock.EXPECT().Get(u.ID).Return(nil, domain.ErrNotFound)
		err := service.Borrow(u, b)
		assert.Equal(t, domain.ErrNotFound, err)
	})
	t.Run("book not found", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		uMock.EXPECT().Get(u.ID).Return(u, nil)
		bMock.EXPECT().Get(b.ID).Return(nil, domain.ErrNotFound)
		err := service.Borrow(u, b)
		assert.Equal(t, domain.ErrNotFound, err)
	})
	t.Run("not enough books to borrow", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		b.Quantity = 0
		uMock.EXPECT().Get(u.ID).Return(u, nil)
		bMock.EXPECT().Get(b.ID).Return(b, nil)
		err := service.Borrow(u, b)
		assert.Equal(t, domain.ErrNotEnoughBooks, err)
	})
	t.Run("book already borrowed", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		u.Books = []entity.ID{b.ID}
		b.Quantity = 1
		uMock.EXPECT().Get(u.ID).Return(u, nil)
		bMock.EXPECT().Get(b.ID).Return(b, nil)
		err := service.Borrow(u, b)
		assert.Equal(t, domain.ErrBookAlreadyBorrowed, err)
	})
	t.Run("sucess", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		uMock.EXPECT().Get(u.ID).Return(u, nil)
		bMock.EXPECT().Get(b.ID).Return(b, nil)
		uMock.EXPECT().Update(u).Return(nil)
		bMock.EXPECT().Update(b).Return(nil)
		err := service.Borrow(u, b)
		assert.Nil(t, err)
	})
}

func Test_Return(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	uMock := umock.NewMockManager(controller)
	bMock := bmock.NewMockManager(controller)
	service := NewUseCase(uMock, bMock)
	t.Run("book not found", func(t *testing.T) {
		b := book.NewFixtureBook()
		bMock.EXPECT().Get(b.ID).Return(nil, domain.ErrNotFound)
		err := service.Return(b)
		assert.Equal(t, domain.ErrNotFound, err)
	})
	t.Run("book not borrowed", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		bMock.EXPECT().Get(b.ID).Return(b, nil)
		uMock.EXPECT().List().Return([]*user.User{u}, nil)
		err := service.Return(b)
		assert.Equal(t, domain.ErrBookNotBorrowed, err)
	})
	t.Run("success", func(t *testing.T) {
		u := user.NewFixtureUser()
		b := book.NewFixtureBook()
		u.Books = []entity.ID{b.ID}
		bMock.EXPECT().Get(b.ID).Return(b, nil)
		uMock.EXPECT().Get(u.ID).Return(u, nil)
		uMock.EXPECT().List().Return([]*user.User{u}, nil)
		uMock.EXPECT().Update(u).Return(nil)
		bMock.EXPECT().Update(b).Return(nil)
		err := service.Return(b)
		assert.Nil(t, err)
	})
}
