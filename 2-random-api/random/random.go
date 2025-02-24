package random

import (
	"math/rand"
	"strconv"
	"time"
)

func Random() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNum := r.Intn(6) + 1

	return strconv.Itoa(randomNum)
}
