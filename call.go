package di

import (
	"errors"
	"reflect"
)

// Call calls a function with dependency injection.
func (container *Container) Call(f any) []reflect.Value {
	funcType := reflect.TypeOf(f)
	if funcType.Kind() != reflect.Func {
		panic(errors.New("f is not a function"))
	}
	numIn := funcType.NumIn()
	args := make([]reflect.Value, numIn)
	for i := 0; i < numIn; i++ {
		inType := funcType.In(i)
		inValue, ok := container.ResolvePosition(i)
		if ok && inValue.Type().AssignableTo(inType) {
			args[i] = inValue
			continue
		}
		inValue, ok = container.ResolvePosition(i - numIn)
		if ok && inValue.Type().AssignableTo(inType) {
			args[i] = inValue
			continue
		}
		inValue, ok = container.ResolveType(inType)
		if !ok {
			inValue = reflect.Zero(inType)
		}
		args[i] = inValue
	}
	return reflect.ValueOf(f).Call(args)
}
