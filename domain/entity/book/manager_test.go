package book

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureBook()
	id, err := m.Create(u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u1 := NewFixtureBook()
	u2 := NewFixtureBook()
	u2.Title = "Lemmy: Biography"

	uID, _ := m.Create(u1)
	_, _ = m.Create(u2)

	t.Run("search", func(t *testing.T) {
		c, err := m.Search("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "I Am Ozzy", c[0].Title)

		c, err = m.Search("dio")
		assert.Equal(t, domain.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.List()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.Get(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Title, saved.Title)
	})
}

func Test_Update(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureBook()
	id, err := m.Create(u)
	assert.Nil(t, err)
	saved, _ := m.Get(id)
	saved.Title = "Lemmy: Biography"
	assert.Nil(t, m.Update(saved))
	updated, err := m.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, "Lemmy: Biography", updated.Title)
}

func TestDelete(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u1 := NewFixtureBook()
	u2 := NewFixtureBook()
	u2ID, _ := m.Create(u2)

	err := m.Delete(u1.ID)
	assert.Equal(t, domain.ErrNotFound, err)

	err = m.Delete(u2ID)
	assert.Nil(t, err)
	_, err = m.Get(u2ID)
	assert.Equal(t, domain.ErrNotFound, err)
}
