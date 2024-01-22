package main

import (
	"fmt"
)

func main() {
	y := 42 // Go assigns the type as int automatically
	fmt.Printf("%v is of type %T \n", y, y)

	var z float32 = 43.234
	var m float64 = 42.222

	// Here float32 type is converted into float64 type

	1
	m = float64(z)
	fmt.Printf("%v of type %T \n", m, m)

}
