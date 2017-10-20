# decnum

This is a Go binding around C decNumber package, for calculation with base-10 floating point numbers.
Decimal data type is important for financial calculations.

The C decNumber package can be found at:
http://speleotrove.com/decimal/

The decNumber package used is "International Components for Unicode (ICU)" version.


### Documentation of original C decNumber package

http://speleotrove.com/decimal/decnumber.html

More specifically, you should read these topics:
   - Context: http://speleotrove.com/decimal/dncont.html
   - decQuad: http://speleotrove.com/decimal/dnfloat.html
   - decQuad example: http://speleotrove.com/decimal/dnusers.html#example7


The original C decNumber package contains two kinds of data type:
   - decNumber, which contains arbitrary-precision numbers. Storage will grow as needed.
   - decQuad, decDouble, decSingle, which are fixed-size data types. They are faster than decNumber.


### This Go package

  The Quad type contains a 128 bits decimal floating point value, and a 16 bits status field.

__This Go package only uses the decQuad data type__ (128 bits), which can store numbers with 34 significant digits.
It is very much like the float64, except that its precision is better (float64 has a precision of only 15 digits), and it works in base-10 instead of base-2.

I have only written the following files:
   - [mydecquad.c](https://github.com/covrom/decnum/blob/master/mydecquad.c)
   - [mydecquad.h](https://github.com/covrom/decnum/blob/master/mydecquad.h)
   - [mydecquad.go](https://github.com/covrom/decnum/blob/master/mydecquad.go)
   - [mydecquad_test.go](https://github.com/covrom/decnum/blob/master/mydecquad_test.go)
   - [mydecquad_run_cowlishaw_test.go](https://github.com/covrom/decnum/blob/master/mydecquad_run_cowlishaw_test.go)
   - [doc.go](https://github.com/covrom/decnum/blob/master/doc.go)

The other .c and .h files in the directory come from the original C decNumber package.

The code of this Go wrapper is quite easy to read, and the pattern for calling C function is always the same.
Parameters are always passed from Go to C as value, and in the other direction too.
Strings passed from C to Go are are also passed by value, as array in struct.
Strings passed from Go to C using C.CString(s).


__Installation__:

    go get github.com/covrom/decnum


### Godoc
https://godoc.org/github.com/covrom/decnum


### Test
The test file [mydecquad_test.go](https://github.com/covrom/decnum/blob/master/mydecquad_test.go) is very instructive.

Run the test:

    go test


### License

All files in this package are under ICU License (see ICU-license.html).



