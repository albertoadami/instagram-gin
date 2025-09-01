package service

import (
	"time"

	"github.com/albertoadami/instagram-gin/internal/command"
	"github.com/albertoadami/instagram-gin/internal/domain"
	"github.com/albertoadami/instagram-gin/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (s *UserService) CreateUser(cmd *command.CreateUserCommand) (uuid.UUID, error) {
	user := domain.User{
		ID:           uuid.New(),
		Username:     cmd.Username,
		Email:        cmd.Email,
		Name:         cmd.Name,
		Surname:      cmd.Surname,
		PasswordHash: cmd.Password, // In a real application, hash the password
		Gender:       cmd.Gender,
		BirthDate:    cmd.BirthDate,
		Status:       domain.Pending,
		CreatedAt:    time.Now(),
		UpdateAt:     time.Now(),
	}
	if err := (s.userRepository).Create(&user); err != nil {
		return uuid.UUID{}, err
	}
	return user.ID, nil

}
