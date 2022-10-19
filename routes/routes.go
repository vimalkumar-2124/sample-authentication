package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vimalkumar-2124/sample-authentication/models"
	"github.com/vimalkumar-2124/sample-authentication/services"
)

type UserRoutes struct {
	UserService services.UserService
}

func NewInstanceOfUserRoutes(userService services.UserService) *UserRoutes {
	return &UserRoutes{UserService: userService}
}

func (u *UserRoutes) SignIn(c *gin.Context) {
	var body models.SignInBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	v := validator.New()
	if err := v.Struct(body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	token, err := u.UserService.SignIn(body)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}
	c.JSON(200, gin.H{"message": "Signed In successfully", "token": token})
}

func (u *UserRoutes) SignUp(c *gin.Context) {
	var body models.SignUpBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	token, err := u.UserService.SignUp(body)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Signed Up", "token": token})

}

func (u *UserRoutes) LogOut(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	authToken = strings.ReplaceAll(authToken, "Bearer ", "")
	err := u.UserService.LogOut(authToken)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Logged Out"})
}
