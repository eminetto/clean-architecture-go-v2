package book

import (
	"github.com/eminetto/clean-architecture-go-v2/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	u := NewFixtureBook()
	id, err := service.Create(u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	u1 := NewFixtureBook()
	u2 := NewFixtureBook()
	u2.Title = "Lemmy: Biography"

	uID, _ := service.Create(u1)
	_, _ = service.Create(u2)

	t.Run("search", func(t *testing.T) {
		c, err := service.Search("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "I Am Ozzy", c[0].Title)

		c, err = service.Search("dio")
		assert.Equal(t, domain.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := service.List()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := service.Get(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Title, saved.Title)
	})
}

func Test_Update(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	u := NewFixtureBook()
	id, err := service.Create(u)
	assert.Nil(t, err)
	saved, _ := service.Get(id)
	saved.Title = "Lemmy: Biography"
	assert.Nil(t, service.Update(saved))
	updated, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, "Lemmy: Biography", updated.Title)
}

func TestDelete(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	u1 := NewFixtureBook()
	u2 := NewFixtureBook()
	u2ID, _ := service.Create(u2)

	err := service.Delete(u1.ID)
	assert.Equal(t, domain.ErrNotFound, err)

	err = service.Delete(u2ID)
	assert.Nil(t, err)
	_, err = service.Get(u2ID)
	assert.Equal(t, domain.ErrNotFound, err)
}
