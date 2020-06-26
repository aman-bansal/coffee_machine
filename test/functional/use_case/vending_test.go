package use_case

import (
	"fmt"
	"github.com/aman-bansal/coffee_machine/pkg/model"
	"github.com/aman-bansal/coffee_machine/pkg/use_case"
	"testing"
)

type vendingTestCase struct {
	name        string
	userChoice  string
	ingredients []*model.Ingredient
	recipe      []*model.DrinkRecipe
}

var testCases []vendingTestCase = []vendingTestCase{
	{
		name:        "Success",
		userChoice: "hot_tea",
		ingredients: []*model.Ingredient{
			{
				Name:     "hot_water",
				Quantity: 100,
			},
		},
		recipe:      []*model.DrinkRecipe{
			{
				RecipeName:  "hot_tea",
				Ingredients: []*model.Ingredient{
					{
						Name:     "hot_water",
						Quantity: 100,
					},
				},
			},
		},
	},
	{
		name:        "recipe not found",
		userChoice: "hot_coffee",
		ingredients: []*model.Ingredient{
			{
				Name:     "hot_water",
				Quantity: 100,
			},
		},
		recipe:      []*model.DrinkRecipe{
			{
				RecipeName:  "hot_tea",
				Ingredients: []*model.Ingredient{
					{
						Name:     "hot_water",
						Quantity: 100,
					},
				},
			},
		},
	},
	{
		name:        "outlet not found",
		userChoice: "hot_tea",
		ingredients: []*model.Ingredient{
			{
				Name:     "hot_water",
				Quantity: 100,
			},
		},
		recipe:      []*model.DrinkRecipe{
			{
				RecipeName:  "hot_tea",
				Ingredients: []*model.Ingredient{
					{
						Name:     "hot_water",
						Quantity: 100,
					},
				},
			},
		},
	},
	{
		name:        "ingredient not available",
		userChoice: "hot_tea",
		ingredients: []*model.Ingredient{
			{
				Name:     "hot_water",
				Quantity: 100,
			},
		},
		recipe:      []*model.DrinkRecipe{
			{
				RecipeName:  "hot_tea",
				Ingredients: []*model.Ingredient{
					{
						Name:     "hot_water",
						Quantity: 100,
					},
					{
						Name:     "hot_milk",
						Quantity: 200,
					},
				},
			},
		},
	},
	{
		name:       "ingredient not enough",
		userChoice: "hot_tea",
		ingredients: []*model.Ingredient{
			{
				Name:     "hot_water",
				Quantity: 100,
			},
		},
		recipe:      []*model.DrinkRecipe{
			{
				RecipeName:  "hot_tea",
				Ingredients: []*model.Ingredient{
					{
						Name:     "hot_water",
						Quantity: 200,
					},
				},
			},
		},
	},
}

func setupVendingTestEnv(testCase vendingTestCase) (use_case.MachineUseCase, use_case.VendingUseCase) {
	machine := use_case.NewMachineUseCaseImpl(1)
	inventory := use_case.NewInventoryUseCaseImpl(testCase.ingredients, testCase.recipe)
	return machine, use_case.NewVendingUseCaseImpl(machine, inventory)
}

func TestDispatchCoffee_Success(t *testing.T) {
	machineUseCase, vendingUseCase := setupVendingTestEnv(testCases[0])
	//build no need to build initial state

	outletId, err := vendingUseCase.DispatchCoffee(testCases[0].userChoice)
	if err != nil {
		t.Fatalf("TestDispatchCoffee_Success failed: got err %v", err.Error())
	}

	if outletId != 1 {
		t.Fatalf("TestDispatchCoffee_Success failed: outlet id is 0 ")
	}

	fmt.Println("TestDispatchCoffee_Success success")
	//destroy
	_ = machineUseCase.MarkOutletAsFree(outletId)
}

func TestDispatchCoffee_RecipeNotFound(t *testing.T) {
	_, vendingUseCase := setupVendingTestEnv(testCases[1])
	//build no need to build initial state

	outletId, err := vendingUseCase.DispatchCoffee(testCases[1].userChoice)
	if err == nil {
		t.Fatalf("TestDispatchCoffee_RecipeNotFound failed: no error")
	}

	if outletId != 0 {
		t.Fatalf("TestDispatchCoffee_RecipeNotFound failed: got valid outlet id ")
	}

	fmt.Println("TestDispatchCoffee_RecipeNotFound success: get the right err ", err.Error())
}

func TestDispatchCoffee_OutletNotFound(t *testing.T) {
	machineUseCase, vendingUseCase := setupVendingTestEnv(testCases[2])
	//build
	_, _ = machineUseCase.GetOutletIfAvailable()

	//assert
	outletId, err := vendingUseCase.DispatchCoffee(testCases[2].userChoice)
	if err == nil {
		t.Fatalf("TestDispatchCoffee_OutletNotFound failed: no error")
	}

	if outletId != 0 {
		t.Fatalf("TestDispatchCoffee_OutletNotFound failed: got valid outlet id ")
	}

	fmt.Println("TestDispatchCoffee_OutletNotFound success: get the right err ", err.Error())

	//destroy
	_ = machineUseCase.Clear()
}

func TestDispatchCoffee_IngredientsNotAvailable(t *testing.T) {
	_, vendingUseCase := setupVendingTestEnv(testCases[3])
	//build no need to build initial state

	outletId, err := vendingUseCase.DispatchCoffee(testCases[3].userChoice)
	if err == nil {
		t.Fatalf("TestDispatchCoffee_IngredientsNotAvailable failed: no error")
	}

	if outletId != 0 {
		t.Fatalf("TestDispatchCoffee_IngredientsNotAvailable failed: got valid outlet id ")
	}

	fmt.Println("TestDispatchCoffee_IngredientsNotAvailable success: get the right err ", err.Error())
}

func TestDispatchCoffee_IngredientsNotEnough(t *testing.T) {
	_, vendingUseCase := setupVendingTestEnv(testCases[4])
	//build no need to build initial state

	outletId, err := vendingUseCase.DispatchCoffee(testCases[4].userChoice)
	if err == nil {
		t.Fatalf("TestDispatchCoffee_IngredientsNotEnough failed: no error")
	}

	if outletId != 0 {
		t.Fatalf("TestDispatchCoffee_IngredientsNotEnough failed: got valid outlet id ")
	}

	fmt.Println("TestDispatchCoffee_IngredientsNotEnough success: get the right err ", err.Error())
}