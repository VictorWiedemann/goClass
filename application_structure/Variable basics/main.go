package main

import "fmt"

func main() {
	age := 30
	var otherAge int = 40 // don't need to add int here since it knows that it is an int from the 40
	fmt.Println(age, " ", otherAge)

	//var name = "Dan"

	helloString := "Learning golang!"
	fmt.Println(helloString)

	//declare 2 new variables
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	//this works because year is new even if car is not.
	car, year := "BMW", 2020
	fmt.Println(car, year)

	//all declarations are default empty or 0 like C#
	var i, j int
	i, j = 5, 8

	//this will mute the errors
	//_, _ = i, j

	j, i = i, j // this is how to swap a variable.

	//a, b := 4, 5.2

	//cannot do the below since types do not match. variables are not dynamic.
	//a = b

	//will actually print the string %d, then 20
	fmt.Println("your age is %d", 20)

	//will replace %d with 21
	fmt.Printf("your age is %d\n", 21)

	//to print escapable chars:
	fmt.Printf("hullo, \"Hello!\"\n")

	steveName := "steve"
	//float is %f, int is %d, string is %s, %q for quoted strings. %v can be any type.
	//%T will print the type
	fmt.Printf("hey %q in variable type %T", steveName, steveName)

	//%t is to print a bool. %b is binary base 2.
	//To print at least n bits:
	fmt.Printf("\n\n number in binary %08b", 55)

	//%x is hex

	//to print with the correct amount of numbers:
	fmt.Printf("\n\n%6.8f", 133.2)
}
