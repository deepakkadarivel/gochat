package trace

import (
	"testing"
	"bytes"
)

func TestNew(t *testing.T) {
	var buf = bytes.Buffer{}
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New cannot return an empty value.")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Tracer should not print %s", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer = Off()
	silentTracer.Trace("No tracing")
}


