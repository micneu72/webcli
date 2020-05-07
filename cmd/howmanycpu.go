package cmd

import (
	"fmt"
	"runtime"
)

func cpu() {
	fmt.Println("CPU:", runtime.NumCPU(), "KERNE")
}