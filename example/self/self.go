package main

import (
	"fmt"

	"github.com/pakuula/instrumenting"
	"github.com/pakuula/instrumenting/fields"
)

func main() {
	target := instrumenting.NewTracer("test specie", true, fields.Bool("ignore", true))

	tracer := instrumenting.NewTracer("self", true)
	for i := 0; i < 1000; i++ {
		target.Trace("trace")
	}
	tracer.Trace("1000 done")
	fmt.Printf("Overhead of a single Trace call, microseconds: %.2f\n", float64(tracer.GetElapsed().Microseconds())/1000.0)
}
