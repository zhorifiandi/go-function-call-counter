package gofunctioncallcounter

import (
	"errors"
	"reflect"
	"sync"
)

type FunctionCallCounter struct {
	Mutex          *sync.Mutex
	FnValue        reflect.Value
	WrappedFnValue reflect.Value
	Counter        int
}

func NewFunctionCallCounter(Fn interface{}) (*FunctionCallCounter, error) {
	f := &FunctionCallCounter{
		Mutex: new(sync.Mutex),
	}

	f.FnValue = reflect.ValueOf(Fn)
	if f.FnValue.Kind() != reflect.Func {
		return nil, errors.New("input value is not a function")
	}

	fnType := f.FnValue.Type()
	f.WrappedFnValue = reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		f.Mutex.Lock()
		f.Counter++
		f.Mutex.Unlock()

		result := f.FnValue.Call(args)

		return result
	})

	return f, nil
}

func (f *FunctionCallCounter) GetFunction() interface{} {
	// Use a type assertion to convert the wrapped function value to a function of the same type as the input function
	return f.WrappedFnValue.Interface()
}

func (f *FunctionCallCounter) GetCounter() int {
	return f.Counter
}

func (f *FunctionCallCounter) ResetCounter() {
	f.Mutex.Lock()
	f.Counter = 0
	f.Mutex.Unlock()
}
