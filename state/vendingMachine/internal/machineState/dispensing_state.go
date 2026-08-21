package machinestate

import (
	"fmt"
	vendingmachine "vendingmachine/internal/vendingMachine"
)

type DispensingState struct {}

func NewDispensingState() *DispensingState {
	return &DispensingState{}
}

func (this *DispensingState) SelectItem(context *vendingmachine.VendingMachine, itemCode string) {
	fmt.Println("Please wait, dispensing in progress.")
}

func (this *DispensingState) InsertCoin(context *vendingmachine.VendingMachine, amount float64) {
	fmt.Println("Please wait, dispensing in progress.")
}

func (this *DispensingState) DispenseItem(context *vendingmachine.VendingMachine) {
	fmt.Println("Item dispensed successfully.")
	context.Reset(NewIdealState())
}
