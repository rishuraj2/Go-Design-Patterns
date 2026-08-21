package main

import (
	machinestate "vendingmachine/internal/machineState"
	vendingmachine "vendingmachine/internal/vendingMachine"
)

func main() {
	machine := vendingmachine.NewVendingMachine(machinestate.NewIdealState())
	machine.SelectItem("2A")
	machine.InsertCoin(5.2)
	machine.InsertCoin(6.3)
	machine.DispenseItem()
}
