```go
package main

import "fmt"

func main(){
	fmt.Print("nop")
}

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

// Multiply multiplies two integers
func Multiply(x, y int) (res int) {
	return x / y
}
```
