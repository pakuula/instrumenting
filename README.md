# Overview

A very simple utility for measuring time between invocations.

It prints to standard output JSON strings.

Tracer object are assigned to a scope that they include in each output record. 
The method `tracer.Trace(args)` prints the name of current location, microseconds elapsed since any previous calls of the tracer methods and optional data fields.
The library is built atop of `uber-zap` tracer, therefore the optional datafields must be specified using `zap.Field` data structure. See [`zap.Field`](https://pkg.go.dev/go.uber.org/zap#Field) for details.

Elapsed time is printed in microseconds. Please, do not put too much trust into those numbers, they might be slightly inaccurate.

To measure the overhead of the tracer you can use the `self` program from `exapmples`.
```bash
git clone https://github.com/pakuula/instrumenting
cd instrumenting
go build ./example/self/ && ./self
```

It will print a bunch of trace lines and a short summary:
```
Overhead of a single Trace call, microseconds: 6.88
```

# Example

How to install:
```bash
go get -u github.com/pakuula/instrumenting
```

Tracing the code
```
package main

import (
	"crypto"

	i_ "github.com/pakuula/instrumenting"
	"go.uber.org/zap"
)

func main() {
	hf := crypto.SHA256.New()
	data := make([]byte, 1024)
	tracer := i_.NewTracer("example", true)
	defer tracer.Finish()

	{
		// child scope
		tracer := tracer.Child("loop")
		hf.Reset()
		tracer.Trace("reset")
		n, err := hf.Write(data)
		tracer.TraceWithError("write", err, zap.Int("num of bytes", n))
		_ = hf.Sum([]byte{})
		tracer.TraceWithError("sum", err)
		tracer.Finish()
	}
}
```

Output :
```json
{"src":"sha256/example.go:13","loc":"started","scope":"example","time":"2022-01-16T21:02:16.953993516+09:00"}
{"src":"sha256/example.go:18","loc":"started","scope":"example:loop","time":"2022-01-16T21:02:16.954240923+09:00"}
{"src":"sha256/example.go:20","loc":"reset","scope":"example:loop","elapsed":0.637}
{"src":"sha256/example.go:22","loc":"write","scope":"example:loop","elapsed":10.06,"num of bytes":1024}
{"src":"sha256/example.go:24","loc":"sum","scope":"example:loop","elapsed":2.263}
{"src":"sha256/example.go:25","loc":"finished","scope":"example:loop","elapsed":123.055}
{"src":"sha256/example.go:27","loc":"finished","scope":"example","elapsed":175.772}
```

