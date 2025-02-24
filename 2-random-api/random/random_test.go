package random_test

import (
	"go/adv-random-api/random"
	"strconv"
	"testing"
)

func TestRandom(t *testing.T) {
	randNum, _ := strconv.Atoi(random.Random())

	if randNum < 1 || randNum > 6 {
		t.Errorf("Ожидалось 1-6, получили %d", randNum)
	}
}
