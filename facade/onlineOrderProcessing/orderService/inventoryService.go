package orderservice

import "fmt"

type InventoryService struct {
	stock map[string]int
}

func NewInventoryService() InventoryService {
	return InventoryService{
		stock: map[string]int{
			"SKU-001": 5,
			"SKU-002": 1,
		},
	}
}

func (this InventoryService) CheckStock(productId string, quantity int) bool {
	if qty, exists := this.stock[productId]; exists && qty >= quantity {
		return true
	}

	return false
}

func (this *InventoryService) ReserveStock(productId string, quantity int) {
	fmt.Printf("Inventory: Product %s in stock (%d available). Reserved %d units\n", productId, this.stock[productId], quantity)
	this.stock[productId] -= quantity

}

func (this *InventoryService) ReleaseStock(productId string, quantity int) {
	this.stock[productId] += quantity

	fmt.Printf("Rolling back: Releasing %d units of %s\n", quantity, productId)
}
