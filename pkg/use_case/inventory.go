package use_case

import (
	"github.com/aman-bansal/coffee_machine/pkg/constant"
	"github.com/aman-bansal/coffee_machine/pkg/dataservice"
	"github.com/aman-bansal/coffee_machine/pkg/model"
	"log"
)

//only dependency is fileDataService
//this will save all the necessary ingredient/recipe information in file based data system
type inventory struct {
	fileDataService dataservice.FileBasedDataservice
}

//to create the initial inventory
func NewInventoryUseCaseImpl(ingredients []*model.Ingredient, recipes []*model.DrinkRecipe) InventoryUseCase {
	daoService := dataservice.NewFileDataServiceImpl()
	for idx := range ingredients {
		err := daoService.Create(constant.INGREDIENT_COLLECTION, ingredients[idx].Name, ingredients[idx])
		if err != nil {
			log.Fatal("error while initializing create ingredients use case")
		}
	}

	for idx := range recipes {
		err := daoService.Create(constant.RECIPE_COLLECTION, recipes[idx].RecipeName, recipes[idx])
		if err != nil {
			log.Fatal("error while initializing create recipes use cases")
		}
	}

	return &inventory{
		fileDataService: daoService,
	}
}

//add new ingredient
// if already present, will throw error. Use Refill Instead
func (i inventory) AddNewIngredient(id string, ingredient *model.Ingredient) error {
	err := i.fileDataService.Create(constant.INGREDIENT_COLLECTION, id, ingredient)
	if err != nil { return err }

	return nil
}

//to refill ingredient
// if ingredient not preset, this will throw an error
// if present, it will add the provided quantity to the existing remaining quantity
func (i inventory) RefillIngredient(ingredientName string, quantity int) error {
	err := i.fileDataService.UpdateIngredients(constant.ADD, []*model.Ingredient{
		{
			Name:     ingredientName,
			Quantity: quantity,
		},
	})

	if err != nil { return err }
	return nil
}

//to use ingredient
// if ingredient not preset, this will throw an error
// if present, it will subtract the provided quantity to the existing remaining quantity if possible otherwise will throw an error
func (i inventory) CheckAndUseIngredients(ingredients []*model.Ingredient) error {
	err := i.fileDataService.UpdateIngredients(constant.SUBTRACT, ingredients)
	if err != nil { return err }

	return nil
}

//to get the recipe
//if not present will throw an error
func (i inventory) GetRecipe(id string) (*model.DrinkRecipe, error) {
	recipe, err := i.fileDataService.Get(constant.RECIPE_COLLECTION, id)
	if err != nil { return nil, err }

	return recipe.(*model.DrinkRecipe), nil
}

//to add the new recipe
// if already present, this will throw an error
func (i inventory) AddRecipe(id string, recipe *model.DrinkRecipe) error {
	err := i.fileDataService.Create(constant.RECIPE_COLLECTION, id, recipe)
	if err != nil { return err }

	return nil
}

//to update the recipe
//if not present, this will throw an error
func (i inventory) UpdateRecipe(id string, recipe *model.DrinkRecipe) error {
	err := i.fileDataService.Upsert(constant.RECIPE_COLLECTION, id, recipe)
	if err != nil { return err }

	return nil
}
