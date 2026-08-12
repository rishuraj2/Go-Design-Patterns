package paymentmethod

import (
	"fmt"
	"sync"
)

type PaymentMethodStore struct {
	store map[string]PaymentMethod
}

var (
	instance *PaymentMethodStore
	once     sync.Once
)

func GetPaymentMethodStore() *PaymentMethodStore {
	once.Do(func() {
		instance = &PaymentMethodStore{
			store: make(map[string]PaymentMethod),
		}
	})

	return instance
}

func (this *PaymentMethodStore) Register(name string, method PaymentMethod) error {
	if _, exists := this.store[name]; exists {
		return fmt.Errorf("method already exists")
	}

	this.store[name] = method
	return nil
}

func (this *PaymentMethodStore) Fetch(name string) (PaymentMethod, error) {
	if method, exists := this.store[name]; exists {
		return method, nil
	}

	return nil, fmt.Errorf("method does not exists")
}
