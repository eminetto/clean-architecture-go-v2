package loan

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	bmock "github.com/eminetto/clean-architecture-go-v2/domain/usecase/book/mock"
	umock "github.com/eminetto/clean-architecture-go-v2/domain/usecase/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Borrow(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	uMock := umock.NewMockUseCase(controller)
	bMock := bmock.NewMockUseCase(controller)
	uc := NewService(uMock, bMock)
	t.Run("user not found", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		uMock.EXPECT().GetUser(u.ID).Return(nil, domain.ErrNotFound)
		err := uc.Borrow(u, b)
		assert.Equal(t, domain.ErrNotFound, err)
	})
	t.Run("book not found", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		uMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bMock.EXPECT().GetBook(b.ID).Return(nil, domain.ErrNotFound)
		err := uc.Borrow(u, b)
		assert.Equal(t, domain.ErrNotFound, err)
	})
	t.Run("not enough books to borrow", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		b.Quantity = 0
		uMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		err := uc.Borrow(u, b)
		assert.Equal(t, domain.ErrNotEnoughBooks, err)
	})
	t.Run("book already borrowed", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		u.Books = []entity.ID{b.ID}
		b.Quantity = 1
		uMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		err := uc.Borrow(u, b)
		assert.Equal(t, domain.ErrBookAlreadyBorrowed, err)
	})
	t.Run("sucess", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		uMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		uMock.EXPECT().UpdateUser(u).Return(nil)
		bMock.EXPECT().UpdateBook(b).Return(nil)
		err := uc.Borrow(u, b)
		assert.Nil(t, err)
	})
}

func Test_Return(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	uMock := umock.NewMockUseCase(controller)
	bMock := bmock.NewMockUseCase(controller)
	uc := NewService(uMock, bMock)
	t.Run("book not found", func(t *testing.T) {
		b := entity.NewFixtureBook()
		bMock.EXPECT().GetBook(b.ID).Return(nil, domain.ErrNotFound)
		err := uc.Return(b)
		assert.Equal(t, domain.ErrNotFound, err)
	})
	t.Run("book not borrowed", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		uMock.EXPECT().ListUsers().Return([]*entity.User{u}, nil)
		err := uc.Return(b)
		assert.Equal(t, domain.ErrBookNotBorrowed, err)
	})
	t.Run("success", func(t *testing.T) {
		u := entity.NewFixtureUser()
		b := entity.NewFixtureBook()
		u.Books = []entity.ID{b.ID}
		bMock.EXPECT().GetBook(b.ID).Return(b, nil)
		uMock.EXPECT().GetUser(u.ID).Return(u, nil)
		uMock.EXPECT().ListUsers().Return([]*entity.User{u}, nil)
		uMock.EXPECT().UpdateUser(u).Return(nil)
		bMock.EXPECT().UpdateBook(b).Return(nil)
		err := uc.Return(b)
		assert.Nil(t, err)
	})
}
