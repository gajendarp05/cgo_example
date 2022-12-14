package main

// #include <stdio.h>
// #include <stdlib.h>
//
// typedef struct {
//		int a, b;
// } numbers;
//
// int sum(numbers *n) {
//		return n->a + n->b;
// }
import "C"
import "fmt"

func main() {
	// Allocated to C memory
	cMemory := C.calloc(1, C.sizeof_numbers)
	// Explicit call free to C memory
	defer C.free(cMemory)
	n := (*C.numbers)(cMemory)
	n.a = 10
	n.b = 20
	fmt.Println(C.sum(n))
}
