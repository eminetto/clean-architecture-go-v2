package user

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/pkg/password"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo, password.NewFakeService())
	u := NewFixtureUser()
	id, err := service.Create(u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo, password.NewFakeService())
	u1 := NewFixtureUser()
	u2 := NewFixtureUser()
	u2.FirstName = "Lemmy"

	uID, _ := service.Create(u1)
	_, _ = service.Create(u2)

	t.Run("search", func(t *testing.T) {
		c, err := service.Search("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Osbourne", c[0].LastName)

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
		assert.Equal(t, u1.FirstName, saved.FirstName)
	})
}

func Test_Update(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo, password.NewFakeService())
	u := NewFixtureUser()
	id, err := service.Create(u)
	assert.Nil(t, err)
	saved, _ := service.Get(id)
	saved.FirstName = "Dio"
	saved.Books = append(saved.Books, entity.NewID())
	assert.Nil(t, service.Update(saved))
	updated, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, "Dio", updated.FirstName)
	assert.False(t, u.UpdatedAt.IsZero())
	assert.Equal(t, 1, len(updated.Books))
}

func TestDelete(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo, password.NewFakeService())
	u1 := NewFixtureUser()
	u2 := NewFixtureUser()
	u2ID, _ := service.Create(u2)

	err := service.Delete(u1.ID)
	assert.Equal(t, domain.ErrNotFound, err)

	err = service.Delete(u2ID)
	assert.Nil(t, err)
	_, err = service.Get(u2ID)
	assert.Equal(t, domain.ErrNotFound, err)
}
