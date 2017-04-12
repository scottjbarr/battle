package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/scottjbarr/battle"
)

func main() {
	count, err := strconv.ParseInt(os.Getenv("COUNT"), 10, 64)

	if err != nil {
		fmt.Printf("COUNT not set, defaulting to 10\n")
		count = 10
	}

	rand.Seed(time.Now().Unix())

	b := battle.NewBattle(count)

	for !b.Finished() {
		b.Turn()
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("%s wins!\n", b.Winner())
}
