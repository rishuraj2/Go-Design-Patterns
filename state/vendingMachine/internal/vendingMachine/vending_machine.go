package vendingmachine

type MachineState interface {
	SelectItem(context *VendingMachine, itemCode string)
	InsertCoin(context *VendingMachine, amount float64)
	DispenseItem(context *VendingMachine)
}

type VendingMachine struct {
	currentState   MachineState
	selectedItem   string
	insertedAmount float64
}

func NewVendingMachine(state MachineState) *VendingMachine {
	return &VendingMachine{
		currentState: state,
	}
}

func (this *VendingMachine) SetState(newState MachineState) {
	this.currentState = newState
}

func (this *VendingMachine) SetSelectedItem(itemCode string) {
	this.selectedItem = itemCode
}

func (this *VendingMachine) SetInsertedAmount(amount float64) {
	this.insertedAmount = amount
}

func (this *VendingMachine) GetSelectedItem() string {
	return this.selectedItem
}

func (this *VendingMachine) SelectItem(itemCode string) {
	this.currentState.SelectItem(this, itemCode)
}

func (this *VendingMachine) InsertCoin(amount float64) {
	this.currentState.InsertCoin(this, amount)
}

func (this *VendingMachine) DispenseItem() {
	this.currentState.DispenseItem(this)
}

func (this *VendingMachine) Reset(state MachineState) {
	this.selectedItem = ""
	this.insertedAmount = 0.0
	this.currentState = state
}
