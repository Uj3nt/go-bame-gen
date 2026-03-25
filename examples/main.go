package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/Uj3nt/go-bame-gen"
)

func main() {
	name, err := bame.GenBameFromFile(rand.Intn(7)+3, "japan.txt")
	
	if err != nil {
		log.Printf("%v", err);
		return
	}
	
	fmt.Printf("Hello, %s!\n", name)
}
