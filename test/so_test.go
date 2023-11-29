package main_test

import (
	"testing"
	"unsafe"

	so "github.com/238Studio/child-nodes-hex-loader-so"
)

func TestSo(t *testing.T) {
	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = uintptr(unsafe.Pointer(&a))
	app := so.InitSoLoader()
	hexPackage, err := app.LoadBinPackage("./test")
	if err != nil {
		println(err.Error())
		return
	}
	re := make([]uintptr, 1)
	var str = "外部"
	println(&str)
	re[0] = (uintptr)(unsafe.Pointer(&str))
	err = hexPackage.Execute("Test1", args, uintptr(unsafe.Pointer(&re)))
	//	re := execute[0]
	println("mew")
	println(*(*string)(unsafe.Pointer(re[0])))
}
