package domain

import (
	"fmt"
	"runtime"
)

// ShowCores displays number of core in the current CPU
func ShowCores() {
	fmt.Println("Num of cores: ", runtime.NumCPU())
}
