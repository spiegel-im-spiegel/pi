/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */
package genpi

import (
	"math/cmplx"
	"time"

	"github.com/spiegel-im-spiegel/pi/gencmplx"
)

//New returns generator of Pi
func New(pc, ec int64, rng gencmplx.RNGs) <-chan float64 {
	ch := make(chan float64)
	pcf := float64(pc)
	go func(pc, ec int64) {
		for i := int64(0); i < ec; i++ {

			c := gencmplx.New(gencmplx.NewRndSource(rng, time.Now().UnixNano()), pc)
			m := int64(0) // plot in circle
			for p := range c {
				if cmplx.Abs(p) <= float64(1) {
					m++
				}
			}
			ch <- float64(4*m) / pcf
		}
		close(ch)
	}(pc, ec)

	return ch
}
