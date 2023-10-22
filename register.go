package di

import (
	"errors"
	"reflect"
)

// RegisterValueTo registers a value to the container.
func RegisterValueTo[T any](container *Container, value T) {
	reflectedValue := reflect.ValueOf(value)
	container.typesContainer[reflect.TypeOf(&value).Elem()] = &reflectedValue
}

// UnregisterValueFrom unregisters a value from the container.
func UnregisterValueFrom[T any](container *Container, value T) {
	container.typesContainer[reflect.TypeOf(&value).Elem()] = nil
}

// RegisterByTypeTo registers a value to the container by type.
func RegisterByTypeTo(container *Container, t reflect.Type, value any) {
	valueType := reflect.TypeOf(value)
	if !valueType.AssignableTo(t) {
		panic(errors.New("value is not assignable to t"))
	}
	reflectedValue := reflect.ValueOf(value)
	container.typesContainer[t] = &reflectedValue
}

// UnregisterByTypeFrom unregisters a value from the container by type.
func UnregisterByTypeFrom(container *Container, t reflect.Type) {
	container.typesContainer[t] = nil
}

// RegisterByPositionTo registers a value to the container by function argument position.
func RegisterByPositionTo(container *Container, position int, value any) {
	reflectedValue := reflect.ValueOf(value)
	container.positionsContainer[position] = &reflectedValue
}

// UnregisterByPositionFrom unregisters a value from the container by function argument position.
func UnregisterByPositionFrom(container *Container, position int) {
	container.positionsContainer[position] = nil
}
