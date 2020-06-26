package use_case

import (
	"github.com/aman-bansal/coffee_machine/pkg/model"
)

type machine struct {
	outletPool *model.MachineOutletPool
}

func NewMachineUseCaseImpl(numOfOutlets int) MachineUseCase {
	return &machine{
		outletPool: model.NewMachineOutletPool(numOfOutlets),
	}
}

//get one outlet which is available. This selects one of the available outlets and return it
// the returned outlet can be used to dispatch the coffee
func (m machine) GetOutletIfAvailable() (int, error) {
	outlet, err := m.outletPool.GetAvailableOutlet()
	if err != nil { return 0, err }

	return outlet, nil
}

// when outlet has fully dispatched the coffee, mark it as free so that it can serve another customer
func (m machine) MarkOutletAsFree(outletId int) error {
	err := m.outletPool.MarkOutletFree(outletId)
	if err != nil { return err }

	return nil
}

// if in consistent state, mark all outlets as available
func (m machine) Clear() error {
	m.outletPool.MarkAllOutletsAvailable()
	return nil
}