package main

import (
	"bytes"
	"sync"
)

var buffers  = sync.Pool{
	New :func() interface {} {
		return new(bytes.Buffer)
	},
}
func main() {
	p := sync.Pool(new(bytes.Buffer))
}