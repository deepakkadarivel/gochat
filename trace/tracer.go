package trace

import (
	"io"
	"fmt"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type niltracer struct {}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *niltracer)Trace(a ...interface{})  {
	
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &niltracer{}
}
