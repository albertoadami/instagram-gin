package dto

import (
	"time"

	"github.com/albertoadami/instagram-gin/internal/domain"
)

type CreateUserRequest struct {
	Username  string        `json:"username" binding:"required,min=3,max=50"`
	Email     string        `json:"email" binding:"required,email"`
	Name      string        `json:"name" binding:"required,min=1,max=100"`
	Surname   string        `json:"surname" binding:"required,min=1,max=100"`
	Password  string        `json:"password" binding:"required,min=8,max=72"`
	Gender    domain.Gender `json:"gender" binding:"required,oneof=male female"`
	BirthDate time.Time     `json:"birth_date" binding:"required"` // Format: YYYY-MM-DD
}
