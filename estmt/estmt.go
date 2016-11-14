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
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, pc, ec int64) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), pointCount: pc, estimateCount: ec}
}

//Execute returns list of points.
func Execute(cxt *Context) error {
	if cxt.pointCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for pcount option", cxt.pointCount)
	}
	if cxt.estimateCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for ecount option", cxt.estimateCount)
	}

	//measurement
	ch := genpi.New(cxt.pointCount, cxt.estimateCount)
	min := float64(10)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	pis := make([]float64, 0, cxt.estimateCount)
	for pi := range ch {
		cxt.ui.Outputln(fmt.Sprintf("%v", pi))
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

	//statistics
	cxt.ui.OutputErrln(fmt.Sprintf("minimum value: %7.5f", min))
	cxt.ui.OutputErrln(fmt.Sprintf("maximum value: %7.5f", max))
	ave := sum / float64(cxt.estimateCount)
	cxt.ui.OutputErrln(fmt.Sprintf("average value: %7.5f", ave))
	devi := math.Sqrt(sum2/float64(cxt.estimateCount) - ave*ave) //variance -> standard deviation
	ct := int64(0)
	for _, pi := range pis {
		if ave-devi <= pi && pi <= ave+devi {
			ct++
		}
	}
	cxt.ui.OutputErrln(fmt.Sprintf("standard deviation: %7.5f (%4.1f%%)", devi, float64(ct)*100.0/float64(cxt.estimateCount)))

	return nil
}
