package gofunctioncallcounter_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	gofunctioncallcounter "github.com/zhorifiandi/go-function-call-counter"
)

func myFunction(a int, b int) int {
	return a + b
}
func TestFunctionCallCounter(t *testing.T) {
	// Create a new CountCalls object for the function
	cc, err := gofunctioncallcounter.NewFunctionCallCounter(myFunction)
	require.NoError(t, err)
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

func myRaceFunction(wg *sync.WaitGroup, a int, b int) int {
	defer wg.Done()
	return myFunction(a, b)
}

func TestRaceFunctionCallCounter(t *testing.T) {
	// Create a new CountCalls object for the function
	cc, err := gofunctioncallcounter.NewFunctionCallCounter(myRaceFunction)
	require.NoError(t, err)

	// Get wrapped function
	wrappedFn := cc.GetFunction().(func(*sync.WaitGroup, int, int) int)

	var wg sync.WaitGroup

	shouldBeCalled := 50
	for i := 0; i < shouldBeCalled; i++ {
		wg.Add(1)
		go wrappedFn(&wg, i, shouldBeCalled)
	}

	wg.Wait()
	require.Equal(t, shouldBeCalled, cc.GetCounter())
}
