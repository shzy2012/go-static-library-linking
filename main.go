package main

import "fmt"

// #cgo CFLAGS: -I./static_libraries
// #cgo LDFLAGS: ./static_libraries/libsum.a
// #include <sum.h>
import "C"

func main() {
	fmt.Printf("result=%v\n", C.sum(1, 10))
}
