package repository

import (
	"database/sql"
	"time"

	"github.com/albertoadami/instagram-gin/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) error
	DeleteById(id uuid.UUID) error
}

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	query := `
        INSERT INTO users (id, username, email, name, surname, password_hash, gender, birth_date, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Name,
		user.Surname,
		user.PasswordHash,
		user.Gender,
		user.BirthDate,
		user.Status,
		user.CreatedAt,
		user.UpdateAt,
	)

	return err
}

func (r *PostgresUserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	query := `
        SELECT id, username, email, name, surname, password_hash, gender, birth_date, created_at, updated_at, status
        FROM users 
        WHERE id = $1`

	var user domain.User
	err := r.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) Update(user *domain.User) error {
	user.UpdateAt = time.Now()

	query := `
        UPDATE users 
        SET username = $2, email = $3, name = $4, surname = $5, password_hash = $6, 
            gender = $7, birth_date = $8, status = $9, updated_at = $10
        WHERE id = $1`

	result, err := r.db.Exec(
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Name,
		user.Surname,
		user.PasswordHash,
		user.Gender,
		user.BirthDate,
		user.Status,
		user.UpdateAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *PostgresUserRepository) DeleteById(id uuid.UUID) error {
	query := `
        DELETE FROM users 
        WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil

}
