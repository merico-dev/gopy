package main

// #cgo pkg-config: python-3.10-embed
// #include <Python.h>
import "C"

import (
    "fmt"
    "unsafe"
)


func main() {
	defer C.Py_Finalize()
	C.Py_Initialize()

    fooModule := C.PyImport_ImportModule(C.CString("foo"))
    if fooModule == nil {
        panic("Error importing module")
    }
    fmt.Println(fooModule)

	cattr_name := C.CString("hello")
	defer C.free(unsafe.Pointer(cattr_name))
	helloFunc := C.PyObject_GetAttrString(fooModule, cattr_name)
    if helloFunc == nil {
        panic("Error importing function")
    }
    fmt.Println(helloFunc)

    C.PyObject_Call(helloFunc, C.PyTuple_New(0), C.PyDict_New())
}
