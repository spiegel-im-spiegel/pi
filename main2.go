package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"

	"github.com/spiegel-im-spiegel/pi/gencmplx"
)

func main() {
	c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), int64(100000))
	n := int64(0) // total
	m := int64(0) // plot in circle
	for p := range c {
		n++
		if cmplx.Abs(p) <= float64(1) {
			m++
		}
	}
	fmt.Printf("n = %v, m = %v, 4m/n = %v\n", n, m, float64(4*m)/float64(n))
}
