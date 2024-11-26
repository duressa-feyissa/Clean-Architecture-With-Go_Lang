package routers

import (
	"cleantaskmanager/infrastructure"
	"cleantaskmanager/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewUserRouter(db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(infrastructure.AuthMiddleware())
	// All Private APIs
	NewTaskRouter(db, protectedRouter)
}
