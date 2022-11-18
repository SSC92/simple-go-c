package simplecgo

/*
#include "./src/myc.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

//export MyCFunction
func MyCFunction(x int) int {
	//calls myc.c function and returns the result
	result := C.myFunction(C.int(x))
	return int(result)
}
