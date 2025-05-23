package tests

import (
	"context"
	"testing"

	"github.com/Helltale/tz-telecom/internal/domain"
	"github.com/Helltale/tz-telecom/internal/usecase"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	savedUser *domain.User
	saveErr   error
}

func (m *mockUserRepo) Save(ctx context.Context, u *domain.User) error {
	m.savedUser = u
	return m.saveErr
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name        string
		input       domain.User
		expectError string
	}{
		{
			name: "valid user",
			input: domain.User{
				FirstName: "Alice", LastName: "Smith", Age: 25, Password: "supersecure",
			},
			expectError: "",
		},
		{
			name: "too young",
			input: domain.User{
				FirstName: "Tom", LastName: "Young", Age: 17, Password: "12345678",
			},
			expectError: "user must be 18+",
		},
		{
			name: "short password",
			input: domain.User{
				FirstName: "Rick", LastName: "Tiny", Age: 30, Password: "123",
			},
			expectError: "password too short",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockUserRepo{}
			uc := usecase.NewUserUseCase(repo)
			err := uc.RegisterUser(context.Background(), &tt.input)

			if tt.expectError != "" {
				assert.EqualError(t, err, tt.expectError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo.savedUser)
				assert.NotEqual(t, "", repo.savedUser.Password) // must be hashing
			}
		})
	}
}
