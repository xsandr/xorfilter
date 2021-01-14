package xorfilter

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuse32Basic(t *testing.T) {
	testsize := 1_000_000_000
	keys := make([]uint64, NUM_KEYS)
	for i := range keys {
		keys[i] = rand.Uint64()
	}
	filter, _ := PopulateFuse32(keys)
	for _, v := range keys {
		assert.Equal(t, true, filter.Contains(v))
	}
	falsesize := 1_000_000_000
	matches := 0
	bpv := float64(len(filter.Fingerprints)) * 32.0 / float64(testsize)
	assert.Equal(t, true, bpv < 36.404)
	for i := 0; i < falsesize; i++ {
		v := rand.Uint64()
		if filter.Contains(v) {
			matches++
		}
	}
	fpp := float64(matches) * 100.0 / float64(falsesize)
	fmt.Println("false positive rate ", fpp)
	assert.Equal(t, true, fpp < 0.0000000000000001)
	keys = keys[:1000]
	for trial := 0; trial < 10; trial++ {
		rand.Seed(int64(trial))
		for i := range keys {
			keys[i] = rand.Uint64()
		}
		filter, _ = PopulateFuse32(keys)
		for _, v := range keys {
			assert.Equal(t, true, filter.Contains(v))
		}

	}
}
