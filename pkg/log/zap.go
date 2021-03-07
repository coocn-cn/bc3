package log

import (
	"context"

	"go.uber.org/zap"
)

type zapLogger struct {
	_log       zap.Logger
	ctx        context.Context
	fields     []zap.Field
	callerSkip int
}

// NewZapLogger a
func NewZapLogger() FieldLogger {
	raw, _ := zap.NewDevelopment()
	return &zapLogger{
		_log: raw,
	}
}

func (m *zapLogger) log() zap.Logger {
	return m._log
}

func (m *zapLogger) clone() zap.Logger {
	fields := make([]zap.Field, 0, len(m.fields))

	for _, v := range m.fields {
		fields = append(fields, v)
	}

	return &zapLogger{
		_log:       m._log,
		ctx:        m.ctx,
		fields:     fields,
		callerSkip: m.callerSkip,
	}

}

func (m *zapLogger) WithError(err error) FieldLogger {
	return m.log()
}

func (m *zapLogger) WithContext(ctx context.Context) FieldLogger {
	return m
}

func (m *zapLogger) WithField(key string, value interface{}) FieldLogger {
	return m
}

func (m *zapLogger) WithFields(fields Fields) FieldLogger {
	return m
}

func (m *zapLogger) Debug(args ...interface{}) {
}

func (m *zapLogger) Debugf(format string, args ...interface{}) {
}

func (m *zapLogger) Info(args ...interface{}) {
}

func (m *zapLogger) Infof(format string, args ...interface{}) {
}

func (m *zapLogger) Warn(args ...interface{}) {
}

func (m *zapLogger) Warnf(format string, args ...interface{}) {
}

func (m *zapLogger) Error(args ...interface{}) {
}

func (m *zapLogger) Errorf(format string, args ...interface{}) {
}
