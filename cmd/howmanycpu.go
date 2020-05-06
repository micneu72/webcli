package cmd
import (
	"fmt"
	"runtime"
)

func howmanycpu() {
	fmt.Println("CPU:", runtime.NumCPU(), "KERNE")
}