package services

import (
	"errors"

	"taskmanager/internal/config"
	"taskmanager/internal/models"
	"taskmanager/internal/repositories"
	"taskmanager/internal/utils"
)

var (
	ErrEmailTaken     = errors.New("an account with this email already exists")
	ErrInvalidCreds   = errors.New("invalid email or password")
)

type AuthService struct {
	userRepo *repositories.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo *repositories.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, cfg: cfg}
}

// Register creates a new user account and returns a signed JWT.
func (s *AuthService) Register(name, email, password string) (*models.User, string, error) {
	if existing, _ := s.userRepo.FindByEmail(email); existing != nil {
		return nil, "", ErrEmailTaken
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hash,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateJWT(user.ID, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// Login verifies credentials and returns a signed JWT on success.
func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", ErrInvalidCreds
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, "", ErrInvalidCreds
	}

	token, err := utils.GenerateJWT(user.ID, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
