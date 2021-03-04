package controllers

import (
	"fmt"
	"net/http"

	"github.com/bookstore_items-api/domain/items"
	"github.com/bookstore_items-api/services"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (ic *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error to the caller
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: Return json to the user
	}

	fmt.Println(result)
	// TODO: return created item as json with http.status 201
}

func (ic *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
