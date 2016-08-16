package injekt

import (
	"fmt"
	"reflect"
)

// Injector is a service injector.
type Injector struct {
	services map[reflect.Type]interface{}
	funcType interface{}
}

// New creates a new Injector. funcType is value of the required function
// type to wrap to.
//  // http.HandlerFunc as required function.
//  var f http.HandlerFunc
//  injekt.New(f)
//
//  // func(c *mypackge.Context) as required function.
//  injekt.New(func(c *mypackge.Context){})
func New(funcType interface{}) *Injector {
	return &Injector{
		funcType: funcType,
	}
}

// Wrap wraps f and return a function compatible with
// the required function. f must be a function, otherwise a panic
// occurs.
//
// If f returns values, f and the required function must have same return types
// to get desired behaviour. If any of the services passed as
// parameters to f is not registered, empty value of the type will be
// passed.
//  // http.HandlerFunc required function.
//  http.HandleFunc("/", inj.Wrap(myFunc).(http.HandlerFunc))
//
//  // func(c *mypackage.Context) required function and custom router.
//  myRouter.Handle("/", inj.Wrap(myFunc).(func(c *mypackage.Context)))
func (inj Injector) Wrap(f interface{}) interface{} {
	return inj.wrapTo(f, inj.funcType)
}

// WrapTo is similar to Wrap but returns a function compatible with
// funcType.
func (inj Injector) WrapTo(f interface{}, funcType interface{}) interface{} {
	return inj.wrapTo(f, funcType)
}

func (inj Injector) wrapTo(f interface{}, funcType interface{}) interface{} {
	mustBeFunc(f, funcType)
	fType := reflect.TypeOf(funcType)
	return reflect.MakeFunc(fType, func(args []reflect.Value) (results []reflect.Value) {
		results = make([]reflect.Value, fType.NumOut())
		for i := range results {
			results[i] = reflect.Zero(fType.Out(i))
		}
		inj := inj.copy() // copy to scope
		for i := range args {
			inj.Register(args[i].Interface())
		}
		if rs := inj.invoke(f); reflect.TypeOf(f).NumOut() == fType.NumOut() {
			// preserve type order
			for i := range rs {
				if rs[i].Type() == fType.Out(i) {
					results[i] = rs[i]
				}
			}
		}
		return
	}).Interface()
}

// Register registers a new service. Services are identifiable by their types.
// Multiple services of same type should be grouped into a struct,
// and the struct should be registered instead.
func (inj *Injector) Register(service interface{}) {
	inj.register(service)
}

func (inj *Injector) register(service interface{}) {
	if inj.services == nil {
		inj.services = make(map[reflect.Type]interface{})
	}
	inj.services[reflect.TypeOf(service)] = service
}

// invoke invokes function f. f must be func type.
func (inj Injector) invoke(f interface{}) []reflect.Value {
	args := make([]reflect.Value, reflect.TypeOf(f).NumIn())
	for i := range args {
		argType := reflect.TypeOf(f).In(i)
		if service, ok := inj.services[argType]; ok {
			args[i] = reflect.ValueOf(service)
		} else {
			// set zero value
			args[i] = reflect.Zero(argType)
		}
	}
	return reflect.ValueOf(f).Call(args)
}

func (inj Injector) copy() Injector {
	i := inj
	i.services = make(map[reflect.Type]interface{})
	for k, v := range inj.services {
		i.services[k] = v
	}
	return i
}

func mustBeFunc(f ...interface{}) {
	for i := range f {
		if reflect.TypeOf(f[i]).Kind() != reflect.Func {
			panic(fmt.Errorf("'%v' is not a func type", reflect.TypeOf(f[i])))
		}
	}
}
