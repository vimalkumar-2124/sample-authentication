package routes

import (
	"log"
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
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	v := validator.New()
	if err := v.Struct(body); err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	token, err := u.UserService.SignIn(body)
	if err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"statusCode": 200, "message": "Signed In successfully", "jwtToken": token})
	return
}

func (u *UserRoutes) SignUp(c *gin.Context) {
	var body models.SignUpBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// token, err := u.UserService.SignUp(body)
	// if err != nil {
	// 	c.JSON(400, gin.H{"message": err.Error()})
	// 	return
	// }
	// c.JSON(200, gin.H{"message": "Signed Up", "token": token})

	result, err := u.UserService.SignUp(body)
	if err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	// c.JSON(200, gin.H{"message": "Signed Up", "token": token})
	c.JSON(201, gin.H{"statusCode": 201, "message": result})
	return

}

func (u *UserRoutes) LogOut(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	authToken = strings.ReplaceAll(authToken, "Bearer ", "")
	err := u.UserService.LogOut(authToken)
	if err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"statusCode": 200, "message": "Logged Out"})
}

// func (u *UserRoutes) ChangePassword(c *gin.Context) {
// 	var body models.ChangeUserPassword
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(400, gin.H{"message": err.Error()})
// 		return
// 	}
// 	err := u.UserService.ChangePassword(body)
// 	if err != nil {
// 		c.JSON(400, gin.H{"message": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "Password changed successfully"})
// 	return
// }

func (u *UserRoutes) ChangePassword(c *gin.Context) {
	id := c.Param("id")
	// log.Println("ID : ", id)
	var body models.ChangeUserPassword
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	err := u.UserService.ChangePassword(body, id)
	if err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"statusCode": 200, "message": "Password changed successfully"})
	return
}

func (u *UserRoutes) AllUsers(c *gin.Context) {
	log.Println("All user route started...")
	allUser, err := u.UserService.AllUser()
	if err != nil {
		c.JSON(400, gin.H{"statusCode": 400, "message": err.Error()})
		return
	}
	log.Println("All user route completed...")
	c.JSON(200, gin.H{"statusCode": 200, "message": "List of users", "data": allUser})
	return
}
