package simplecgo

/*
#include "myc.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

//sample function to use cgo
func MyCFunction(x int) int {
	//calls myc.c function and returns the result
	result := C.myFunction(C.int(x))
	return int(result)
}
