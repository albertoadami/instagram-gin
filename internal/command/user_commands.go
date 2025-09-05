package command

import (
	"time"

	"github.com/albertoadami/instagram-gin/internal/domain"
)

type CreateUserCommand struct {
	Username  string
	Email     string
	Name      string
	Surname   string
	Password  string
	Gender    domain.Gender
	BirthDate time.Time
}
