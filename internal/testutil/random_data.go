package testutil

import (
	"time"

	"github.com/albertoadami/instagram-gin/internal/command"
	"github.com/albertoadami/instagram-gin/internal/domain"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func init() {
	gofakeit.Seed(0) // 0 uses current time as seed
}

func CreateRandomUser() *domain.User {
	return &domain.User{
		ID:           uuid.New(),
		Username:     gofakeit.Username(),
		Email:        gofakeit.Email(),
		Name:         gofakeit.FirstName(),
		Surname:      gofakeit.LastName(),
		PasswordHash: gofakeit.Password(true, true, true, true, false, 60),
		Gender:       domain.Male,
		BirthDate:    gofakeit.DateRange(time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2005, 12, 31, 0, 0, 0, 0, time.UTC)),
		Status:       domain.Active,
		CreatedAt:    gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		UpdateAt:     gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
	}
}

func CreateRandomCreateUserCommand() *command.CreateUserCommand {
	return &command.CreateUserCommand{
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Name:      gofakeit.FirstName(),
		Surname:   gofakeit.LastName(),
		Password:  gofakeit.Password(true, true, true, true, false, 60),
		Gender:    domain.Male,
		BirthDate: gofakeit.DateRange(time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2005, 12, 31, 0, 0, 0, 0, time.UTC)),
	}
}
