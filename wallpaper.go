package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No file specified. Returning.")
		return
	}
	if _, err := os.Stat(os.Args[1]); err != nil {
		fmt.Println("Can't find file at path specified, ", err)
		return
	}
	path := os.Args[1]
	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()

	systemParametersInfo := user32.MustFindProc("SystemParametersInfoW")
	retCode, _, err := systemParametersInfo.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))), 3)
	if retCode != 1 {
		fmt.Println("failed", err)
	}
}
