package injekt

import (
	"reflect"
	"testing"
)

type (
	Func func(a int) int

	A struct{ a string }
)

func newInjector() *Injector {
	var f Func
	injector := New(f)
	var a = A{"a"}
	injector.Register(a)
	return injector
}

func TestServiceInjector_Register(t *testing.T) {
	injector := newInjector()
	var a A
	if _, ok := injector.services[reflect.TypeOf(a)]; !ok {
		t.Error("a is not registered")
	}
}

func TestInjector_invoke(t *testing.T) {
	injector := newInjector()
	injector.invoke(func(a A) {
		if a.a != "a" {
			t.Error("invoke failed")
		}
	})

	injector.invoke(func(a A, b int) {
		if a.a != "a" {
			t.Error("invoke failed")
		}
		if b != 0 {
			t.Error("b should be 0")
		}
	})

	defer func() {
		if err := recover(); err == nil {
			t.Error("panic expected")
		}
	}()
	injector.invoke(nil)
}

func TestInjector_Wrap(t *testing.T) {
	injector := newInjector()
	injector.Register("Hello world")
	injector.Wrap(func(a A, n int, b string) {
		if b != "Hello world" {
			t.Errorf("Expected %s found %s", "Hello world", b)
		}
		if n != 10 {
			t.Errorf("Expected %d found %d", 10, n)
		}
		exp := A{a: "a"}
		if a != exp {
			t.Errorf("Expected %v found %v", exp, a)
		}
	}).(Func)(10)
}

func TestInjector_WrapTo(t *testing.T) {
	var a A
	injector := New(a)
	injector.Register("Something")
	var f func(int) int
	if num := injector.WrapTo(func(n int, s string) int {
		return n
	}, f).(func(int) int)(10); num != 10 {
		t.Errorf("Expected %d found %d", 10, num)
	}
	defer func() {
		if err := recover(); err == nil {
			t.Error("panic expected")
		}
	}()
	injector.WrapTo(nil, nil)
}
