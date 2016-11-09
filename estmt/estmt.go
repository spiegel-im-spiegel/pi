package estmt

import (
	"fmt"
	"io"
	"math"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/pi/genpi"
)

//Context is context for plot package.
type Context struct {
	ui            *gocli.UI
	pointCount    int64
	estimateCount int64
	histClass     float64
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, pc, ec int64, hc float64) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), pointCount: pc, estimateCount: ec, histClass: hc}
}

//Execute returns estimate of Pi.
func Execute(cxt *Context) error {
	if cxt.pointCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for pcount option", cxt.pointCount)
	}
	if cxt.estimateCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for ecount option", cxt.estimateCount)
	}

	ch := genpi.New(cxt.pointCount, cxt.estimateCount)
	min := float64(10)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	pis := make([]float64, 0, cxt.estimateCount)
	for pi := range ch {
		pis = append(pis, pi)
		if pi < min {
			min = pi
		}
		if pi > max {
			max = pi
		}
		sum += pi
		sum2 += pi * pi
	}
	cxt.ui.OutputErrln(fmt.Sprintf("minimum value: %7.5f", min))
	cxt.ui.OutputErrln(fmt.Sprintf("maximum value: %7.5f", max))
	ave := sum / float64(cxt.estimateCount)
	cxt.ui.OutputErrln(fmt.Sprintf("average value: %7.5f", ave))
	vari := sum2/float64(cxt.estimateCount) - ave*ave
	cxt.ui.OutputErrln(fmt.Sprintf("standard deviation: %7.5f", math.Sqrt(vari)))

	if cxt.histClass <= 0.0 {
		for _, pi := range pis {
			cxt.ui.Outputln(fmt.Sprintf("%7.5f", pi))
		}
		return nil
	}

	classCount := int((max-min)/cxt.histClass) + 1
	freqs := make([]int, classCount)
	for _, pi := range pis {
		class := 0
		for i := 0; i < classCount; i++ {
			if min+cxt.histClass*float64(i) <= pi {
				class = i
			} else {
				break
			}
		}
		freqs[class]++
	}
	for i, freq := range freqs {
		mid := min + cxt.histClass*float64(i) + cxt.histClass/2.0
		cxt.ui.Outputln(fmt.Sprintf("%7.5f\t%v", mid, freq))
	}

	return nil
}
