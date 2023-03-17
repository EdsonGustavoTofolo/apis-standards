package database

import (
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Joh", "mail@mail.com", "12345")

	userDb := NewUser(db)

	err = userDb.Create(user)

	assert.NoError(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Joh", "mail@mail.com", "12345")

	userDb := NewUser(db)

	err = userDb.Create(user)

	assert.NoError(t, err)

	userFound, err := userDb.FindByEmail(user.Email)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}
