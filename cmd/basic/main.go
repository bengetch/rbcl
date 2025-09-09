package main

import (
	"fmt"

	"github.com/bengetch/rbcl"
)

func main() {
	a := rbcl.Ristretto255Bytes()
	fmt.Println(a)
}
