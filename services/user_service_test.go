package services

import (
	"context"
	"errors"
	"project-fa/mocks"
	"project-fa/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewUserService(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Test Initiate New User Service", func(t *testing.T) {
		mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)
		userService := NewUserService(mockUserRepo)
		require.NotNil(t, userService)
	})
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Success Create User", func(t *testing.T) {
		mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)
		userService := NewUserService(mockUserRepo)

		user := models.CreateUserRequest{
			Username:    "usamah",
			Email:       "usamah@gmail.com",
			Password:    "usamah",
			PhoneNumber: "1234",
			Age:         30,
		}

		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

		err := userService.CreateUser(context.Background(), user)

		require.Nil(t, err)
	})

	t.Run("Failed, Error: Username is Required", func(t *testing.T) {
		mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)
		userService := NewUserService(mockUserRepo)

		user := models.CreateUserRequest{
			Username:    "",
			Email:       "usamah@gmail.com",
			Password:    "usamah",
			PhoneNumber: "1234",
			Age:         30,
		}

		err := userService.CreateUser(context.Background(), user)

		require.Equal(t, errors.New("username is required"), err)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Success Update User", func(t *testing.T) {
		mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)
		userService := NewUserService(mockUserRepo)

		updateUserRequest := models.UpdateUserRequest{
			Username: "usamah",
			Email:    "usamah@gmail.com",
		}

		idToken := 1

		getUser := models.User{
			Id:          1,
			Username:    "abdurrahman",
			Email:       "abdurrahman@gmail.com",
			Password:    "abdurrahman",
			PhoneNumber: "1234",
			Age:         30,
		}

		updateUser := models.User{
			Id:          1,
			Username:    "usamah",
			Email:       "usamah@gmail.com",
			Password:    "abdurrahman",
			PhoneNumber: "1234",
			Age:         30,
		}

		mockUserRepo.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(getUser, nil)

		mockUserRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(updateUser, nil)

		res, err := userService.UpdateUser(context.Background(), updateUserRequest, idToken)

		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, "usamah", res.Username)
		require.NotEqual(t, "age", 0)
	})
}
