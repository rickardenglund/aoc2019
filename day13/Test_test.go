package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_test(t *testing.T) {
	m := sync.Map{}

	m.Store(pos{1, 2}, 2)

	v, _ := m.Load(pos{1, 2})
	fmt.Printf("%v\n", v)
}
