
IMPORTANT:
==========



Tech note 1:
------------

The original files decCommon.c and decBasic.c HAVE BEEN MODIFIED, so that they are not compiled by cgo.

In the first line of these two files, this line has been added:

// +build ignore

followed by a blank line.

If not, "go build" will try to compile these two files, and errors occur.
This is because these files are "imported" by files like decQuad.c.
They are not meant to be compiled stand-alone, but inside other decQuad.c, etc.





Tech note 2:
------------

By typing "go build -x", we have the confirmation that the proper flags for multithreading are used ( -fPIC and -pthread ), as shown in the following line from the output:

	...
	gcc -I . -g -O2 -fPIC -m64 -pthread -I $WORK/decnum/_obj/ -o $WORK/decnum/_obj/decNumber.o -c ./decNumber.c
	...
	...

