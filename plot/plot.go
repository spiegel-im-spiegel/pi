/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package plot

import (
	"fmt"
	"io"
	"time"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/pi/gencmplx"
)

//Context is context for plot package.
type Context struct {
	ui         *gocli.UI
	rngType    gencmplx.RNGs
	pointCount int64
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, rng gencmplx.RNGs, ct int64) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), rngType: rng, pointCount: ct}
}

//Execute output list of points.
func Execute(cxt *Context) error {
	if cxt.pointCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for pcount option", cxt.pointCount)
	}
	rng := "default"
	switch cxt.rngType {
	case gencmplx.LCG:
		rng = "LCG"
	case gencmplx.MT:
		rng = "MT"
	default:
	}
	cxt.ui.OutputErrln(fmt.Sprintf("random number generator: %s", rng))

	ch := gencmplx.New(gencmplx.NewRndSource(cxt.rngType, time.Now().UnixNano()), cxt.pointCount)
	for p := range ch {
		cxt.ui.Outputln(fmt.Sprintf("%v\t%v", real(p), imag(p)))
	}

	return nil
}
