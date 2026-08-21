package machinestate

import (
	"fmt"
	"strconv"
	vendingmachine "vendingmachine/internal/vendingMachine"
)

type ItemSelectedState struct {}

func NewItemSelectedState() *ItemSelectedState {
	return &ItemSelectedState{}
}

func (this *ItemSelectedState) SelectItem(context *vendingmachine.VendingMachine, itemCode string) {
	fmt.Println("Item already selected: " + itemCode)
}

func (this *ItemSelectedState) InsertCoin(context *vendingmachine.VendingMachine, amount float64) {
	fmt.Println("Inserted $" + strconv.FormatFloat(amount, 'f', -1, 64) + " for item: " + context.GetSelectedItem())
	context.SetInsertedAmount(amount)
	context.SetState(NewHasMoneyState())
}

func (this *ItemSelectedState) DispenseItem(context *vendingmachine.VendingMachine) {
	fmt.Println("Insert coin before dispensing.")
}
