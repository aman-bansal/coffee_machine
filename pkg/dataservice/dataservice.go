package dataservice

import (
	"github.com/aman-bansal/coffee_machine/pkg/constant"
	"github.com/aman-bansal/coffee_machine/pkg/model"
)

// this service actually imitates databases and lock are introduced to manage concurrency.
type FileBasedDataservice interface {
	UpdateIngredients(opType constant.OpType, ingredients []*model.Ingredient) error
	Upsert(collectionName constant.CollectionName, id string, object interface{}) error
	Create(collectionName constant.CollectionName, id string, object interface{}) error
	Get(collectionName constant.CollectionName, id string) (interface{}, error)
}
