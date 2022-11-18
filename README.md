# simple-go-c
Example for using C in Golang

## How to use

This is the simplest example for using C in Golang.

It consists of two parts:

### C part

```myc.c``` contains the C code.
```myc.h``` contains the header file for the C code.

You should replace the ```myc.c``` and ```myc.h``` with your own C code and header file.

### Golang part

```simplescgo.go``` contains the Golang code.

You should replace the include in ```simplescgo.go```  with your actual C code as you can see in the example. You can also use standard C library in your Golang code.

```
/*
#include "myc.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
```