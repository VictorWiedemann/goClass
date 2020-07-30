//iota is an incrementing non-typed integer
package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3
	)

	//even though we already used iota, this will restart at 0.
	const (
		north = iota
		east
		south
		west
	)

	//must be used in brackets, this doesn't weork
	//var1 := iota
	//var2 := iota

	//c1 and c2 are explicitly incremented, c3 is incemented due to being equal to c2 (which is iota)
	fmt.Println(c1, c2, c3)
	fmt.Println(north, east, south, west)

	const (
		a = (iota * 2) + 1
		b
		c
	)
	fmt.Println(a, b, c) //1, 3, 5

	//we want x -2, y -4 and z-5
	const (
		x = -(iota + 2)
		_
		y
		z
	)
	println(x, y, z)

}
