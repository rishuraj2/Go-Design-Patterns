package data

import (
	"fmt"
	"sync"
)

type ImageDataRegistry struct {
	registry map[string][]byte
}

var (
	instance *ImageDataRegistry
	once     sync.Once
)

func GetImageDataRegistry() *ImageDataRegistry {
	once.Do(func() {
		instance = &ImageDataRegistry{
			registry: make(map[string][]byte),
		}
	})

	return instance
}

func (this *ImageDataRegistry) Register(filename string, data []byte) error {
	if _, exists := this.registry[filename]; exists {
		return fmt.Errorf(`filename "%s" already exists`, filename)
	}

	this.registry[filename] = data
	return nil
}

func (this *ImageDataRegistry) Fetch(filename string) ([]byte, error) {
	if data, exists := this.registry[filename]; exists {
		return data, nil
	}

	return []byte{}, nil
}
