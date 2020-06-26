package use_case

import (
	"errors"
)

type vending struct {
	inventory InventoryUseCase
	machine   MachineUseCase
}

func NewVendingUseCaseImpl(machine MachineUseCase, inventory InventoryUseCase) VendingUseCase {
	return &vending{
		inventory: inventory,
		machine:   machine,
	}
}

func (v vending) DispatchCoffee(drinkName string) (int, error) {
	drinkRecipe, err := v.inventory.GetRecipe(drinkName)
	if err != nil { return 0, errors.New("given recipe is not available") }

	outlet, err := v.machine.GetOutletIfAvailable()
	if err != nil { return 0, err }

	err = v.inventory.CheckAndUseIngredients(drinkRecipe.Ingredients)
	if err != nil {
		// to remain in consistent state. if ingredients not available. then mark selected outlet as free
		_ = v.machine.MarkOutletAsFree(outlet)
		return 0, err
	}

	return outlet, nil
}
