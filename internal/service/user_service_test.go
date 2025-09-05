package service

import (
	"testing"

	"github.com/albertoadami/instagram-gin/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserSucessfully(t *testing.T) {

	mockRepo := new(testutil.MockUserRepository)
	userService := NewUserService(mockRepo)

	cmd := testutil.CreateRandomCreateUserCommand()
	mockRepo.On("Create", mock.Anything).Return(nil)

	result, err := userService.CreateUser(cmd)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	assert.NotEqual(t, uuid.UUID{}, result)
}
