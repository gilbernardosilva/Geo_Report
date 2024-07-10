package model

import "fmt"

type MockUserRepository struct {
	Users map[uint64]*User // Simulated storage
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Users: make(map[uint64]*User),
	}
}

func (m *MockUserRepository) CreateUser(user *User) error {
	// Simulate database behavior: insert user into mock storage
	m.Users[user.ID] = user
	return nil
}

func (m *MockUserRepository) UpdateUser(user *User) error {
	// Simulate database behavior: update user in mock storage
	if _, ok := m.Users[user.ID]; !ok {
		return fmt.Errorf("user with ID %d not found", user.ID)
	}
	m.Users[user.ID] = user
	return nil
}

func (m *MockUserRepository) DeleteUser(userID uint64) error {
	// Simulate database behavior: delete user from mock storage
	if _, ok := m.Users[userID]; !ok {
		return fmt.Errorf("user with ID %d not found", userID)
	}
	delete(m.Users, userID)
	return nil
}

func (m *MockUserRepository) GetUser(userID uint64) (*User, error) {
	// Simulate database behavior: retrieve user from mock storage
	user, ok := m.Users[userID]
	if !ok {
		return nil, fmt.Errorf("user with ID %d not found", userID)
	}
	return user, nil
}
