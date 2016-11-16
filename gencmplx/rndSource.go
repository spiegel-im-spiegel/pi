/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package gencmplx

import (
	"math/rand"

	"github.com/davidminor/gorand/lcg"
	"github.com/seehuhn/mt19937"
)

//RNGs is kind of RNG
type RNGs int

const (
	NULL RNGs = iota
	GO
	LCG
	MT
)

//NewRndSource returns Source of random numbers
func NewRndSource(rng RNGs, seed int64) rand.Source {
	switch rng {
	case LCG:
		return lcg.NewLcg64(uint64(seed))
	case MT:
		mt := mt19937.New()
		mt.Seed(seed)
		return mt
	default:
		return rand.NewSource(seed)
	}
}
