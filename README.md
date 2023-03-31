# go-function-call-counter
Go Function Call Counter


## Installing
```
go get github.com/zhorifiandi/go-function-call-counter
```

## Usage
```go
package main

import (
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
}

```
