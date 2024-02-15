package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// TestGet validates logger get
func TestGet(t *testing.T) {
	tests := []struct {
		name string
		log  *zap.Logger
		fn   func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "when is nil",
			log:  nil,
			fn:   assert.Nil,
		},
		{
			name: "when is not nil",
			log:  &zap.Logger{},
			fn:   assert.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log = tt.log
			tt.fn(t, Get())
		})
	}
}

// TestInfo handle log with info level
func TestInfo(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	log = zap.New(observedZapCore)

	type args struct {
		message string
		fields  []zap.Field
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Without any field",
			args: args{
				message: "whithout fields",
			},
		},
		{
			name: "With some fields",
			args: args{
				message: "whith fields",
				fields: []zap.Field{
					zap.String("position", "4411cacb-6df6-42d1-b4dc-e804147b6e7f"),
					zap.Bool("test", true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.message, tt.args.fields...)
			logs := observedLogs.FilterMessage(tt.args.message).All()
			assert.ElementsMatch(t, tt.args.fields, logs[0].Context)
		})
	}
}

// TestDebug handle log with Debug level
func TestDebug(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.DebugLevel)
	log = zap.New(observedZapCore)

	type args struct {
		message string
		fields  []zap.Field
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Without any field",
			args: args{
				message: "whithout fields",
			},
		},
		{
			name: "With some fields",
			args: args{
				message: "whith fields",
				fields: []zap.Field{
					zap.String("position", "4411cacb-6df6-42d1-b4dc-e804147b6e7f"),
					zap.Bool("test", true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.message, tt.args.fields...)
			logs := observedLogs.FilterMessage(tt.args.message).All()
			assert.ElementsMatch(t, tt.args.fields, logs[0].Context)
		})
	}
}

// TestError handle log with Error level
func TestError(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.ErrorLevel)
	log = zap.New(observedZapCore)

	type args struct {
		message string
		fields  []zap.Field
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Without any field",
			args: args{
				message: "whithout fields",
			},
		},
		{
			name: "With some fields",
			args: args{
				message: "whith fields",
				fields: []zap.Field{
					zap.String("position", "4411cacb-6df6-42d1-b4dc-e804147b6e7f"),
					zap.Bool("test", true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.message, tt.args.fields...)
			logs := observedLogs.FilterMessage(tt.args.message).All()
			assert.ElementsMatch(t, tt.args.fields, logs[0].Context)
		})
	}
}

// TestWarn handle log with Warn level
func TestWarn(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	log = zap.New(observedZapCore)

	type args struct {
		message string
		fields  []zap.Field
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Without any field",
			args: args{
				message: "whithout fields",
			},
		},
		{
			name: "With some fields",
			args: args{
				message: "whith fields",
				fields: []zap.Field{
					zap.String("position", "4411cacb-6df6-42d1-b4dc-e804147b6e7f"),
					zap.Bool("test", true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.message, tt.args.fields...)
			logs := observedLogs.FilterMessage(tt.args.message).All()
			assert.ElementsMatch(t, tt.args.fields, logs[0].Context)
		})
	}
}

// TestPanic handle log with Panic level
func TestPanic(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.PanicLevel)
	log = zap.New(observedZapCore)

	type args struct {
		message string
		fields  []zap.Field
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Without any field",
			args: args{
				message: "whithout fields",
			},
		},
		{
			name: "With some fields",
			args: args{
				message: "whith fields",
				fields: []zap.Field{
					zap.String("position", "4411cacb-6df6-42d1-b4dc-e804147b6e7f"),
					zap.Bool("test", true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				Panic(tt.args.message, tt.args.fields...)
			}, "Expected panic")

			logs := observedLogs.FilterMessage(tt.args.message).All()
			assert.ElementsMatch(t, tt.args.fields, logs[0].Context)
		})
	}
}

// TestFatal handle log with Fatal level
func TestFatal(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.FatalLevel)
	log = zap.New(observedZapCore)

	// this will force to panic instead of call exit(1)
	log = log.WithOptions(zap.OnFatal(zapcore.WriteThenPanic))

	type args struct {
		message string
		fields  []zap.Field
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Without any field",
			args: args{
				message: "whithout fields",
			},
		},
		{
			name: "With some fields",
			args: args{
				message: "whith fields",
				fields: []zap.Field{
					zap.String("position", "4411cacb-6df6-42d1-b4dc-e804147b6e7f"),
					zap.Bool("test", true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Panics(t, func() {
				Fatal(tt.args.message, tt.args.fields...)
			}, "Expected panic instead of fatal with exit(1)")

			logs := observedLogs.FilterMessage(tt.args.message).All()
			assert.ElementsMatch(t, tt.args.fields, logs[0].Context)
		})
	}
}
