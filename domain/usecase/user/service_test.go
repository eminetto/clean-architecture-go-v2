package user

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/domain/repository/user"

	"github.com/eminetto/clean-architecture-go-v2/pkg/password"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := user.NewInmemRepository()
	m := NewService(repo, password.NewFakeService())
	u := entity.NewFixtureUser()
	id, err := m.CreateUser(u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := user.NewInmemRepository()
	m := NewService(repo, password.NewFakeService())
	u1 := entity.NewFixtureUser()
	u2 := entity.NewFixtureUser()
	u2.FirstName = "Lemmy"

	uID, _ := m.CreateUser(u1)
	_, _ = m.CreateUser(u2)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchUsers("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Osbourne", c[0].LastName)

		c, err = m.SearchUsers("dio")
		assert.Equal(t, domain.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListUsers()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetUser(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.FirstName, saved.FirstName)
	})
}

func Test_Update(t *testing.T) {
	repo := user.NewInmemRepository()
	m := NewService(repo, password.NewFakeService())
	u := entity.NewFixtureUser()
	id, err := m.CreateUser(u)
	assert.Nil(t, err)
	saved, _ := m.GetUser(id)
	saved.FirstName = "Dio"
	saved.Books = append(saved.Books, entity.NewID())
	assert.Nil(t, m.UpdateUser(saved))
	updated, err := m.GetUser(id)
	assert.Nil(t, err)
	assert.Equal(t, "Dio", updated.FirstName)
	assert.False(t, u.UpdatedAt.IsZero())
	assert.Equal(t, 1, len(updated.Books))
}

func TestDelete(t *testing.T) {
	repo := user.NewInmemRepository()
	m := NewService(repo, password.NewFakeService())
	u1 := entity.NewFixtureUser()
	u2 := entity.NewFixtureUser()
	u2ID, _ := m.CreateUser(u2)

	err := m.DeleteUser(u1.ID)
	assert.Equal(t, domain.ErrNotFound, err)

	err = m.DeleteUser(u2ID)
	assert.Nil(t, err)
	_, err = m.GetUser(u2ID)
	assert.Equal(t, domain.ErrNotFound, err)

	u3 := entity.NewFixtureUser()
	u3.Books = []entity.ID{entity.NewID()}
	_, _ = m.CreateUser(u3)
	err = m.DeleteUser(u3.ID)
	assert.Equal(t, domain.ErrCannotBeDeleted, err)
}
