package lru

import (
	"log"
	"testing"
)

func TestLru(t *testing.T) {
	h := Constructor(2)
	h.Put(1, 1)
	//h.Get(1)
	h.Put(2, 2)
	log.Println(h.Get(1))
	h.Put(3, 3)
	h.Print()
}
