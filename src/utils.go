package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func shuffleSliceInPlace(s []int) {
	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func getNumberOfDigits(a int) int {
	return len(strconv.Itoa(a))
}

func printProgressIndicator(shouldStop chan bool) {
	fmt.Print("Solving")
	go func() {
		for {
			select {
			case <-shouldStop:
				fmt.Println()
				return
			default:
				fmt.Print(".")
				time.Sleep(time.Second)
			}
		}
	}()
}
