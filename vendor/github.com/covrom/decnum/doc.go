/*
Package decnum is a Go binding around C decNumber package, for calculation with decimal floating point numbers.
Decimal base-10 data type is important for financial calculations.

Godoc: https://godoc.org/github.com/rin01/decnum

The type Quad contains a 128 bit decimal floating point number, and a status.

Quad is immutable, as the value and status it contains are immutable.

Quad is simply passed by value to methods or functions.


Status

The status field in Quad contains all the flags set by all operations that have generated the value.

Use the Error() method to check for errors.

The status field of a Quad returned by any operation contains the combined status of the arguments, as well as the flags set by the operation.
Status flags accumulate and are never cleared. This way, you can make a series of operations, and just check the final result for errors.


Example of use

	package main

	import (
		"fmt"
		"log"
		"os"

		"github.com/rin01/decnum"
	)

	func main() {
		var (
			err error
			a   decnum.Quad  //    uninitialized value is 0e-6176. It is really zero, but with the highest negative exponent for this type.
			b   decnum.Quad  //    If you prefer a variable to be 0, that is, 0e0, do      x = decnum.Zero()
			r   decnum.Quad
		)

		if a, err = decnum.FromString(os.Args[1]); err != nil { // convert string to Quad
			log.Fatalf("ERROR OCCURRED !   %v\n", err)
		}

		if b, err = decnum.FromString(os.Args[2]); err != nil { // err is same as b.Error()
			log.Fatalf("ERROR OCCURRED !   %v\n", err)
		}

		r = a.Add(b) // r = a + b
		// ...
		// you can put other operations here. You can also chain them, e.g. r = a.Div(b).Add(c)
		// ...

		fmt.Println("result =", r.String())

		if err := r.Error(); err != nil { // you can just check the final result status after a series of operations have been done, as status flags accumulate and are never cleared.
			log.Fatalf("ERROR OCCURRED !   %v\n", err)
		}
	}


Internal representation of numbers

It is easier to work with this package if you keep in mind the following representation for numbers:

         (-1)^sign  coefficient * 10^exponent
         where coefficient is an integer storing 34 digits.

         12.345678e2    is     12345678E-4
         123e5          is          123E+5
         0              is            0E+0
         1              is            1E+0
         1.00           is          100E+0
         34.560         is        34560E-3

This representation is important to grasp when using functions like ToIntegral, Quantize, etc.


Represention of numbers for display

When numbers are displayed, the functions that convert them to string like ToString use a different format:

         (-1)^sign  c.oefficient * 10^exp
         where c.oefficient is a fractional number with one digit before fractional point

         1234.567e-12       is printed as     1.234567E-9
         650e4              is printed as         6.50E+6

This representation is well suited for displaying numbers, but not when using functions like ToIntegral, Quantize, etc.


Test

The test file for decnum package is
https://github.com/rin01/decnum/blob/master/mydecquad_test.go

The tests existing in the original decnumber C package have also been run.
https://github.com/rin01/decnum/blob/master/mydecquad_run_cowlishaw_test.go


Tech note

This package uses cgo to call functions in the C decNumber package.

Between Go and C world, all parameters are passed by value, because they are small.
This way, there is no need to make complicated things to pass pointers, and it is as fast, or even faster.

*/
package decnum
