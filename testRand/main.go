package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i < 10; i++ {
		newInt := fmt.Sprintf("%02d", rand.Intn(10)+1)
		fmt.Println(newInt)
	}
}
