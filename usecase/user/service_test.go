package user

import (
	"testing"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		Email:     "ozzy@metalgods.net",
		Password:  "123456",
		FirstName: "Ozzy",
		LastName:  "Osbourne",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureUser()
	_, err := m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.FirstName = "Lemmy"

	uID, _ := m.CreateUser(u1.Email, u1.Password, u1.FirstName, u1.LastName)
	_, _ = m.CreateUser(u2.Email, u2.Password, u2.FirstName, u2.LastName)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchUsers("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Osbourne", c[0].LastName)

		c, err = m.SearchUsers("dio")
		assert.Equal(t, entity.ErrNotFound, err)
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
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureUser()
	id, err := m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	assert.Nil(t, err)
	saved, _ := m.GetUser(id)
	saved.FirstName = "Dio"
	saved.Books = append(saved.Books, entity.NewID())
	assert.Nil(t, m.UpdateUser(saved))
	updated, err := m.GetUser(id)
	assert.Nil(t, err)
	assert.Equal(t, "Dio", updated.FirstName)
	assert.False(t, updated.UpdatedAt.IsZero())
	assert.Equal(t, 1, len(updated.Books))
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2ID, _ := m.CreateUser(u2.Email, u2.Password, u2.FirstName, u2.LastName)

	err := m.DeleteUser(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteUser(u2ID)
	assert.Nil(t, err)
	_, err = m.GetUser(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)

	u3 := newFixtureUser()
	id, _ := m.CreateUser(u3.Email, u3.Password, u3.FirstName, u3.LastName)
	saved, _ := m.GetUser(id)
	saved.Books = []entity.ID{entity.NewID()}
	_ = m.UpdateUser(saved)
	err = m.DeleteUser(id)
	assert.Equal(t, entity.ErrCannotBeDeleted, err)
}
