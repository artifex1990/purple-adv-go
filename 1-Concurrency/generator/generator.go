package generator

import (
	"math/rand"
	"time"
)

// Функция генерации слайса с случайными числами произвольного размера
func Generate(n int) []int {
	slice := make([]int, n)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < n; i++ {
		slice[i] = r.Intn(101)
	}

	return slice
}
