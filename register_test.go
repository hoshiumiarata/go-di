package di_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-di"
)

func TestRegisterValueTo(t *testing.T) {
	container := di.New()
	i := 1
	di.RegisterValueTo(container, i)
	value, ok := container.ResolveType(reflect.TypeOf(i))
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered")
	}
}

func TestUnregisterValueFrom(t *testing.T) {
	container := di.New()
	i := 1
	di.RegisterValueTo(container, i)
	di.UnregisterValueFrom(container, i)
	_, ok := container.ResolveType(reflect.TypeOf(i))
	if ok {
		t.Error("expected i to not be registered")
	}
}

func TestRegisterByTypeTo(t *testing.T) {
	t.Run("Assignable", func(t *testing.T) {
		container := di.New()
		i := 1
		di.RegisterByTypeTo(container, reflect.TypeOf(i), i)
		value, ok := container.ResolveType(reflect.TypeOf(i))
		if !ok || value.Int() != 1 {
			t.Error("expected i to be registered")
		}
	})
	t.Run("Not Assignable", func(t *testing.T) {
		container := di.New()
		i := 1
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		di.RegisterByTypeTo(container, reflect.TypeOf(""), i)
	})
}

func TestUnregisterByTypeFrom(t *testing.T) {
	container := di.New()
	i := 1
	di.RegisterByTypeTo(container, reflect.TypeOf(i), i)
	di.UnregisterByTypeFrom(container, reflect.TypeOf(i))
	_, ok := container.ResolveType(reflect.TypeOf(i))
	if ok {
		t.Error("expected i to not be registered")
	}
}

func TestRegisterByPositionTo(t *testing.T) {
	container := di.New()
	i := 1
	di.RegisterByPositionTo(container, 0, i)
	value, ok := container.ResolvePosition(0)
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered")
	}
}

func TestUnregisterByPositionFrom(t *testing.T) {
	container := di.New()
	i := 1
	di.RegisterByPositionTo(container, 0, i)
	di.UnregisterByPositionFrom(container, 0)
	_, ok := container.ResolvePosition(0)
	if ok {
		t.Error("expected i to not be registered")
	}
}
