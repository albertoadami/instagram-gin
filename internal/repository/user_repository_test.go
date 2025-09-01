package repository

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/albertoadami/instagram-gin/internal/domain"
	"github.com/albertoadami/instagram-gin/internal/testutil"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/stretchr/testify/assert"
)

var (
	testContainer *testutil.PostgresContainer
	testRepo      *PostgresUserRepository
	testDB        *sqlx.DB
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Setup
	container, err := testutil.SetupPostgres(ctx, "../../migrations")
	if err != nil {
		fmt.Printf("Failed to setup test container: %v\n", err)
		os.Exit(1)
	}

	testContainer = container
	testDB = container.DB
	testRepo = NewPostgresUserRepository(testDB)

	// Run tests
	code := m.Run()

	// Teardown
	testContainer.Teardown(ctx)

	os.Exit(code)
}

func TestCreateUserSuccessfully(t *testing.T) {
	cleanupTables(t)
	user := testutil.CreateRandomUser()

	err := testRepo.Create(user)
	assert.NoError(t, err)

	fetchedUser, err := testRepo.FindByID(user.ID)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, fetchedUser.ID, user.ID)
	assert.Equal(t, fetchedUser.Username, user.Username)
	assert.Equal(t, fetchedUser.Email, user.Email)
	assert.Equal(t, fetchedUser.Name, user.Name)
	assert.Equal(t, fetchedUser.Surname, user.Surname)
	assert.Equal(t, fetchedUser.PasswordHash, user.PasswordHash)
	assert.Equal(t, fetchedUser.Gender, user.Gender)
	assert.Equal(t, fetchedUser.Status, user.Status)
	expectedDate := user.BirthDate.Format("2006-01-02")
	actualDate := fetchedUser.BirthDate.Format("2006-01-02")
	assert.Equal(t, expectedDate, actualDate)
	assert.WithinDuration(t, fetchedUser.CreatedAt, user.CreatedAt, time.Second)
	assert.WithinDuration(t, fetchedUser.UpdateAt, user.UpdateAt, time.Second)

}

func TestFindUserByIDNotFound(t *testing.T) {
	cleanupTables(t)

	user, err := testRepo.FindByID(uuid.New())
	assert.Nil(t, user)
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func TestFindUserByIDFound(t *testing.T) {
	cleanupTables(t)

	originalUser := testutil.CreateRandomUser()
	err := testRepo.Create(originalUser)
	assert.NoError(t, err)

	fetchedUser, err := testRepo.FindByID(originalUser.ID)

	assert.NoError(t, err)
	assert.Equal(t, originalUser.ID, fetchedUser.ID)
	assert.Equal(t, fetchedUser.ID, originalUser.ID)
	assert.Equal(t, fetchedUser.Username, originalUser.Username)
	assert.Equal(t, fetchedUser.Email, originalUser.Email)
	assert.Equal(t, fetchedUser.Name, originalUser.Name)
	assert.Equal(t, fetchedUser.Surname, originalUser.Surname)
	assert.Equal(t, fetchedUser.PasswordHash, originalUser.PasswordHash)
	assert.Equal(t, fetchedUser.Gender, originalUser.Gender)
	assert.Equal(t, fetchedUser.Status, originalUser.Status)
}

func TestUpdateUserSuccessfully(t *testing.T) {
	cleanupTables(t)

	originalUser := testutil.CreateRandomUser()
	err := testRepo.Create(originalUser)
	assert.NoError(t, err)

	// Modify some fields
	originalUser.Name = "UpdatedName"
	originalUser.Surname = "UpdatedSurname"
	originalUser.Status = domain.Inactive
	err = testRepo.Update(originalUser)
	assert.NoError(t, err)

	fetchedUser, err := testRepo.FindByID(originalUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, "UpdatedName", fetchedUser.Name)
	assert.Equal(t, "UpdatedSurname", fetchedUser.Surname)
	assert.Equal(t, domain.Inactive, fetchedUser.Status)

}

func TestUpdateUserNotFound(t *testing.T) {
	cleanupTables(t)
	nonExistentUser := testutil.CreateRandomUser() // Not inserted into DB

	err := testRepo.Update(nonExistentUser)
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func TestDeleteUserByIdNotFound(t *testing.T) {
	cleanupTables(t)
	err := testRepo.DeleteById(uuid.New())
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func TestDeleteUserByIdFound(t *testing.T) {
	cleanupTables(t)
	originalUser := testutil.CreateRandomUser()
	err := testRepo.Create(originalUser)
	assert.NoError(t, err)

	err = testRepo.DeleteById(originalUser.ID)
	assert.NoError(t, err)

	fetchedUser, err := testRepo.FindByID(originalUser.ID)
	assert.Nil(t, fetchedUser)
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func cleanupTables(t *testing.T) {
	t.Helper() // This marks the function as a test helper

	_, err := testDB.Exec("TRUNCATE TABLE users")
	if err != nil {
		t.Fatalf("Failed to truncate table %s: %v", "users", err)
	}
}
