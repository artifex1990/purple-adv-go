package generator_test

import (
	"go/adv-dz1/generator"
	"testing"
)

func TestGenerateLen(t *testing.T) {
	lenSlice := 10
	lenGenSlice := len(generator.Generate(lenSlice))

	if lenSlice != lenGenSlice {
		t.Errorf("Ожидалось %d, получили %d", lenSlice, lenGenSlice)
	}
}
