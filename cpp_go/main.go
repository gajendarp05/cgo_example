package main

// #cgo CFLAGS:
// #cgo LDFLAGS: -ldl
// #include <stdio.h>
// #include <dlfcn.h>
// #include <stdlib.h>
//
// int adder(void *fn, int x, int y) {
// 	int (*Adder)(int, int);
// 	Adder = fn;
//  return Adder(x, y);
// }
import "C"
import (
	"fmt"
	"log"
)

func main() {
	libHandle := C.dlopen(C.CString("./libadd.so"), C.RTLD_LAZY)
	if libHandle == nil {
		log.Fatal("failed to load\n")
	}

	fnAdd := C.dlsym(libHandle, C.CString("Add"))
	fmt.Println(C.adder(fnAdd, 40, 20))
}
