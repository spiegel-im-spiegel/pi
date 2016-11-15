/**
 * These codes are licensed under CC0 (exclude qnorm function).
 * http://creativecommons.org/publicdomain/zero/1.0/
 */
package estmt

import (
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/pi/genpi"
)

//Context is context for estmt package.
type Context struct {
	ui            *gocli.UI
	pointCount    int64
	estimateCount int64
	qqFlag        bool
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, pc, ec int64, qq bool) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), pointCount: pc, estimateCount: ec, qqFlag: qq}
}

//Execute output list of estimate data.
func Execute(cxt *Context) error {
	if cxt.pointCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for pcount option", cxt.pointCount)
	}
	if cxt.estimateCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for ecount option", cxt.estimateCount)
	}
	ecf := float64(cxt.estimateCount)

	//measurement
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

	//statistics
	cxt.ui.OutputErrln(fmt.Sprintf("minimum value: %7.5f", min))
	cxt.ui.OutputErrln(fmt.Sprintf("maximum value: %7.5f", max))
	ave := sum / ecf
	cxt.ui.OutputErrln(fmt.Sprintf("average value: %7.5f", ave))
	devi := math.Sqrt(sum2/ecf - ave*ave) //variance -> standard deviation
	ct := int64(0)
	for _, pi := range pis {
		if ave-devi <= pi && pi <= ave+devi {
			ct++
		}
	}
	cxt.ui.OutputErrln(fmt.Sprintf("standard deviation: %7.5f (%4.1f%%)", devi, float64(ct)*100.0/ecf))

	if !cxt.qqFlag {
		//output
		for _, pi := range pis {
			cxt.ui.Outputln(fmt.Sprintf("%v", pi))
		}
		return nil
	}

	//Q-Q plot
	sort.Float64s(pis)
	//rank := make([]float64, 0, cxt.estimateCount)
	ppnds := make([]float64, 0, cxt.estimateCount)
	for i, _ := range pis {
		r := (float64(i+1) - 0.5) / ecf
		//rank = append(rank, r)
		ppnds = append(ppnds, qnorm(r))
	}

	//output
	for i, pi := range pis {
		//cxt.ui.Outputln(fmt.Sprintf("%v\t%v\t%v", pi, rank[i], ppnds[i]))
		cxt.ui.Outputln(fmt.Sprintf("%v\t%v", ppnds[i], pi))
	}
	return nil
}

//qnorm function refers to http://rangevoting.org/Qnorm.html
// This function is licensed under GNU GPL version 3 or later.
func qnorm(p float64) float64 {
	const (
		split = 0.42
		a0    = 2.50662823884
		a1    = -18.61500062529
		a2    = 41.39119773534
		a3    = -25.44106049637
		b1    = -8.47351093090
		b2    = 23.08336743743
		b3    = -21.06224101826
		b4    = 3.13082909833
		c0    = -2.78718931138
		c1    = -2.29796479134
		c2    = 4.85014127135
		c3    = 2.32121276858
		d1    = 3.54388924762
		d2    = 1.63706781897
	)

	q := p - 0.5
	ppnd := float64(0)
	if math.Abs(q) <= split {
		r := q * q
		ppnd = q * (((a3*r+a2)*r+a1)*r + a0) / ((((b4*r+b3)*r+b2)*r+b1)*r + 1)
	} else {
		r := p
		if q > 0 {
			r = 1 - p
		}
		if r > 0 {
			r = math.Sqrt(-math.Log(r))
			ppnd = (((c3*r+c2)*r+c1)*r + c0) / ((d2*r+d1)*r + 1)
			if q < 0 {
				ppnd = -ppnd
			}
		}
	}
	return ppnd
}
