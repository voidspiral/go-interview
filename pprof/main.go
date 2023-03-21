package main

import (
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	trace.Start(f)
	defer trace.Stop()
}
