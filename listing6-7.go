//Listing 6.7: Multiplication first
package main

import (
	"fmt"
)

func main() {
	celsius := 21.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "° F")
}
