package plot

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/pi/gencmplx"
)

//Context is context for plot package.
type Context struct {
	ui         *gocli.UI
	pointCount int64
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, ct int64) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), pointCount: ct}
}

//Execute returns list of points.
func Execute(cxt *Context) error {
	if cxt.pointCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for pcount option", cxt.pointCount)
	}

	ch := gencmplx.New(rand.NewSource(time.Now().UnixNano()), cxt.pointCount)
	for p := range ch {
		cxt.ui.Outputln(fmt.Sprintf("%v\t%v", real(p), imag(p)))
	}

	return nil
}
