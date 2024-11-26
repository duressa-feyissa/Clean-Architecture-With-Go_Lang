package main

import (
	"cleantaskmanager/db"
	"cleantaskmanager/delivery/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := db.Client.Database(os.Getenv("DB_NAME"))

	r := gin.Default()
	routers.Setup(db, r)
	r.Run(":8080")

}
