package main

import "fmt"

func main() {

	//classic const structure
	//must be init when declared.
	//don't need to use this to compile like other vars.
	const days int = 7

	const x, y int = 5, 0
	//consts will throw certain errors at compile time. Example is divide by 0
	//below will fail the compile.
	//fmt.Printf(x / y)

	//this will save -500 to all of the variables below
	const (
		min1 = -500
		min2
		min
	)

	fmt.Println(min1, min2, min)
	const a float64 = 5.1
	//untyped constant is:
	const b = 6.7
	fmt.Printf("%T", b) //prints float64

	//below inits are allowed:
	const c float64 = a * b     //because this is 2 consts
	const str = "Hello" + "Go!" // works because of string literals (I guess?)

	const d = 5 > 10

	//below will throw and error (cannot use 2.2 * x (type int) as type float64 in const initializer)
	//const xx int = 5
	//const yy float64 = 2.2 * x

	//this works though

	const xx = 5 //this is an untyped int type. That allows xx * yy without issues.
	const yy = 4.2
	zz := xx * yy
	fmt.Printf("\n\n%v %T, %T, %T", zz, xx, yy, zz)

	//So we can convert untyped int to int and var without issue
	var i int = xx
	var j float64 = xx
	var p byte = xx

	fmt.Println(i, j, p)
}
