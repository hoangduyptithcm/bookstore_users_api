package app

import "github.com/hoangduyptithcm/bookstore_users_api/controllers"

func mapUrls() {

	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_ud", controllers.FindUserByName)
	router.POST("/users", controllers.CreateUser)
}
