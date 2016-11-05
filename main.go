package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spiegel-im-spiegel/pi/gencmplx"
)

func main() {
	c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), int64(10000))
	for p := range c {
		fmt.Printf("%v\t%v\n", real(p), imag(p))
	}
}
