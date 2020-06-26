package use_case

import "github.com/aman-bansal/coffee_machine/pkg/model"

//this is for vending use case, i.e. dispatching coffee
//in future can be extended to override or cancel
type VendingUseCase interface {
	DispatchCoffee(drinkName string) (int, error)
}

//This relates to the use cases related to coffee machine inventory
//use cases related to add/refill/use ingredients
//use cases related to get/add/update recipe
type InventoryUseCase interface {
	AddNewIngredient(id string, recipe *model.Ingredient) error
	RefillIngredient(ingredientName string, quantity int) error
	CheckAndUseIngredients(ingredients []*model.Ingredient) error
	GetRecipe(id string) (*model.DrinkRecipe, error)
	AddRecipe(id string, recipe *model.DrinkRecipe) error
	UpdateRecipe(id string, recipe *model.DrinkRecipe) error
}

//this relates to the use cases to maintain and check machine state
//use cases like which outlets are available will be maintained by this
// in case of any error or inconsistent state, clear() function can be used
type MachineUseCase interface {
	GetOutletIfAvailable() (int, error)
	MarkOutletAsFree(outletId int) error
	Clear() error
}