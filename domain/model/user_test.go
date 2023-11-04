package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	user := NewUser(&User{Email: "", ID: 1, FirstName: "", LastName: "", Password: "", Phone: ""})

	t.Run("Email Cannot be Empty", func(t *testing.T) {
		customErr := user.Validate()
		err := fmt.Errorf(customErr.Message)

		assert.NotNil(t, customErr)

		assert.EqualError(t, err, "Email cannot be empty")

	})

	t.Run("Phone cannot be empty", func(t *testing.T) {
		user := NewUser(&User{Email: "ziadshimy@gmail.com", ID: 1, FirstName: "", LastName: "", Password: "asdasdasd", Phone: ""})
		customErr := user.Validate()

		err := fmt.Errorf(customErr.Message)

		assert.NotNil(t, customErr)

		assert.EqualError(t, err, "Phone cannot be empty")
	})

	t.Run("Password cannot be empty", func(t *testing.T) {
		user := NewUser(&User{Email: "ziadshimy@gmail.com", ID: 1, FirstName: "", LastName: "", Password: "", Phone: "asdasdasd"})
		customErr := user.Validate()

		err := fmt.Errorf(customErr.Message)

		assert.NotNil(t, customErr)

		assert.EqualError(t, err, "Password cannot be empty")
	})

	t.Run("First name cannot be empty", func(t *testing.T) {
		user := NewUser(&User{Email: "ziadshimy@gmail.com", ID: 1, FirstName: "", LastName: "asdasd", Password: "asdasd", Phone: "asdasdasd"})
		customErr := user.Validate()

		err := fmt.Errorf(customErr.Message)

		assert.NotNil(t, customErr)

		assert.EqualError(t, err, "First name cannot be empty")
	})

	t.Run("Last name cannot be empty", func(t *testing.T) {
		user := NewUser(&User{Email: "ziadshimy@gmail.com", ID: 1, FirstName: "asdasd", LastName: "", Password: "asdasd", Phone: "asdasdasd"})
		customErr := user.Validate()

		err := fmt.Errorf(customErr.Message)

		assert.NotNil(t, customErr)

		assert.EqualError(t, err, "Last name cannot be empty")
	})
}
