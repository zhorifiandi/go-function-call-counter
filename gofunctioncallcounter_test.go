package gofunctioncallcounter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	gofunctioncallcounter "github.com/zhorifiandi/go-function-call-counter"
)

func myFunction(a int, b int) int {
	return a + b
}
func TestFunctionCallCounter(t *testing.T) {
	// Create a new FunctionCallCounter object for the function
	cc := gofunctioncallcounter.NewFunctionCallCounter(myFunction)

	require.Equal(t, 0, cc.GetCounter())

	// Call Function using original function
	resultFromPlainFunction := myFunction(1, 2)

	// Get wrapped function
	wrappedFn := cc.GetFunction().(func(int, int) int)

	// Call Function using wrapped function
	resultFromWrappedFunction := wrappedFn(1, 2)

	require.Equal(t, resultFromPlainFunction, resultFromWrappedFunction)
	require.Equal(t, 1, cc.GetCounter())

	wrappedFn(3, 4)
	require.Equal(t, 2, cc.GetCounter())

	cc.ResetCounter()
	require.Equal(t, 0, cc.GetCounter())
}
