package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vimalkumar-2124/sample-authentication/db"
	"github.com/vimalkumar-2124/sample-authentication/repositories"
	"github.com/vimalkumar-2124/sample-authentication/routes"
	"github.com/vimalkumar-2124/sample-authentication/services"
	"github.com/vimalkumar-2124/sample-authentication/tokens"
)

// Middleware functions

func TokenValidity(userRepo repositories.UserRepo) gin.HandlerFunc {
	log.Println("Token Validity")
	return func(c *gin.Context) {
		token := tokens.ExtractToken(c.Request)
		found, _, err := userRepo.GetSessinById(token)
		if err != nil {
			c.JSON(403, gin.H{"message": err.Error()})
			c.AbortWithStatus(403)
			return
		}
		if !found {
			c.JSON(403, gin.H{"message": "Session is not found"})
			c.AbortWithStatus(403)
			return
		}
		_, err = tokens.ExtractTokenMetaData(c.Request)
		if err != nil {
			err := userRepo.MarkSessionAsExpired(token)
			if err != nil {
				c.JSON(403, gin.H{"message": err.Error()})
				c.AbortWithStatus(403)
				return
			}
			c.JSON(403, gin.H{"message": err.Error()})
			c.AbortWithStatus(403)
			return
		} else {
			c.Next()
		}
	}
}

func AdminGuard() gin.HandlerFunc {
	log.Println("Admin Guard")
	return func(c *gin.Context) {
		role, err := tokens.ExtractTokenMetaData(c.Request)
		if err != nil {
			c.JSON(403, gin.H{"message": err.Error()})
			c.AbortWithStatus(403)
			return
		}
		if role.Role == "admin" || role.Role == "Admin" {

			c.Next()
		} else {
			c.JSON(403, gin.H{"message": "You're not authorized"})
			c.AbortWithStatus(403)
			return
		}

	}
}

func ValidateAuth(userRepo repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := tokens.ExtractToken(c.Request)
		found, session, err := userRepo.GetSessinById(authToken)
		if err != nil {
			c.JSON(403, gin.H{"message": err.Error()})
			c.AbortWithStatus(403)
			return
		}
		if !found {
			log.Println("Token not found")
			c.JSON(403, gin.H{"message": "Token is not found"})
			c.AbortWithStatus(403)
			return
		}
		c.Set("session", session)
		c.Next()

	}
}

// To avoid CORS error
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-control-allow-origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
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
	router.Use(CORSMiddleware())
	userApI := router.Group("/users")
	{
		userApI.GET("/all", AdminGuard(), TokenValidity(userRepo), userHandler.AllUsers)
		userApI.POST("/signin", userHandler.SignIn)
		userApI.POST("/signup", userHandler.SignUp)
		userApI.POST("/change-password/:id", userHandler.ChangePassword)
		userApI.POST("/logout", ValidateAuth(userRepo), userHandler.LogOut)
	}

	router.Run(":8000")

}
