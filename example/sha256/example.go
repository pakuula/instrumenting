package main

import (
	"crypto"
	"time"

	"github.com/pakuula/instrumenting"
	"github.com/pakuula/instrumenting/fields"
)

func main() {
	hf := crypto.SHA256.New()
	data := make([]byte, 1024)
	tracer := instrumenting.NewTracer("example", true, fields.Int("size", 1024))
	defer tracer.Finish()

	{
		// child scope
		tracer := tracer.Child("inner", fields.Bool("child", true))
		hf.Reset()
		tracer.Trace("reset")
		n, err := hf.Write(data)
		tracer.TraceWithError("write", err, fields.Int("num of bytes", n))
		_ = hf.Sum([]byte{})
		tracer.TraceWithError("sum", err)
		tracer.Finish(fields.Time("optional field", time.Now()))
	}
}
