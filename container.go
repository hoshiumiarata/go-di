package di

import "reflect"

// Container is a dependency injection container.
type Container struct {
	parent             *Container
	typesContainer     map[reflect.Type]*reflect.Value
	positionsContainer map[int]*reflect.Value
}

// New creates a new dependency injection container.
func New() *Container {
	container := &Container{
		typesContainer:     make(map[reflect.Type]*reflect.Value),
		positionsContainer: make(map[int]*reflect.Value),
	}
	RegisterValueTo(container, container)
	return container
}

// New creates a new child dependency injection container.
func (container *Container) New() *Container {
	child := New()
	child.parent = container
	return child
}

// Parent returns the parent dependency injection container.
func (container *Container) Parent() *Container {
	return container.parent
}
