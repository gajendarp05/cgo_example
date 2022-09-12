package main

// #include <stdio.h>
// #include <stdlib.h>
//
// typedef struct {
// int x, y;
// } numbers;
// numbers *num;
// void storeRef(numbers *n) {
//     num = n;
// }
// int sum() {
//     return num->x + num->y;
// }
import "C"
import (
	"fmt"
)

func main() {
	passMemory()
	callingSum()
}
func passMemory() {
	goMemory := &C.numbers{x: 1, y: 3}
	C.storeRef(goMemory)
}
func callingSum() {
	fmt.Println("Calling for Sum ", C.sum())
}
