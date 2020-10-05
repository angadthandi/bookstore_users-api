package app

import (
	"log"

	"github.com/angadthandi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	err := users_db.Client.Ping()
	if err != nil {
		log.Fatalf("unable to connect to mysql db error: %v", err)
	}

	log.Println("StartApplication...")

	mapUrls()
	router.Run(":8080")
}
