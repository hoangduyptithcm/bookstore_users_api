package app

import "https/github.com/hoangduyptithcm/bookstore_users_api/controllers"

func mapUrls() {

	router.GET("/ping", controllers.Ping)
}
