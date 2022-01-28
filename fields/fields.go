package fields

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field
type ObjectMarshaler = zapcore.ObjectMarshaler
type ObjectMarshalerFunc = zapcore.ObjectMarshalerFunc
type ArrayMarshaler = zapcore.ArrayMarshaler
type ArrayMarshalerFunc = zapcore.ArrayMarshalerFunc

var (
	Object      = zap.Object
	Array       = zap.Array
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	Bools       = zap.Bools
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex128s = zap.Complex128s
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Complex64s  = zap.Complex64s
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float64s    = zap.Float64s
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Float32s    = zap.Float32s
	Int         = zap.Int
	Intp        = zap.Intp
	Ints        = zap.Ints
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int64s      = zap.Int64s
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int32s      = zap.Int32s
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int16s      = zap.Int16s
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	Int8s       = zap.Int8s
	String      = zap.String
	Stringp     = zap.Stringp
	Strings     = zap.Strings
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uints       = zap.Uints
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint64s     = zap.Uint64s
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint32s     = zap.Uint32s
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint16s     = zap.Uint16s
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Binary      = zap.Binary
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Uintptrs    = zap.Uintptrs
	Time        = zap.Time
	Timep       = zap.Timep
	Times       = zap.Times
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Durations   = zap.Durations
	NamedError  = zap.NamedError
	Errors      = zap.Errors
	Stringer    = zap.Stringer
	Reflect     = zap.Reflect

	Skip   = zap.Skip
	Inline = zap.Inline
)

// type fields struct {
// 	Object      func(string, ObjectMarshaler) Field
// 	Array       func(string, ArrayMarshaler) Field
// 	Bool        func(string, bool) Field
// 	Boolp       func(string, *bool) Field
// 	Bools       func(string, []bool) Field
// 	Complex128  func(string, complex128) Field
// 	Complex128p func(string, *complex128) Field
// 	Complex128s func(string, []complex128) Field
// 	Complex64   func(string, complex64) Field
// 	Complex64p  func(string, *complex64) Field
// 	Complex64s  func(string, []complex64) Field
// 	Float64     func(string, float64) Field
// 	Float64p    func(string, *float64) Field
// 	Float64s    func(string, []float64) Field
// 	Float32     func(string, float32) Field
// 	Float32p    func(string, *float32) Field
// 	Float32s    func(string, []float32) Field
// 	Int         func(string, int) Field
// 	Intp        func(string, *int) Field
// 	Ints        func(string, []int) Field
// 	Int64       func(string, int64) Field
// 	Int64p      func(string, *int64) Field
// 	Int64s      func(string, []int64) Field
// 	Int32       func(string, int32) Field
// 	Int32p      func(string, *int32) Field
// 	Int32s      func(string, []int32) Field
// 	Int16       func(string, int16) Field
// 	Int16p      func(string, *int16) Field
// 	Int16s      func(string, []int16) Field
// 	Int8        func(string, int8) Field
// 	Int8p       func(string, *int8) Field
// 	Int8s       func(string, []int8) Field
// 	String      func(string, string) Field
// 	Stringp     func(string, *string) Field
// 	Strings     func(string, []string) Field
// 	Uint        func(string, uint) Field
// 	Uintp       func(string, *uint) Field
// 	Uints       func(string, []uint) Field
// 	Uint64      func(string, uint64) Field
// 	Uint64p     func(string, *uint64) Field
// 	Uint64s     func(string, []uint64) Field
// 	Uint32      func(string, uint32) Field
// 	Uint32p     func(string, *uint32) Field
// 	Uint32s     func(string, []uint32) Field
// 	Uint16      func(string, uint16) Field
// 	Uint16p     func(string, *uint16) Field
// 	Uint16s     func(string, []uint16) Field
// 	Uint8       func(string, uint8) Field
// 	Uint8p      func(string, *uint8) Field
// 	Binary      func(string, []byte) Field
// 	Uintptr     func(string, uintptr) Field
// 	Uintptrp    func(string, *uintptr) Field
// 	Uintptrs    func(string, []uintptr) Field
// 	Time        func(string, time.Time) Field
// 	Timep       func(string, *time.Time) Field
// 	Times       func(string, []time.Time) Field
// 	Duration    func(string, time.Duration) Field
// 	Durationp   func(string, *time.Duration) Field
// 	Durations   func(string, []time.Duration) Field
// 	NamedError  func(string, error) Field
// 	Errors      func(string, []error) Field
// 	Stringer    func(string, fmt.Stringer) Field
// 	Reflect     func(string, interface{}) Field

// 	Skip   func() Field
// 	Inline func(ObjectMarshaler) Field
// }

// var Fields = fields{
// 	Object:      zap.Object,
// 	Array:       zap.Array,
// 	Bool:        zap.Bool,
// 	Boolp:       zap.Boolp,
// 	Bools:       zap.Bools,
// 	Complex128:  zap.Complex128,
// 	Complex128p: zap.Complex128p,
// 	Complex128s: zap.Complex128s,
// 	Complex64:   zap.Complex64,
// 	Complex64p:  zap.Complex64p,
// 	Complex64s:  zap.Complex64s,
// 	Float64:     zap.Float64,
// 	Float64p:    zap.Float64p,
// 	Float64s:    zap.Float64s,
// 	Float32:     zap.Float32,
// 	Float32p:    zap.Float32p,
// 	Float32s:    zap.Float32s,
// 	Int:         zap.Int,
// 	Intp:        zap.Intp,
// 	Ints:        zap.Ints,
// 	Int64:       zap.Int64,
// 	Int64p:      zap.Int64p,
// 	Int64s:      zap.Int64s,
// 	Int32:       zap.Int32,
// 	Int32p:      zap.Int32p,
// 	Int32s:      zap.Int32s,
// 	Int16:       zap.Int16,
// 	Int16p:      zap.Int16p,
// 	Int16s:      zap.Int16s,
// 	Int8:        zap.Int8,
// 	Int8p:       zap.Int8p,
// 	Int8s:       zap.Int8s,
// 	String:      zap.String,
// 	Stringp:     zap.Stringp,
// 	Strings:     zap.Strings,
// 	Uint:        zap.Uint,
// 	Uintp:       zap.Uintp,
// 	Uints:       zap.Uints,
// 	Uint64:      zap.Uint64,
// 	Uint64p:     zap.Uint64p,
// 	Uint64s:     zap.Uint64s,
// 	Uint32:      zap.Uint32,
// 	Uint32p:     zap.Uint32p,
// 	Uint32s:     zap.Uint32s,
// 	Uint16:      zap.Uint16,
// 	Uint16p:     zap.Uint16p,
// 	Uint16s:     zap.Uint16s,
// 	Uint8:       zap.Uint8,
// 	Uint8p:      zap.Uint8p,
// 	Binary:      zap.Binary,
// 	Uintptr:     zap.Uintptr,
// 	Uintptrp:    zap.Uintptrp,
// 	Uintptrs:    zap.Uintptrs,
// 	Time:        zap.Time,
// 	Timep:       zap.Timep,
// 	Times:       zap.Times,
// 	Duration:    zap.Duration,
// 	Durationp:   zap.Durationp,
// 	Durations:   zap.Durations,
// 	NamedError:  zap.NamedError,
// 	Errors:      zap.Errors,
// 	Stringer:    zap.Stringer,
// 	Reflect:     zap.Reflect,

// 	Skip:   zap.Skip,
// 	Inline: zap.Inline,
// }
