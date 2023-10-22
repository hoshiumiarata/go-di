package di_test

import (
	"testing"

	"github.com/hoshiumiarata/go-di"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func (B) F() {}

type I interface {
	F()
}

func TestCall(t *testing.T) {
	t.Run("By Type", func(t *testing.T) {
		f := func(a int, b string, c int8, d float32, e float64, f bool, g A, h *A, i I, j func() bool, k func(int) int) {
			if a != 1 {
				t.Errorf("expected a to be 1, but got %d", a)
			}
			if b != "2" {
				t.Errorf("expected b to be 2, but got %s", b)
			}
			if c != 3 {
				t.Errorf("expected c to be 3, but got %d", c)
			}
			if d != 4 {
				t.Errorf("expected d to be 4, but got %f", d)
			}
			if e != 5 {
				t.Errorf("expected e to be 5, but got %f", e)
			}
			if f != true {
				t.Errorf("expected f to be true, but got %t", f)
			}
			if g.Name != "7" {
				t.Errorf("expected g.Name to be 7, but got %s", g.Name)
			}
			if h.Name != "8" {
				t.Errorf("expected h.Name to be 8, but got %s", h.Name)
			}
			if _, ok := i.(B); !ok {
				t.Errorf("expected i to be B, but got %T", i)
			}
			if !j() {
				t.Errorf("expected j() to be true, but got false")
			}
			if k(10) != 10 {
				t.Errorf("expected k(10) to be 10, but got %d", k(10))
			}
		}

		container := di.New()
		di.RegisterValueTo(container, 1)
		di.RegisterValueTo(container, "2")
		di.RegisterValueTo(container, int8(3))
		di.RegisterValueTo(container, float32(4))
		di.RegisterValueTo(container, float64(5))
		di.RegisterValueTo(container, true)
		di.RegisterValueTo(container, A{Name: "7"})
		di.RegisterValueTo(container, &A{Name: "8"})
		var i I = B{Name: "9"}
		di.RegisterValueTo(container, i)
		di.RegisterValueTo(container, func() bool { return true })
		di.RegisterValueTo(container, func(a int) int { return a })
		container.Call(f)
	})
	t.Run("By Position", func(t *testing.T) {
		f := func(a int, b string, c int8, d float32, e float64, f bool, g A, h *A, i I, j func() bool, k func(int) int) {
			if a != 1 {
				t.Errorf("expected a to be 1, but got %d", a)
			}
			if b != "2" {
				t.Errorf("expected b to be 2, but got %s", b)
			}
			if c != 3 {
				t.Errorf("expected c to be 3, but got %d", c)
			}
			if d != 4 {
				t.Errorf("expected d to be 4, but got %f", d)
			}
			if e != 5 {
				t.Errorf("expected e to be 5, but got %f", e)
			}
			if f != true {
				t.Errorf("expected f to be true, but got %t", f)
			}
			if g.Name != "7" {
				t.Errorf("expected g.Name to be 7, but got %s", g.Name)
			}
			if h.Name != "8" {
				t.Errorf("expected h.Name to be 8, but got %s", h.Name)
			}
			if _, ok := i.(B); !ok {
				t.Errorf("expected i to be B, but got %T", i)
			}
		}

		container := di.New()
		di.RegisterByPositionTo(container, 0, 1)
		di.RegisterByPositionTo(container, 1, "2")
		di.RegisterByPositionTo(container, 2, int8(3))
		di.RegisterByPositionTo(container, 3, float32(4))
		di.RegisterByPositionTo(container, 4, float64(5))
		di.RegisterByPositionTo(container, 5, true)
		di.RegisterByPositionTo(container, -5, A{Name: "7"})
		di.RegisterByPositionTo(container, -4, &A{Name: "8"})
		var i I = B{Name: "9"}
		di.RegisterByPositionTo(container, -3, i)
		di.RegisterByPositionTo(container, -2, func() bool { return true })
		di.RegisterByPositionTo(container, -1, func(a int) int { return a })
		container.Call(f)
	})
	t.Run("Priority", func(t *testing.T) {
		f := func(a int, b int, c int) {
			if a != 1 {
				t.Errorf("expected a to be 1, but got %d", a)
			}
			if b != 2 {
				t.Errorf("expected b to be 2, but got %d", b)
			}
			if c != 3 {
				t.Errorf("expected c to be 3, but got %d", c)
			}
		}

		container := di.New()
		di.RegisterValueTo(container, 2)
		di.RegisterByPositionTo(container, 0, 1)
		di.RegisterByPositionTo(container, 2, 3)
		container.Call(f)
	})
	t.Run("Not Assignable", func(t *testing.T) {
		f := func(a int, b string) {
			if a != 0 {
				t.Errorf("expected a to be 0, but got %d", a)
			}
			if b != "" {
				t.Errorf("expected b to be \"\", but got %s", b)
			}
		}

		container := di.New()
		di.RegisterByPositionTo(container, 0, "1")
		di.RegisterByPositionTo(container, -1, 2)
		container.Call(f)
	})
	t.Run("Not Function", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, but not")
			}
		}()

		container := di.New()
		container.Call(1)
	})
}

func BenchmarkCall(b *testing.B) {
	f := func(a int, b string, c int8, d float32, e float64, f bool, g A, h *A, i I, j func() bool, k func(int) int) {
	}
	b.Run("Without DI", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(1, "2", 3, 4, 5, true, A{Name: "7"}, &A{Name: "8"}, B{Name: "9"}, func() bool { return true }, func(a int) int { return a })
		}
	})
	b.Run("With DI using Type", func(b *testing.B) {
		container := di.New()
		di.RegisterValueTo(container, 1)
		di.RegisterValueTo(container, "2")
		di.RegisterValueTo(container, int8(3))
		di.RegisterValueTo(container, float32(4))
		di.RegisterValueTo(container, float64(5))
		di.RegisterValueTo(container, true)
		di.RegisterValueTo(container, A{Name: "7"})
		di.RegisterValueTo(container, &A{Name: "8"})
		var i I = B{Name: "9"}
		di.RegisterValueTo(container, i)
		di.RegisterValueTo(container, func() bool { return true })
		di.RegisterValueTo(container, func(a int) int { return a })
		for i := 0; i < b.N; i++ {
			container.Call(f)
		}
	})
	b.Run("With DI using Position", func(b *testing.B) {
		container := di.New()
		di.RegisterByPositionTo(container, 0, 1)
		di.RegisterByPositionTo(container, 1, "2")
		di.RegisterByPositionTo(container, 2, int8(3))
		di.RegisterByPositionTo(container, 3, float32(4))
		di.RegisterByPositionTo(container, 4, float64(5))
		di.RegisterByPositionTo(container, 5, true)
		di.RegisterByPositionTo(container, 6, A{Name: "7"})
		di.RegisterByPositionTo(container, 7, &A{Name: "8"})
		var i I = B{Name: "9"}
		di.RegisterByPositionTo(container, 8, i)
		di.RegisterByPositionTo(container, 9, func() bool { return true })
		di.RegisterByPositionTo(container, 10, func(a int) int { return a })
		for i := 0; i < b.N; i++ {
			container.Call(f)
		}
	})
}
