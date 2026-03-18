package main

import (
	"fmt"
	"math/rand"

	"github.com/Uj3nt/go-bame-gen"
)

func main() {
	name := bame.GenBameFromFile(rand.Intn(7)+3, "japan.txt")
	fmt.Printf("Hello, %s!\n", name)
}
