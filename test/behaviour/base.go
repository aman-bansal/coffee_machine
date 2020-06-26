package behaviour

import (
	"encoding/json"
	"github.com/aman-bansal/coffee_machine/pkg/model"
	"github.com/aman-bansal/coffee_machine/pkg/use_case"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

type UseCaseData struct {
	Machine *Machine `json:"machine"`
}

type Machine struct {
	Outlets *Outlets `json:"outlets"`
	TotalItemQuantity map[string]int `json:"total_items_quantity"`
	Beverages map[string]map[string]int `json:"beverages"`
}

type Outlets struct {
	Count int `json:"count_n"`
}

func GetUsersAndTheirChoice() map[string]string {
	var testCase map[string]string = make(map[string]string)
	testCase["user 1"] = "hot_tea"
	testCase["user 2"] = "hot_coffee"
	testCase["user 3"] = "black_tea"
	testCase["user 4"] = "green_tea"
	testCase["user 5"] = "hot_tea"
	testCase["user 6"] = "black_tea"
	testCase["user 7"] = "green_tea"
	testCase["user 8"] = "hot_coffee"
	testCase["user 9"] = "black_tea"
	testCase["user 10"] = "green_tea"

	return testCase
}

func GetIngredientsRefillRequest() map[string]int {
	var testCase map[string]int = make(map[string]int)
	testCase["hot_water"] = 500
	testCase["hot_milk"] = 500
	testCase["ginger_syrup"] = 500
	testCase["sugar_syrup"] = 500
	testCase["tea_leaves_syrup"] = 500

	return testCase
}

func Setup() (use_case.MachineUseCase, use_case.InventoryUseCase, use_case.VendingUseCase) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	bytes, err := ioutil.ReadFile(path.Join(path.Dir(currentFilePath), "test_cases/test_case1.json"))
	if err != nil { log.Fatal("not able to open test cases file", err) }

	testData := new(UseCaseData)
	err = json.Unmarshal(bytes, testData)
	if err != nil { log.Fatal("not able to read test cases data into struct objects") }

	if testData.Machine == nil || testData.Machine.Outlets == nil || testData.Machine.Outlets.Count == 0 {
		log.Fatal("outlets are not found. please check the test cases")
	}

	machine := use_case.NewMachineUseCaseImpl(testData.Machine.Outlets.Count)
	inventory := use_case.NewInventoryUseCaseImpl(transformToIngredients(testData.Machine.TotalItemQuantity), transformToDrinkRecipe(testData.Machine.Beverages))

	vending := use_case.NewVendingUseCaseImpl(machine, inventory)

	return machine, inventory, vending
}

func transformToIngredients(items map[string]int) []*model.Ingredient {
	if len(items) == 0 { }
	result := make([]*model.Ingredient, 0)
	for key, val := range items {
		result = append(result, &model.Ingredient{
			Name:     key,
			Quantity: val,
		})
	}
	return result
}

func transformToDrinkRecipe(recipes map[string]map[string]int) []*model.DrinkRecipe {
	result := make([]*model.DrinkRecipe, 0)
	for recipeName, val := range recipes {
		ingredients := make([]*model.Ingredient, 0)
		for ingName, quant := range val {
			ingredients = append(ingredients, &model.Ingredient{
				Name:     ingName,
				Quantity: quant,
			})
		}

		result = append(result, &model.DrinkRecipe{
			RecipeName:  recipeName,
			Ingredients: ingredients,
		})
	}
	return result
}