package estmt

import (
	"fmt"
	"io"

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

	ch := genpi.New(cxt.pointCount, cxt.estimateCount)
	for pi := range ch {
		cxt.ui.Outputln(fmt.Sprintf("%v", pi))
	}

	return nil
}
