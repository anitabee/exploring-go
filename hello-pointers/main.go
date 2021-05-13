package main

import "fmt"

// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroVal(iVal int) {
	iVal = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory
// address to the current value at that address. Assigning a value to a dereferenced pointer
// changes the value at the referenced address.
func zeroPtr(iPtr *int) {
	*iPtr = 0 // set iPtr through the pointer
}

// Another way to get a pointer is to use the built-in new function.
// new takes a type as an argument, allocates enough memory to fit a
// value of that type and returns a pointer to it.
func oneVal(iPtr *int) {
	*iPtr = 1 // set iPtr through the pointer
}

func square(a *float64) {
	*a = *a * *a
}

// http://www.golang-book.com/books/intro/8
// Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1)
func swap(x *int, y *int) {
	*x, *y = *y, *x

}

func main() {
	i := 5
	fmt.Println("initial:", i) // print initial value 5

	zeroVal(i)
	fmt.Println("zeroVal:", i) // i is still 5

	zeroPtr(&i)
	fmt.Println("zeroPtr", i) // i is 0

	fmt.Println("pointer:", &i) // print pointer value

	iPtr := new(int)
	oneVal(iPtr)
	fmt.Println("oneVal", *iPtr) // iPtr is 1

	a := 1.5
	square(&a)
	fmt.Println("square", a) // a is 2.25

	x := 1
	y := 2
	swap(&x, &y)
	fmt.Printf("swap x: %v, y: %v\n", x, y) // swap x and y values 

}
