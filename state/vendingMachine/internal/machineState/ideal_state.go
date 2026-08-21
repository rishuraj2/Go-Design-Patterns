package machinestate

import (
	"fmt"
	vendingmachine "vendingmachine/internal/vendingMachine"
)

type IdealState struct{}

func NewIdealState() *IdealState {
	return &IdealState{}
}

func (this *IdealState) SelectItem(context *vendingmachine.VendingMachine, itemCode string) {
	fmt.Println("Item selected: " + itemCode)
	context.SetSelectedItem(itemCode)
	context.SetState(NewItemSelectedState())
}

func (this *IdealState) InsertCoin(context *vendingmachine.VendingMachine, amount float64) {
	fmt.Println("Please select an item before inserting coins.")
}

func (this *IdealState) DispenseItem(context *vendingmachine.VendingMachine) {
	fmt.Println("No item selected. Nothing to dispense.")
}
