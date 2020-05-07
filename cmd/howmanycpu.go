package cmd

import (
	"fmt"
	"runtime"
)

func Howmanycpu() {
	fmt.Println("CPU:", runtime.NumCPU(), "KERNE")
}