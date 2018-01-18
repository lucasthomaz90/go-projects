package main

import (
	"fmt"
	"math"
	"time"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println("hello world")

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("\nCASE 3")

	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "short"` in this case.
	f := "short"
	fmt.Println(f)

	fmt.Println("\nCASE 4 - Constant")
	fmt.Println("S = ", s)
	//fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const m = 3e20 / n
	fmt.Println("M = ", m)

	// A numeric constant has no type until it's given
	// one, such as by an explicit cast.
	fmt.Println("M int64 = ", int64(m))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	fmt.Println("N math = ", math.Sin(n))

	fmt.Println("\nCASE 5 - FOR")

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("\nCASE 6 - IF")
	//Here's a basic example.
	if 8%2 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	// You can have an `if` statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in all
	// branches.
	if num := -3; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Note that you don't need parentheses around conditions
	// in Go, but that the braces are required.

	fmt.Println("\nCASE 7 - switch")

	// Here's a basic `switch`.
	k := 2
	fmt.Print("Write ", k, " as ")
	switch k {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type `switch` compares types instead of values.  You
	// can use this to discover the the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool", t)
		case int:
			fmt.Println("I'm an int", t)
		default:
			fmt.Printf("Don't know type %T\n", t)
			fmt.Println("Don't know type ", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	fmt.Println("\nCASE 8 - arrays")

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var o [5]int
	fmt.Println("emp:", o)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	o[4] = 100
	fmt.Println("set:", o)
	fmt.Println("get:", o[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(o))

	// Use this syntax to declare and initialize an array
	// in one line.
	p := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", p)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
