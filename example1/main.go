package main

// #include <stdio.h>
// #include <stdlib.h>
// void Print(char *name) {
//	 printf("Name: %s ", name);
// }
import "C"
import (
	"unsafe"
)

func main() {
	// C.CString coverts Go to C strings and this conversions make a copy of data in C memory
	name := C.CString("John")
	C.Print(name)
	// Since Memory allocation in C are not known to Go's memory manager,
	// Always free the C memory manually when you are done
	C.free(unsafe.Pointer(name))
}
