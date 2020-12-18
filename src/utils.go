package main

import (
	"math/rand"
	"time"
)

func shuffleSliceInPlace(s []int) {
	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
