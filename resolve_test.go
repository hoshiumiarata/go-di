package di_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-di"
)

func TestResolveType(t *testing.T) {
	parent := di.New()
	i := 1
	di.RegisterValueTo(parent, i)
	value, ok := parent.ResolveType(reflect.TypeOf(i))
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in parent")
	}
	child := parent.New()
	value, ok = child.ResolveType(reflect.TypeOf(i))
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in child")
	}
	di.UnregisterValueFrom(child, i)
	_, ok = child.ResolveType(reflect.TypeOf(i))
	if ok {
		t.Error("expected i to not be registered in child")
	}
	value, ok = parent.ResolveType(reflect.TypeOf(i))
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in parent")
	}
	i = 2
	di.RegisterValueTo(child, i)
	value, ok = child.ResolveType(reflect.TypeOf(i))
	if !ok || value.Int() != 2 {
		t.Error("expected i to be registered in child")
	}
	value, ok = parent.ResolveType(reflect.TypeOf(i))
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in parent")
	}
}

func TestResolvePosition(t *testing.T) {
	parent := di.New()
	i := 1
	di.RegisterByPositionTo(parent, 0, i)
	value, ok := parent.ResolvePosition(0)
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in parent")
	}
	child := parent.New()
	value, ok = child.ResolvePosition(0)
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in child")
	}
	di.UnregisterByPositionFrom(child, 0)
	_, ok = child.ResolvePosition(0)
	if ok {
		t.Error("expected i to not be registered in child")
	}
	value, ok = parent.ResolvePosition(0)
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in parent")
	}
	i = 2
	di.RegisterByPositionTo(child, 0, i)
	value, ok = child.ResolvePosition(0)
	if !ok || value.Int() != 2 {
		t.Error("expected i to be registered in child")
	}
	value, ok = parent.ResolvePosition(0)
	if !ok || value.Int() != 1 {
		t.Error("expected i to be registered in parent")
	}
}
