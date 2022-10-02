package main

// #include <stdio.h>
// #include <stdlib.h>
//
// typedef struct {
// int x, y;
// } numbers;
// numbers *num;
// int sum(numbers *n) {
//     return n->x + n->y;
// }
import "C"
import (
	"fmt"
)

func main() {
	// Allocated in Golang memory
	goMemory := &C.numbers{x: 1, y: 3}
	fmt.Println("Calling for Sum ", C.sum(goMemory))
}
