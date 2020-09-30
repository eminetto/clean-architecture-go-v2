package book

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/eminetto/clean-architecture-go-v2/domain/repository/book"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := book.NewInmemRepository()
	m := NewService(repo)
	u := entity.NewFixtureBook()
	id, err := m.CreateBook(u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := book.NewInmemRepository()
	m := NewService(repo)
	u1 := entity.NewFixtureBook()
	u2 := entity.NewFixtureBook()
	u2.Title = "Lemmy: Biography"

	uID, _ := m.CreateBook(u1)
	_, _ = m.CreateBook(u2)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchBooks("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "I Am Ozzy", c[0].Title)

		c, err = m.SearchBooks("dio")
		assert.Equal(t, domain.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListBooks()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetBook(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Title, saved.Title)
	})
}

func Test_Update(t *testing.T) {
	repo := book.NewInmemRepository()
	m := NewService(repo)
	u := entity.NewFixtureBook()
	id, err := m.CreateBook(u)
	assert.Nil(t, err)
	saved, _ := m.GetBook(id)
	saved.Title = "Lemmy: Biography"
	assert.Nil(t, m.UpdateBook(saved))
	updated, err := m.GetBook(id)
	assert.Nil(t, err)
	assert.Equal(t, "Lemmy: Biography", updated.Title)
}

func TestDelete(t *testing.T) {
	repo := book.NewInmemRepository()
	m := NewService(repo)
	u1 := entity.NewFixtureBook()
	u2 := entity.NewFixtureBook()
	u2ID, _ := m.CreateBook(u2)

	err := m.DeleteBook(u1.ID)
	assert.Equal(t, domain.ErrNotFound, err)

	err = m.DeleteBook(u2ID)
	assert.Nil(t, err)
	_, err = m.GetBook(u2ID)
	assert.Equal(t, domain.ErrNotFound, err)
}
