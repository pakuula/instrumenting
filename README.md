# Overview

A simple utility for measuring time between invocations.

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
```go
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
```

Output :
```json
{"src":"sha256/example.go:14","loc":"started","scope":"example","size":1024,"time":"2022-01-28T21:25:06.970741402+09:00"}
{"src":"sha256/example.go:19","loc":"started","scope":"example:inner","size":1024,"child":true,"time":"2022-01-28T21:25:06.971008139+09:00"}
{"src":"sha256/example.go:21","loc":"reset","scope":"example:inner","size":1024,"child":true,"elapsed":0.79}
{"src":"sha256/example.go:23","loc":"write","scope":"example:inner","size":1024,"child":true,"elapsed":10.917,"num of bytes":1024}
{"src":"sha256/example.go:25","loc":"sum","scope":"example:inner","size":1024,"child":true,"elapsed":2.43}
{"src":"sha256/example.go:26","loc":"finished","scope":"example:inner","size":1024,"child":true,"elapsed":120.998,"optional field":"2022-01-28T21:25:06.971165004+09:00"}
{"src":"sha256/example.go:28","loc":"finished","scope":"example","size":1024,"elapsed":207.726}
```

# Fields

Instrumenting supports all Zap fields available at the moment of writing. 

In the table below the left column is the name of the field. The right column 
is the signature of field constructor.

|Field Name | Constructor Signature|
|---|---|
|`fields.Object`| `func(string, ObjectMarshaler) Field`|
|`fields.Array`| `func(string, ArrayMarshaler) Field`|
|`fields.Bool`| `func(string, bool) Field`|
|`fields.Boolp`| `func(string, *bool) Field`|
|`fields.Bools`| `func(string, []bool) Field`|
|`fields.Complex128`| `func(string, complex128) Field`|
|`fields.Complex128p`| `func(string, *complex128) Field`|
|`fields.Complex128s`| `func(string, []complex128) Field`|
|`fields.Complex64`| `func(string, complex64) Field`|
|`fields.Complex64p`| `func(string, *complex64) Field`|
|`fields.Complex64s`| `func(string, []complex64) Field`|
|`fields.Float64`| `func(string, float64) Field`|
|`fields.Float64p`| `func(string, *float64) Field`|
|`fields.Float64s`| `func(string, []float64) Field`|
|`fields.Float32`| `func(string, float32) Field`|
|`fields.Float32p`| `func(string, *float32) Field`|
|`fields.Float32s`| `func(string, []float32) Field`|
|`fields.Int`| `func(string, int) Field`|
|`fields.Intp`| `func(string, *int) Field`|
|`fields.Ints`| `func(string, []int) Field`|
|`fields.Int64`| `func(string, int64) Field`|
|`fields.Int64p`| `func(string, *int64) Field`|
|`fields.Int64s`| `func(string, []int64) Field`|
|`fields.Int32`| `func(string, int32) Field`|
|`fields.Int32p`| `func(string, *int32) Field`|
|`fields.Int32s`| `func(string, []int32) Field`|
|`fields.Int16`| `func(string, int16) Field`|
|`fields.Int16p`| `func(string, *int16) Field`|
|`fields.Int16s`| `func(string, []int16) Field`|
|`fields.Int8`| `func(string, int8) Field`|
|`fields.Int8p`| `func(string, *int8) Field`|
|`fields.Int8s`| `func(string, []int8) Field`|
|`fields.String`| `func(string, string) Field`|
|`fields.Stringp`| `func(string, *string) Field`|
|`fields.Strings`| `func(string, []string) Field`|
|`fields.Uint`| `func(string, uint) Field`|
|`fields.Uintp`| `func(string, *uint) Field`|
|`fields.Uints`| `func(string, []uint) Field`|
|`fields.Uint64`| `func(string, uint64) Field`|
|`fields.Uint64p`| `func(string, *uint64) Field`|
|`fields.Uint64s`| `func(string, []uint64) Field`|
|`fields.Uint32`| `func(string, uint32) Field`|
|`fields.Uint32p`| `func(string, *uint32) Field`|
|`fields.Uint32s`| `func(string, []uint32) Field`|
|`fields.Uint16`| `func(string, uint16) Field`|
|`fields.Uint16p`| `func(string, *uint16) Field`|
|`fields.Uint16s`| `func(string, []uint16) Field`|
|`fields.Uint8`| `func(string, uint8) Field`|
|`fields.Uint8p`| `func(string, *uint8) Field`|
|`fields.Binary`| `func(string, []byte) Field`|
|`fields.Uintptr`| `func(string, uintptr) Field`|
|`fields.Uintptrp`| `func(string, *uintptr) Field`|
|`fields.Uintptrs`| `func(string, []uintptr) Field`|
|`fields.Time`| `func(string, time.Time) Field`|
|`fields.Timep`| `func(string, *time.Time) Field`|
|`fields.Times`| `func(string, []time.Time) Field`|
|`fields.Duration`| `func(string, time.Duration) Field`|
|`fields.Durationp`| `func(string, *time.Duration) Field`|
|`fields.Durations`| `func(string, []time.Duration) Field`|
|`fields.NamedError`| `func(string, error) Field`|
|`fields.Errors`| `func(string, []error) Field`|
|`fields.Stringer`| `func(string, fmt.Stringer) Field`|
|`fields.Reflect`| `func(string, interface{}) Field`|
|`fields.Skip`| `func() Field`|
|`fields.Inline`| `func(interface{}) Field`|
