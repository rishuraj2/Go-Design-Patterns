package machinestate

import (
	"fmt"
	vendingmachine "vendingmachine/internal/vendingMachine"
)

type HasMoneyState struct {}

func NewHasMoneyState() *HasMoneyState {
	return &HasMoneyState{}
}

func (this *HasMoneyState) SelectItem(context *vendingmachine.VendingMachine, itemCode string) {
	fmt.Println("Cannot change item after inserting money.")
}

func (this *HasMoneyState) InsertCoin(context *vendingmachine.VendingMachine, amount float64) {
	fmt.Println("Money already inserted.")
}

func (this *HasMoneyState) DispenseItem(context *vendingmachine.VendingMachine) {
	fmt.Println("Dispensing item: " + context.GetSelectedItem())
	context.SetState(NewDispensingState())
}
