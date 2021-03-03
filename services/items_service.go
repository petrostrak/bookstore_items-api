package services

import (
	"net/http"

	"github.com/bookstore_items-api/domain/items"
	"github.com/bookstore_items-api/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
}

type itemsService struct{}

func (is *itemsService) Create(item items.Item) (*items.Item, errors.RestErr) {
	return nil, errors.NewRestError("TODO: imple", http.StatusNotImplemented, "not implemented", nil)
}

func (is *itemsService) Get(string) (*items.Item, errors.RestErr) {
	return nil, errors.NewRestError("TODO: imple", http.StatusNotImplemented, "not implemented", nil)
}
