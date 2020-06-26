package use_case

import (
	"fmt"
	"github.com/aman-bansal/coffee_machine/pkg/use_case"
	"testing"
)

const numOfOutlet = 5
func setupMachineTestCase() use_case.MachineUseCase{
	return use_case.NewMachineUseCaseImpl(numOfOutlet)
}

func TestGetOutletIfAvailable_Error(t *testing.T) {
	machineUseCase := setupMachineTestCase()
	//build
	//occupy all the outlets
	idx := 0
	for {
		if idx == numOfOutlet { break }
		_, _ = machineUseCase.GetOutletIfAvailable()
		idx = idx + 1
	}
	//assert
	outletId, err := machineUseCase.GetOutletIfAvailable()
	if outletId == 0 && err != nil {
		fmt.Println(err.Error())
	} else {
		t.Fatalf("TestGetOutletIfAvailable_Error failed.")
	}
	//destroy
	_ = machineUseCase.Clear()
}

func TestGetOutletIfAvailable_Success(t *testing.T) {
	machineUseCase := setupMachineTestCase()
	//build
	//occupy all the outlets
	idx := 0
	for {
		if idx == numOfOutlet - 1 { break }
		_, _ = machineUseCase.GetOutletIfAvailable()
		idx = idx + 1
	}

	//assert
	outletId, err := machineUseCase.GetOutletIfAvailable()
	if outletId == 0 || err != nil {
		t.Fatalf("TestGetOutletIfAvailable_Success failed.")
	}

	fmt.Println("found the outlet id ", outletId)
	//destroy
	//clear all the outlets
	_ = machineUseCase.Clear()
}

func TestMarkOutletAsFree_Error(t *testing.T) {
	machineUseCase := setupMachineTestCase()
	var outletId int = 1
	//build not required
	idx := 0
	for {
		if idx == numOfOutlet - 1 { break }
		outletId, _ = machineUseCase.GetOutletIfAvailable()
		idx = idx + 1
	}
	//assert
	err := machineUseCase.MarkOutletAsFree(outletId)
	if err == nil {
		t.Fatalf("TestMarkOutletAsFree_Error failed.")
	}

	fmt.Println("free the outlet id with error ", err.Error())
	//destroy
	//clear all the outlets
	_ = machineUseCase.Clear()
}

func TestMarkOutletAsFree_Success(t *testing.T) {
	machineUseCase := setupMachineTestCase()
	var outletId int
	//build
	//get one outlet
	idx := 0
	for {
		if idx == numOfOutlet - 1 { break }
		outletId, _ = machineUseCase.GetOutletIfAvailable()
		idx = idx + 1
	}

	//assert : try to mark it free
	err := machineUseCase.MarkOutletAsFree(outletId)
	if err != nil {
		t.Fatalf("TestMarkOutletAsFree_Success failed.")
	}

	fmt.Println("free the outlet id ", outletId)
	//destroy
	//clear all the outlets
	_ = machineUseCase.Clear()
}