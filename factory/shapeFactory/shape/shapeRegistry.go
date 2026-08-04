package shape

import (
	"errors"
	"sync"
)

type shapeRegistry struct {
	registry map[string]Shape
	mu       sync.RWMutex
}

var (
	errDuplicateShape = errors.New("shape already registered")
	errShapeNotFound  = errors.New("shape not found")
)

var (
	instance *shapeRegistry
	once     sync.Once
)

func getShapeRegistry() *shapeRegistry {
	once.Do(func() {
		instance = &shapeRegistry{
			registry: make(map[string]Shape),
		}
	})

	return instance
}

func (r *shapeRegistry) register(name string, shape Shape) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.registry[name]; exists {
		return errDuplicateShape
	}

	r.registry[name] = shape
	return nil
}

func (r *shapeRegistry) getShape(name string) (Shape, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if shape, exists := r.registry[name]; exists {
		return shape, nil
	}

	return nil, errShapeNotFound
}
