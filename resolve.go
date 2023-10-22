package di

import "reflect"

// ResolveType resolves a value from the container by type.
func (container *Container) ResolveType(t reflect.Type) (value reflect.Value, ok bool) {
	valuePtr, ok := container.typesContainer[t]
	if valuePtr != nil {
		return *valuePtr, true
	}
	if ok {
		return value, false
	}
	if container.parent != nil {
		return container.parent.ResolveType(t)
	}
	return
}

// ResolvePosition resolves a value from the container by function argument position.
func (container *Container) ResolvePosition(position int) (value reflect.Value, ok bool) {
	valuePtr, ok := container.positionsContainer[position]
	if valuePtr != nil {
		return *valuePtr, true
	}
	if ok {
		return value, false
	}
	if container.parent != nil {
		return container.parent.ResolvePosition(position)
	}
	return
}
