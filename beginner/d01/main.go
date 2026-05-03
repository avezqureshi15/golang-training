package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	quotes := []string{
		"I guess some things just take their own time.",
		"It’s okay… not everything needs to be figured out today.",
		"Some days feel different, and that’s fine.",
		"I don’t think too much about it anymore.",
		"Things change quietly most of the time",
	}

	fmt.Println("Avez Q")

	now := time.Now()
	fmt.Println(now)
	
	randNum := rand.Intn(len(quotes) - 1)
	fmt.Println(quotes[randNum])
}
