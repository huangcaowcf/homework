package main

import "unsafe"

func main() {
	var x1 byte
	var x int
	var y int32 //  (-2^16)--(2^16-1)
	var z int64
	var h uint
	var i uint32 //  (0-2^32-1)
	var j uint64
	var i8 int8
	var ui8 uint8
	println(x1, x, y, z, h, i, j)
	println(unsafe.Sizeof(x1))
	println(unsafe.Sizeof(x))
	println(unsafe.Sizeof(y))
	println(unsafe.Sizeof(z))
	println(unsafe.Sizeof(h))
	println(unsafe.Sizeof(i))
	println(unsafe.Sizeof(j))
	println(unsafe.Sizeof(i8))

	i8 = 127
	i8 = i8 + 2
	ui8 = 255
	ui8 = ui8 + 1
	println(i8, ui8)
}
