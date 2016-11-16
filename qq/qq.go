/**
 * These codes are licensed under CC0 (exclude qnorm function).
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package qq

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/gocli"
)

//Context is context for estmt package.
type Context struct {
	ui *gocli.UI
}

//NewContext returns Context instance
func NewContext(r io.Reader, w, e io.Writer) *Context {
	return &Context{ui: gocli.NewUI(r, w, e)}
}

//Execute output Q-Q plot data.
func Execute(cxt *Context) error {
	scanner := bufio.NewScanner(cxt.ui.Reader())
	pis := make([]float64, 0)
	for scanner.Scan() {
		pi, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return errors.Wrap(err, "invalid data")
		}
		pis = append(pis, pi)
	}
	ecf := float64(len(pis))

	sort.Float64s(pis)
	for i, pi := range pis {
		rank := (float64(i+1) - 0.5) / ecf
		cxt.ui.Outputln(fmt.Sprintf("%v\t%v", qnorm(rank), pi))
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
