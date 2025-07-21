package service

import (
	"errors"
	"study-go-controller/internal/domain/user/entity"
	"study-go-controller/internal/domain/user/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService defines the contract for user business logic
type UserService interface {
	CreateUser(username, email, password, name string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	UpdateUser(id uint, username, email, name string) (*entity.User, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]*entity.User, error)
	ValidatePassword(password, hashedPassword string) bool
}

// userService implements UserService interface
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user with hashed password
func (s *userService) CreateUser(username, email, password, name string) (*entity.User, error) {
	// Check if user already exists
	if existingUser, _ := s.userRepo.GetByEmail(email); existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	if existingUser, _ := s.userRepo.GetByUsername(username); existingUser != nil {
		return nil, errors.New("user with this username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id uint) (*entity.User, error) {
	return s.userRepo.GetByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(email string) (*entity.User, error) {
	return s.userRepo.GetByEmail(email)
}

// GetUserByUsername retrieves a user by username
func (s *userService) GetUserByUsername(username string) (*entity.User, error) {
	return s.userRepo.GetByUsername(username)
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id uint, username, email, name string) (*entity.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if email is already taken by another user
	if existingUser, _ := s.userRepo.GetByEmail(email); existingUser != nil && existingUser.ID != id {
		return nil, errors.New("email is already taken")
	}

	// Check if username is already taken by another user
	if existingUser, _ := s.userRepo.GetByUsername(username); existingUser != nil && existingUser.ID != id {
		return nil, errors.New("username is already taken")
	}

	user.Username = username
	user.Email = email
	user.Name = name

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

// GetAllUsers retrieves all users
func (s *userService) GetAllUsers() ([]*entity.User, error) {
	return s.userRepo.GetAll()
}

// ValidatePassword validates if the provided password matches the hashed password
func (s *userService) ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
