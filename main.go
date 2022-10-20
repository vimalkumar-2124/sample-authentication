package main

import (
	"context"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vimalkumar-2124/sample-authentication/db"
	"github.com/vimalkumar-2124/sample-authentication/repositories"
	"github.com/vimalkumar-2124/sample-authentication/routes"
	"github.com/vimalkumar-2124/sample-authentication/services"
)

// Middleware function
func ValidateAuth(userRepo repositories.UserRepo) gin.HandlerFunc {

	// TODO : Yet to implemet JWT auth
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.AbortWithStatus(403)
			return
		}
		authToken = strings.ReplaceAll(authToken, "Bearer ", "")
		found, session, err := userRepo.GetSessinById(authToken)
		if err != nil {
			log.Println("Session is not found ", err)
			c.AbortWithStatus(403)
			return
		}
		if !found {
			log.Println("Token not found")
			c.AbortWithStatus(403)
			return
		}
		c.Set("session", session)
		c.Next()

	}
}

func main() {
	log.Println("Starting...")
	dbName := "users"
	client, err := db.CreateDbConnection(dbName)
	if err != nil {
		log.Println("Failed to connect DB")
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	db := client.Database(dbName)

	// Repositories
	userRepo := repositories.NewInstanceOfUserRepo(db)

	// Services
	userService := services.NewInstanceOfUserService(userRepo)

	// Handlers
	userHandler := routes.NewInstanceOfUserRoutes(userService)

	router := gin.Default()

	userApI := router.Group("/user")
	{
		userApI.POST("/signin", userHandler.SignIn)
		userApI.POST("/signup", userHandler.SignUp)
		userApI.POST("/change-password", userHandler.ChangePassword)
		userApI.POST("/logout", ValidateAuth(userRepo), userHandler.LogOut)
	}

	router.Run(":8000")

}
