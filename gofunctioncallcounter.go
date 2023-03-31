package gofunctioncallcounter

import (
	"fmt"
	"reflect"
)

type FunctionCallCounter struct {
	FnValue        reflect.Value
	WrappedFnValue reflect.Value
	Counter        int
}

func NewFunctionCallCounter(Fn interface{}) *FunctionCallCounter {
	f := &FunctionCallCounter{}

	f.FnValue = reflect.ValueOf(Fn)
	if f.FnValue.Kind() != reflect.Func {
		panic("Input value is not a function")
	}
	fnType := f.FnValue.Type()
	f.WrappedFnValue = reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		f.Counter++
		fmt.Printf("Function called %d times\n", f.Counter)

		result := f.FnValue.Call(args)

		return result
	})

	return f
}

func (f *FunctionCallCounter) GetFunction() interface{} {
	// Use a type assertion to convert the wrapped function value to a function of the same type as the input function
	return f.WrappedFnValue.Interface()
}

func (f *FunctionCallCounter) GetCounter() int {
	return f.Counter
}

func (f *FunctionCallCounter) ResetCounter() {
	f.Counter = 0
}
