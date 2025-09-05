package handlers

import (
	"github.com/albertoadami/instagram-gin/internal/command"
	"github.com/albertoadami/instagram-gin/internal/dto"
	"github.com/albertoadami/instagram-gin/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userSvc service.UserService) *UserHandler {
	return &UserHandler{
		userService: userSvc,
	}
}
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	command := &command.CreateUserCommand{
		Username:  req.Username,
		Email:     req.Email,
		Name:      req.Name,
		Surname:   req.Surname,
		Password:  req.Password,
		Gender:    req.Gender,
		BirthDate: req.BirthDate,
	}
	id, err := h.userService.CreateUser(command)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, gin.H{"id": id})

}
