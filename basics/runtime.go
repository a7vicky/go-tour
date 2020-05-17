package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT)
	fmt.Println(runtime.GoroutineProfile)
	fmt.Println(runtime.MemProfile)
}
