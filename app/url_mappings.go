package app

import (
	"github.com/angadthandi/bookstore_users-api/controllers/ping"
	"github.com/angadthandi/bookstore_users-api/controllers/users"
)

func mapUrls(
// dbClient *sql.DB,
) {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	// router.GET("/users/search", users.Search)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
