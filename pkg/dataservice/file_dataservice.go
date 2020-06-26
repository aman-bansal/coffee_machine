package dataservice

import (
	"errors"
	"github.com/aman-bansal/coffee_machine/pkg/constant"
	"github.com/aman-bansal/coffee_machine/pkg/model"
	"sync"
)

//file data service has map structure to handle and save data
// there is mutex lock per collection which is being used to ensure concurrency
// with proper integration this lock situation can be easily handled via transaction from DB say MYSQL
type FileDataService struct {
	ingredients map[string]*model.Ingredient
	drinksRecipe map[string]*model.DrinkRecipe

	mutexLocks map[constant.CollectionName]sync.Mutex
}

func NewFileDataServiceImpl() FileBasedDataservice {
	locks := make(map[constant.CollectionName]sync.Mutex)
	locks[constant.INGREDIENT_COLLECTION] = sync.Mutex{}
	locks[constant.RECIPE_COLLECTION] = sync.Mutex{}

	return &FileDataService{
		ingredients:  make(map[string]*model.Ingredient),
		drinksRecipe: make(map[string]*model.DrinkRecipe),
		mutexLocks:   locks,
	}
}

//to update the ingredients
// add means ingredients are added. so given quantity is added to the previous one
// subtract means ingredients are being used, mean while dispatching the coffee
//This can be easily made generic via using update by query functionality of database say MYSQL
func (f *FileDataService) UpdateIngredients(opType constant.OpType, ingredients []*model.Ingredient) error {
	lock, _ := f.mutexLocks[constant.INGREDIENT_COLLECTION]
	lock.Lock()
	defer lock.Unlock()

	switch opType {
	case constant.ADD:
		for idx := range ingredients {
			// if ingredient present, then add the new quantity to the previous one
			if val, ok := f.ingredients[ingredients[idx].Name]; ok {
				val.Quantity = val.Quantity + ingredients[idx].Quantity
				continue
			}

			//if not present,return error. If new ingredient then create first
			return errors.New(ingredients[idx].Name + " ingredient specified is not available")
		}
	case constant.SUBTRACT:
		for idx := range ingredients {
			//this is for dispatch beverage case. If present and also quantity is greater than or equal to the required one
			// then only proceed
			if val, ok := f.ingredients[ingredients[idx].Name]; ok {
				if val.Quantity < ingredients[idx].Quantity {
					return errors.New(ingredients[idx].Name + " ingredient is low. please refill")
				}
				val.Quantity = val.Quantity - ingredients[idx].Quantity
				continue
			}

			//if not present return error
			return errors.New(ingredients[idx].Name + " ingredient specified is not available")
		}
	}

	return nil
}

//upsert function to replace the existing one
//this imitates db upsert operation. It will override the existing one or will create new one
func (f *FileDataService) Upsert(collectionName constant.CollectionName, id string, object interface{}) error {
	if collectionName != constant.INGREDIENT_COLLECTION && collectionName != constant.RECIPE_COLLECTION { return errors.New("collection is invalid")}

	switch collectionName {
	case constant.RECIPE_COLLECTION:
		lock, _ := f.mutexLocks[constant.RECIPE_COLLECTION]
		lock.Lock()
		defer lock.Unlock()

		f.drinksRecipe[id] = object.(*model.DrinkRecipe)
	case constant.INGREDIENT_COLLECTION:
		lock, _ := f.mutexLocks[constant.INGREDIENT_COLLECTION]
		lock.Lock()
		defer lock.Unlock()

		f.ingredients[id] = object.(*model.Ingredient)
	}

	return nil
}

//create function to insert a new resource
//this imitates db insert function
//if already present will throw error
func (f *FileDataService) Create(collectionName constant.CollectionName, id string, object interface{}) error {
	if collectionName != constant.INGREDIENT_COLLECTION && collectionName != constant.RECIPE_COLLECTION { return errors.New("collection is invalid")}

	switch collectionName {
	case constant.RECIPE_COLLECTION:
		lock, _ := f.mutexLocks[constant.RECIPE_COLLECTION]
		lock.Lock()
		defer lock.Unlock()

		if _, ok := f.drinksRecipe[id]; ok { return errors.New("given recipe already exist") }
		f.drinksRecipe[id] = object.(*model.DrinkRecipe)
	case constant.INGREDIENT_COLLECTION:
		lock, _ := f.mutexLocks[constant.INGREDIENT_COLLECTION]
		lock.Lock()
		defer lock.Unlock()

		if _, ok := f.ingredients[id]; ok { return errors.New("given ingredients already exist") }
		f.ingredients[id] = object.(*model.Ingredient)
	}

	return nil
}

//get function to retrieve the current resource
// imitates db get function
//if not present, will return error
func (f *FileDataService) Get(collectionName constant.CollectionName, id string) (interface{}, error) {
	if collectionName != constant.INGREDIENT_COLLECTION && collectionName != constant.RECIPE_COLLECTION { return nil, errors.New("collection is invalid")}

	switch collectionName {
	case constant.RECIPE_COLLECTION:
		lock, _ := f.mutexLocks[constant.RECIPE_COLLECTION]
		lock.Lock()
		defer lock.Unlock()

		if _, ok := f.drinksRecipe[id]; !ok { return nil, errors.New("recipe is not found") }
		return f.drinksRecipe[id], nil
	case constant.INGREDIENT_COLLECTION:
		lock, _ := f.mutexLocks[constant.INGREDIENT_COLLECTION]
		lock.Lock()
		defer lock.Unlock()

		if _, ok := f.ingredients[id]; !ok { return nil, errors.New("ingredients is not found") }
		return f.ingredients[id], nil
	}

	return nil, nil
}