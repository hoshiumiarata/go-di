package di_test

import (
	"reflect"
	"testing"

	"github.com/hoshiumiarata/go-di"
)

func TestNewContainer(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		container := di.New()
		if container == nil {
			t.Errorf("expected container to not be nil")
		}
		value, ok := container.ResolveType(reflect.TypeOf(container))
		if !ok || value.Interface() != container {
			t.Errorf("expected container to be registered")
		}
	})
	t.Run("New Child", func(t *testing.T) {
		parent := di.New()
		child := parent.New()
		if child == nil {
			t.Errorf("expected child to not be nil")
		}
		if child.Parent() != parent {
			t.Errorf("expected child to have parent")
		}
		value, ok := child.ResolveType(reflect.TypeOf(child))
		if !ok || value.Interface() != child {
			t.Errorf("expected child to be registered")
		}
		value, ok = parent.ResolveType(reflect.TypeOf(parent))
		if !ok || value.Interface() != parent {
			t.Errorf("expected parent to be registered")
		}
	})
}
