package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 0x0100
	ptr := unsafe.Pointer(&i)
	if 0x01 == *(*byte)(ptr) {
		fmt.Println("Big Endian")
	} else if 0x00 == *(*byte)(ptr) {
		fmt.Println("Little Endian")
	} else {
		// ...
		println("over")
	}
}
