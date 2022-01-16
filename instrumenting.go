package instrumenting

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Tracer interface {
	Init(scope string, withFields ...zap.Field) Tracer
	Child(segment string, withFields ...zap.Field) Tracer

	GetStartTime() time.Time
	GetTimeStamp() time.Time
	GetElapsed() time.Duration

	Trace(location string, data ...zap.Field)
	TraceWithError(location string, err error, data ...zap.Field)
	Finish(data ...zap.Field)
}

type _DummyTracer struct {
	StartedTStamp time.Time
	TStamp        time.Time
	Elapsed       time.Duration
}

func (t *_DummyTracer) Init(scope string, withFields ...zap.Field) Tracer {
	t.StartedTStamp = time.Now()
	t.TStamp = t.StartedTStamp
	return t
}
func (t *_DummyTracer) Child(segment string, data ...zap.Field) Tracer {
	tt := (&_DummyTracer{})
	return tt.Init(segment)
}

func (t *_DummyTracer) GetStartTime() time.Time {
	return t.StartedTStamp
}
func (t *_DummyTracer) GetTimeStamp() time.Time {
	return t.TStamp
}
func (t *_DummyTracer) GetElapsed() time.Duration {
	return t.Elapsed
}
func (t *_DummyTracer) Trace(location string, data ...zap.Field) {
	now := time.Now()
	t.Elapsed = now.Sub(t.TStamp)
	t.TStamp = time.Now()
}

func (t *_DummyTracer) TraceWithError(location string, err error, data ...zap.Field) {
	t.Trace(location)
}
func (t *_DummyTracer) Finish(data ...zap.Field) {
	t.Trace("finished")
}

const (
	SCOPE_KEY   = "scope"
	TIME_KEY    = "time"
	ELAPSED_KEY = "elapsed"
)

var (
	DefaultLogger *zap.Logger
)

func MicrosecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Microsecond))
}
func init() {
	var err error
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}

	config.EncoderConfig.MessageKey = "loc"
	config.EncoderConfig.CallerKey = "src"

	config.EncoderConfig.LevelKey = zapcore.OmitKey
	config.EncoderConfig.TimeKey = zapcore.OmitKey

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	config.EncoderConfig.EncodeDuration = MicrosecondsDurationEncoder

	DefaultLogger, err = config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
	if err != nil {
		panic("Failed to initialize Zap logger: " + err.Error())
	}

}

type DefaultTracer struct {
	StartedTStamp time.Time
	TStamp        time.Time
	Elapsed       time.Duration
	Scope         string
	WithFields    []zap.Field
	Logger        zap.Logger
}

func (t *DefaultTracer) Init(scope string, withFields ...zap.Field) Tracer {
	t.Scope = scope
	t.WithFields = withFields
	if len(scope) > 0 {
		t.Logger = *DefaultLogger.With(zap.String(SCOPE_KEY, scope))
	} else {
		t.Logger = *DefaultLogger
	}
	if len(withFields) > 0 {
		t.Logger = *t.Logger.With(withFields...)
	}
	t.Logger.Info("started", zap.Time("time", time.Now()))
	t.TStamp = time.Now()
	t.StartedTStamp = t.TStamp
	return t
}

func (t *DefaultTracer) GetStartTime() time.Time {
	return t.StartedTStamp
}
func (t *DefaultTracer) GetElapsed() time.Duration {
	return t.Elapsed
}

func (t *DefaultTracer) GetTimeStamp() time.Time {
	return t.TStamp
}

func (t *DefaultTracer) traceInternal(location string, err error, showErr bool, data ...zap.Field) {
	now := time.Now()
	t.Elapsed = now.Sub(t.TStamp)
	logger := t.Logger.
		With(zap.Duration(ELAPSED_KEY, time.Duration(t.Elapsed))).
		With(data...)
	if showErr {
		logger = logger.With(zap.Error(err))
	}
	logger.Info(location)
	t.TStamp = time.Now()
}

func (t *DefaultTracer) Trace(location string, data ...zap.Field) {
	t.traceInternal(location, nil, false, data...)
}
func (t *DefaultTracer) TraceWithError(location string, err error, data ...zap.Field) {
	t.traceInternal(location, err, true, data...)
}

func (t *DefaultTracer) Finish(data ...zap.Field) {
	t.TStamp = t.StartedTStamp
	t.traceInternal("finished", nil, false, data...)
}

func (t *DefaultTracer) Child(segment string, withFields ...zap.Field) Tracer {
	var scope string
	if len(t.Scope) > 0 {
		scope = t.Scope + ":" + segment
	} else {
		scope = segment
	}
	res := &DefaultTracer{}
	if len(t.WithFields) > 0 {
		withFields = append(t.WithFields, withFields...)
	}
	return res.Init(scope, withFields...)
}

func NewTracer(scope string, active bool, withFields ...zap.Field) Tracer {
	if !active {
		return (&_DummyTracer{}).Init(scope)
	}
	tracer := &DefaultTracer{}
	tracer.Init(scope, withFields...)
	return tracer
}
