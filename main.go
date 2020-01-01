package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type X struct {
	A int
	B string
}

func NewX() *X {
	x := &X{
		A: 123,
		B: "abc",
	}
	fmt.Printf("before: %v\n", x)
	runtime.SetFinalizer(x, finalizer)
	return x
}

func newXAsIntPtr() int {
	x := NewX()
	xp := unsafe.Pointer(x)
	xup := uintptr(xp)
	return int(xup)
}

func finalizer(x *X) {
	fmt.Printf("finalize for -> %v\n", x)
	return
}

func main() {
	xip := newXAsIntPtr()
	runtime.GC()
	xp := unsafe.Pointer(uintptr(xip))
	x := (*X)(xp)
	fmt.Printf("after: %v\n", x)
}
