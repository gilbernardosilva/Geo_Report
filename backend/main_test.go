package main_test

import (
	"geo_report_api/model"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserStruct(t *testing.T) {
	expected := model.User{
		ID:           0,
		RoleID:       0,
		Role:         model.Role{},
		FirstName:    "",
		LastName:     "",
		UserName:     "",
		Password:     "",
		Email:        "",
		CreationDate: time.Time{},
	}
	user := model.User{}
	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Expected user to be %+v, but got %+v", expected, user)
	}
}

func TestCreateUser(t *testing.T) {
	mockRepo := model.NewMockUserRepository()
	user := &model.User{
		ID:           1,
		RoleID:       1,
		Role:         model.Role{},
		FirstName:    "John",
		LastName:     "Doe",
		UserName:     "johndoe",
		Password:     "password",
		Email:        "john.doe2@example.com",
		CreationDate: time.Now(),
	}
	err := mockRepo.CreateUser(user)
	assert.NoError(t, err, "CreateUser should not return an error")
	createdUser := mockRepo.Users[user.ID]
	assert.NotNil(t, createdUser, "Created user should not be nil")
	assert.Equal(t, user.ID, createdUser.ID, "User ID should match")
	assert.Equal(t, user.RoleID, createdUser.RoleID, "User RoleID should match")
	assert.Equal(t, user.FirstName, createdUser.FirstName, "User FirstName should match")
	assert.Equal(t, user.LastName, createdUser.LastName, "User LastName should match")
	assert.Equal(t, user.UserName, createdUser.UserName, "User UserName should match")
	assert.Equal(t, user.Password, createdUser.Password, "User Password should match")
	assert.Equal(t, user.Email, createdUser.Email, "User Email should match")
	assert.Equal(t, user.CreationDate.Unix(), createdUser.CreationDate.Unix(), "User CreationDate should match")
}

func TestUpdateUser(t *testing.T) {
	mockRepo := model.NewMockUserRepository()
	initialUser := &model.User{
		ID:           1,
		RoleID:       1,
		Role:         model.Role{},
		FirstName:    "John",
		LastName:     "Doe",
		UserName:     "johndoe",
		Password:     "password",
		Email:        "john.doe3@example.com",
		CreationDate: time.Now(),
	}
	mockRepo.Users[initialUser.ID] = initialUser
	updatedUser := &model.User{
		ID:           initialUser.ID,
		RoleID:       initialUser.RoleID,
		Role:         initialUser.Role,
		FirstName:    "Updated John",
		LastName:     "Updated Doe",
		UserName:     "updatedjohndoe",
		Password:     "updatedpassword",
		Email:        "updated.john.doe@example.com",
		CreationDate: initialUser.CreationDate,
	}
	err := mockRepo.UpdateUser(updatedUser)
	assert.NoError(t, err, "UpdateUser should not return an error")
	updatedUserFromRepo := mockRepo.Users[initialUser.ID]
	assert.NotNil(t, updatedUserFromRepo, "Updated user should not be nil")
	assert.Equal(t, updatedUser.ID, updatedUserFromRepo.ID, "User ID should match after update")
	assert.Equal(t, updatedUser.FirstName, updatedUserFromRepo.FirstName, "User FirstName should match after update")
	assert.Equal(t, updatedUser.LastName, updatedUserFromRepo.LastName, "User LastName should match after update")
	assert.Equal(t, updatedUser.UserName, updatedUserFromRepo.UserName, "User UserName should match after update")
	assert.Equal(t, updatedUser.Password, updatedUserFromRepo.Password, "User Password should match after update")
	assert.Equal(t, updatedUser.Email, updatedUserFromRepo.Email, "User Email should match after update")
	assert.Equal(t, updatedUser.CreationDate.Unix(), updatedUserFromRepo.CreationDate.Unix(), "User CreationDate should match after update")
}

func TestGetUser(t *testing.T) {
	mockRepo := model.NewMockUserRepository()
	expectedUser := &model.User{
		ID:           1,
		RoleID:       1,
		Role:         model.Role{},
		FirstName:    "John",
		LastName:     "Doe",
		UserName:     "johndoe",
		Password:     "password",
		Email:        "john.doe@example.com",
		CreationDate: time.Now(),
	}
	mockRepo.Users[expectedUser.ID] = expectedUser
	user, err := mockRepo.GetUser(expectedUser.ID)
	assert.NoError(t, err, "GetUser should not return an error")
	assert.NotNil(t, user, "Retrieved user should not be nil")
	assert.Equal(t, expectedUser.ID, user.ID, "User ID should match")
	assert.Equal(t, expectedUser.FirstName, user.FirstName, "User FirstName should match")
	assert.Equal(t, expectedUser.LastName, user.LastName, "User LastName should match")
	assert.Equal(t, expectedUser.UserName, user.UserName, "User UserName should match")
	assert.Equal(t, expectedUser.Password, user.Password, "User Password should match")
	assert.Equal(t, expectedUser.Email, user.Email, "User Email should match")
	assert.Equal(t, expectedUser.CreationDate.Unix(), user.CreationDate.Unix(), "User CreationDate should match")
}

func TestDeleteUser(t *testing.T) {
	mockRepo := model.NewMockUserRepository()
	userToDelete := &model.User{
		ID:           1,
		RoleID:       1,
		Role:         model.Role{},
		FirstName:    "John",
		LastName:     "Doe",
		UserName:     "johndoe",
		Password:     "password",
		Email:        "john.doe@example.com",
		CreationDate: time.Now(),
	}
	mockRepo.Users[userToDelete.ID] = userToDelete
	err := mockRepo.DeleteUser(userToDelete.ID)
	assert.NoError(t, err, "DeleteUser should not return an error")
	deletedUser := mockRepo.Users[userToDelete.ID]
	assert.Nil(t, deletedUser, "Deleted user should be nil after deletion")
}

func TestUserNotFound(t *testing.T) {
	mockRepo := model.NewMockUserRepository()
	initialUser := &model.User{
		ID:        1,
		FirstName: "John Doe",
		Email:     "john.doe@example.com",
	}
	mockRepo.CreateUser(initialUser)
	nonExistentUser := &model.User{
		ID:        2,
		FirstName: "Jane Doe",
		Email:     "jane.doe@example.com",
	}
	err := mockRepo.UpdateUser(nonExistentUser)
	assert.Error(t, err, "Expected UpdateUser to return an error for non-existent user")
	_, err = mockRepo.GetUser(2)
	assert.Error(t, err, "Expected GetUser to return an error for non-existent user")
}

func TestPasswordHashingAndVerification(t *testing.T) {
	user := &model.User{
		ID:        1,
		FirstName: "John Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
	}
	hashedPassword, err := model.Hash(user.Password)
	assert.NoError(t, err, "Hashing password should not return an error")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")
	user.Password = hashedPassword
	passwordCorrect := model.VerifyPassword(user.Password, "password123")
	assert.True(t, passwordCorrect, "Password verification should succeed for correct password")
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(10)
	}
}

func square(x int) int {
	return x * x
}

func TestFuzzySquare(t *testing.T) {
	for i := 0; i < 100; i++ {
		input := rand.Intn(201) - 100
		expected := input * input
		result := square(input)
		if result != expected {
			t.Errorf("For input %d, expected %d but got %d", input, expected, result)
		}
	}
}
