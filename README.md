# go-function-call-counter
Go Function Call Counter


## Installing
```
go get github.com/zhorifiandi/go-function-call-counter
```

## Usage
Sample: https://go.dev/play/p/DPSIjK-RJWm
```go
package main

import (
	"fmt"

	gofunctioncallcounter "github.com/zhorifiandi/go-function-call-counter"
)

func myFunction(a int, b int) int {
	return a + b
}

func main() {
	// Create a new CountCalls object for the function
	cc := gofunctioncallcounter.NewFunctionCallCounter(myFunction)

	// Call Function using original function
	resultFromPlainFunction := myFunction(1, 2)

	// Get wrapped function
	wrappedFn := cc.GetFunction().(func(int, int) int)

	// Call Function using wrapped function
	resultFromWrappedFunction := wrappedFn(1, 2)

	fmt.Printf("resultFromPlainFunction: %d\n", resultFromPlainFunction)
	fmt.Printf("resultFromWrappedFunction: %d\n", resultFromWrappedFunction)

	fmt.Printf("Function called %d times\n", cc.GetCounter())

	// Call the function for several times
	wrappedFn(4, 7)
	wrappedFn(5, 8)
	wrappedFn(6, 9)
	fmt.Printf("Function called %d times\n", cc.GetCounter())

	// Resetting Counter
	cc.ResetCounter()
	fmt.Printf("Function called %d times\n", cc.GetCounter())
}

```
