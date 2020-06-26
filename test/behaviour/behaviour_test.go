package behaviour

import (
	"github.com/aman-bansal/coffee_machine/pkg/use_case"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

var r = rand.New(rand.NewSource(99))

// this behaviour will never stop. its in infinite loop. Change configuration in base setup and this test will behave accordingly
func TestBehaviour(t *testing.T) {
	machine, inventory, vending := Setup()
	wg := sync.WaitGroup{}
	wg.Add(1)

	requestForCoffeeRepeatedlyAfterRandomeInterval(vending, machine)
	addIngredientsRepeatdlyAfterIntervals(inventory)

	wg.Wait()
}

func addIngredientsRepeatdlyAfterIntervals(inventory use_case.InventoryUseCase) {
	ingredientsRefillChoice := GetIngredientsRefillRequest()
	for key, val := range ingredientsRefillChoice {
		go func(ingName string, quant int) {
			for {
				// this is to specify gap when new ingredients gets added
				//will be added every time in between [0,10] seconds
				time.Sleep(time.Duration(time.Second.Nanoseconds() * 10))
				err := inventory.RefillIngredient(ingName, quant)
				if err != nil {
					log.Println("ingredient " + ingName + " addition failed because " + err.Error())
					continue
				}

				log.Println("ingredient " + ingName + " addition success ")
			}
		}(key, val)
	}
}

func requestForCoffeeRepeatedlyAfterRandomeInterval(vending use_case.VendingUseCase, machine use_case.MachineUseCase) {
	usersVsChoice := GetUsersAndTheirChoice()
	for key, val := range usersVsChoice {
		go func(user string, choice string) {
			for {
				//this is to specify gap between users request
				time.Sleep(time.Duration(time.Second.Nanoseconds() * r.Int63n(5)))
				outlet, err := vending.DispatchCoffee(choice)
				if err != nil {
					log.Println(user + " request for beverage " + choice + " cannot be served because " + err.Error())
					continue
				}

				log.Println(user + " request for beverage " + choice + " will be served from outlet id " + strconv.Itoa(outlet))
				//this is to specify wait users has to do to get the coffee
				time.Sleep(time.Duration(time.Second.Nanoseconds() * 2))
				_ = machine.MarkOutletAsFree(outlet)
				log.Println(user + " request for beverage " + choice + " is served")
			}
		}(key, val)
	}
}