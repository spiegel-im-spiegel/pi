/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package gencmplx

import "math/rand"

//New returns generator of random complex number
func New(s rand.Source, count int64) <-chan complex128 {
	ch := make(chan complex128)
	r := rand.New(s)
	go func(r *rand.Rand, count int64) {
		for i := int64(0); i < count; i++ {
			ch <- complex(float64(r.Int63n(10000001))/float64(10000000), float64(r.Int63n(10000001))/float64(10000000))
		}
		close(ch)
	}(r, count)

	return ch
}
