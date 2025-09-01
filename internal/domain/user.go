package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `db:"id"`
	Username     string     `db:"username"`
	Email        string     `db:"email"`
	Name         string     `db:"name"`
	Surname      string     `db:"surname"`
	PasswordHash string     `db:"password_hash"`
	Gender       Gender     `db:"gender"`
	BirthDate    time.Time  `db:"birth_date"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdateAt     time.Time  `db:"updated_at"`
	Status       UserStatus `db:"status"`
}
