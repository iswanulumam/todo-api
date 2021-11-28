package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	todoTitle = "Eat banana"
)

func TestSetup(t *testing.T) {

	authRepository := New()

	authLogin(t, authRepository)
}

func authLogin(t *testing.T, repository *AuthRepository) {
	t.Run("AuthLogin", func(t *testing.T) {
		token, _ := repository.Login("admin", "admin")
		assert.Equal(t, len(token), 159)
	})
}
