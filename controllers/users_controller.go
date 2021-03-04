package controllers

import "net/http"

var (
	UsersController usersControllerInterface = &itemsController{}
)

type usersControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type usersController struct{}

func (uc *usersController) Create() {

}

func (uc *usersController) Get() {

}
