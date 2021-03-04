package app

import (
	"net/http"

	"github.com/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
