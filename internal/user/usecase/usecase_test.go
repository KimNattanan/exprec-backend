package usecase

import (
	"errors"
	"testing"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct {
	findByEmailFunc func(email string) (*entities.User, error)
	findByIDFunc    func(id string) (*entities.User, error)
	findAllFunc     func() ([]*entities.User, error)
	saveFunc        func(user *entities.User) error
	patchFunc       func(id string, user *entities.User) error
	deleteFunc      func(id string) error
}

func (m *mockUserRepo) FindByEmail(email string) (*entities.User, error) {
	return m.findByEmailFunc(email)
}
func (m *mockUserRepo) FindByID(id string) (*entities.User, error) {
	return m.findByIDFunc(id)
}
func (m *mockUserRepo) FindAll() ([]*entities.User, error) {
	return m.findAllFunc()
}
func (m *mockUserRepo) Save(user *entities.User) error {
	return m.saveFunc(user)
}
func (m *mockUserRepo) Patch(id string, user *entities.User) error {
	return m.patchFunc(id, user)
}
func (m *mockUserRepo) Delete(id string) error {
	return m.deleteFunc(id)
}

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockUserRepo{
			saveFunc: func(user *entities.User) error {
				return nil
			},
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
		}
		service := NewUserService(repo)

		err := service.Register(&entities.User{Email: "john@example.com", Password: "password123"})
		assert.NoError(t, err)
	})
	t.Run("save fails", func(t *testing.T) {
		repo := &mockUserRepo{
			saveFunc: func(user *entities.User) error {
				return errors.New("save failed")
			},
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
		}
		service := NewUserService(repo)

		err := service.Register(&entities.User{Email: "john@example.com", Password: "password123"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "save failed")
	})
	t.Run("duplicate email", func(t *testing.T) {
		repo := &mockUserRepo{
			saveFunc: func(user *entities.User) error {
				return nil
			},
			findByEmailFunc: func(email string) (*entities.User, error) {
				return &entities.User{ID: uuid.New(), Email: "john@example.com", Password: "password123"}, nil
			},
		}
		service := NewUserService(repo)

		err := service.Register(&entities.User{Email: "john@example.com", Password: "password123"})
		assert.ErrorIs(t, err, apperror.ErrAlreadyExists)
	})
	t.Run("findByEmail unknown error", func(t *testing.T) {
		repo := &mockUserRepo{
			saveFunc: func(user *entities.User) error {
				return nil
			},
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, errors.New("database error")
			},
		}
		service := NewUserService(repo)

		err := service.Register(&entities.User{Email: "john@example.com", Password: "password123"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
	})
}
func TestLogin(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &entities.User{ID: uuid.New(), Email: "john@example.com", Password: string(hashed)}

	t.Run("success", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return user, nil
			},
		}
		service := NewUserService(repo)

		token, u, err := service.Login("john@example.com", "password123")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, user.Email, u.Email)
	})

	t.Run("user not found", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
		}
		service := NewUserService(repo)

		token, u, err := service.Login("john@example.com", "password123")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, u)
	})

	t.Run("wrong password", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return user, nil
			},
		}
		service := NewUserService(repo)

		token, u, err := service.Login("john@example.com", "wrongpass")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, u)
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, errors.New("db error")
			},
		}
		service := NewUserService(repo)

		token, u, err := service.Login("john@example.com", "password123")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, u)
	})
}
func TestFindByEmail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &entities.User{ID: uuid.New(), Email: "john@example.com"}
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return user, nil
			},
		}
		service := NewUserService(repo)

		result, err := service.FindByEmail("john@example.com")
		assert.NoError(t, err)
		assert.Equal(t, user, result)
	})
	t.Run("not found", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
		}
		service := NewUserService(repo)

		result, err := service.FindByEmail("unknown@example.com")
		assert.ErrorIs(t, err, apperror.ErrRecordNotFound)
		assert.Nil(t, result)
	})
}
func TestFindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &entities.User{ID: uuid.New(), Email: "john@example.com"}
		repo := &mockUserRepo{
			findByIDFunc: func(id string) (*entities.User, error) {
				return user, nil
			},
		}
		service := NewUserService(repo)

		result, err := service.FindByID(user.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, user, result)
	})
	t.Run("not found", func(t *testing.T) {
		id := uuid.New()
		repo := &mockUserRepo{
			findByIDFunc: func(uid string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
		}
		service := NewUserService(repo)

		result, err := service.FindByID(id.String())
		assert.ErrorIs(t, err, apperror.ErrRecordNotFound)
		assert.Nil(t, result)
	})
}
func TestFindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		users := []*entities.User{
			{ID: uuid.New(), Email: "john@example.com"},
			{ID: uuid.New(), Email: "jane@example.com"},
		}
		repo := &mockUserRepo{
			findAllFunc: func() ([]*entities.User, error) {
				return users, nil
			},
		}
		service := NewUserService(repo)

		result, err := service.FindAll()
		assert.NoError(t, err)
		assert.Equal(t, users, result)
	})
	t.Run("find failed", func(t *testing.T) {
		repo := &mockUserRepo{
			findAllFunc: func() ([]*entities.User, error) {
				return nil, errors.New("db error")
			},
		}
		service := NewUserService(repo)

		result, err := service.FindAll()
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
func TestPatch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := uuid.New()
		user := &entities.User{ID: id, Email: "john@example.com"}
		repo := &mockUserRepo{
			patchFunc: func(uid string, u *entities.User) error {
				return nil
			},
			findByIDFunc: func(uid string) (*entities.User, error) {
				return user, nil
			},
		}
		service := NewUserService(repo)

		result, err := service.Patch(id.String(), user)
		assert.NoError(t, err)
		assert.Equal(t, user, result)
	})
	t.Run("patch fails", func(t *testing.T) {
		id := uuid.New()
		user := &entities.User{Email: "john@example.com"}
		repo := &mockUserRepo{
			patchFunc: func(uid string, u *entities.User) error {
				return errors.New("patch failed")
			},
		}
		service := NewUserService(repo)

		result, err := service.Patch(id.String(), user)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("findByID fails after patch", func(t *testing.T) {
		id := uuid.New()
		user := &entities.User{Email: "john@example.com"}
		repo := &mockUserRepo{
			patchFunc: func(uid string, u *entities.User) error {
				return nil
			},
			findByIDFunc: func(uid string) (*entities.User, error) {
				return nil, errors.New("find failed")
			},
		}
		service := NewUserService(repo)

		result, err := service.Patch(id.String(), user)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := uuid.New()
		repo := &mockUserRepo{
			deleteFunc: func(uid string) error {
				return nil
			},
		}
		service := NewUserService(repo)

		err := service.Delete(id.String())
		assert.NoError(t, err)
	})
	t.Run("delete fails", func(t *testing.T) {
		id := uuid.New()
		repo := &mockUserRepo{
			deleteFunc: func(uid string) error {
				return errors.New("delete failed")
			},
		}
		service := NewUserService(repo)

		err := service.Delete(id.String())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "delete failed")
	})
}
func TestLoginOrRegisterWithGoogle(t *testing.T) {
	t.Run("new user created", func(t *testing.T) {
		email := "john@example.com"
		name := "John Doe"
		var savedUser *entities.User

		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
			saveFunc: func(user *entities.User) error {
				savedUser = user
				user.ID = uuid.New()
				return nil
			},
		}
		service := NewUserService(repo)

		token, user, err := service.LoginOrRegisterWithGoogle(
			map[string]interface{}{"email": email, "name": name}, nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.NotNil(t, savedUser)
	})

	t.Run("existing user", func(t *testing.T) {
		u := &entities.User{ID: uuid.New(), Email: "john@example.com", Name: "John"}
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return u, nil
			},
			saveFunc: func(user *entities.User) error {
				return nil
			},
		}
		service := NewUserService(repo)

		token, user, err := service.LoginOrRegisterWithGoogle(
			map[string]interface{}{"email": u.Email, "name": u.Name}, nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, u.Email, user.Email)
	})
	t.Run("missing email", func(t *testing.T) {
		repo := &mockUserRepo{}
		service := NewUserService(repo)

		token, user, err := service.LoginOrRegisterWithGoogle(
			map[string]interface{}{"name": "John"}, nil)
		assert.ErrorIs(t, err, apperror.ErrInvalidData)
		assert.Empty(t, token)
		assert.Nil(t, user)
	})

	t.Run("repo save fails", func(t *testing.T) {
		email := "john@example.com"
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, apperror.ErrRecordNotFound
			},
			saveFunc: func(user *entities.User) error {
				return errors.New("save failed")
			},
		}
		service := NewUserService(repo)

		token, user, err := service.LoginOrRegisterWithGoogle(
			map[string]interface{}{"email": email, "name": "John"}, nil)
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
	})

	t.Run("repo findByEmail unknown error", func(t *testing.T) {
		email := "john@example.com"
		repo := &mockUserRepo{
			findByEmailFunc: func(email string) (*entities.User, error) {
				return nil, errors.New("db error")
			},
		}
		service := NewUserService(repo)

		token, user, err := service.LoginOrRegisterWithGoogle(
			map[string]interface{}{"email": email, "name": "John"}, nil)
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
	})
}
