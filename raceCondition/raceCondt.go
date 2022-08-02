/*
Create a sample program that simulates a race condition. So it should run fine when run sequentially but 
cause trouble when run in two or more goroutines.
*/


package main

import (
	"fmt"
)

var value int = 50

func increase(a int, b int) {
	fmt.Printf("%d: Value before increase: %d \n", a, value)
	value = value + b
	fmt.Printf("%d: Value after increase: %d \n", a, value)
}

//Value can not go below 50
func decrease(a int, b int) {
	fmt.Printf("%d: Value before decrease %d\n", a, value)
	if value - a < 50 {
		fmt.Printf("")
	} else {
		value = value - b
		fmt.Printf("%d: Value after decrease %d\n", a, value)

	}
	fmt.Printf("%d: Value after heating up: %d \n", b, value)
}

func main() {
	fmt.Printf("Value at the beginning: %d\n", value)
	for a := 1; a <= 10; a++ { 
		go decrease(a, 20)
		go increase(a, 20)
	}
	fmt.Printf("Value at the end: %d\n", value)
}

/* Output
Value at the beginning: 50
1: Value before decrease 50
20: Value after heating up: 50 
4: Value before decrease 50
20: Value after heating up: 50 
3: Value before decrease 50
20: Value after heating up: 50 
3: Value before increase: 50 
3: Value after increase: 70 
6: Value before increase: 70 
6: Value after increase: 90 
Value at the end: 90
4: Value before increase: 90 
4: Value after increase: 110 
5: Value before decrease 110
*/