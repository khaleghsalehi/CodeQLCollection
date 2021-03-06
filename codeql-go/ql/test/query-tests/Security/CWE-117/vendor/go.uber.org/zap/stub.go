// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for go.uber.org/zap, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: go.uber.org/zap (exports: Logger,SugaredLogger; functions: NewProduction)

// Package zap is a stub of go.uber.org/zap, generated by depstubber.
package zap

type Logger struct{}

func (_ *Logger) Check(_ interface{}, _ string) interface{} {
	return nil
}

func (_ *Logger) Core() interface{} {
	return nil
}

func (_ *Logger) DPanic(_ string, _ ...interface{}) {}

func (_ *Logger) Debug(_ string, _ ...interface{}) {}

func (_ *Logger) Error(_ string, _ ...interface{}) {}

func (_ *Logger) Fatal(_ string, _ ...interface{}) {}

func (_ *Logger) Info(_ string, _ ...interface{}) {}

func (_ *Logger) Named(_ string) *Logger {
	return nil
}

func (_ *Logger) Panic(_ string, _ ...interface{}) {}

func (_ *Logger) Sugar() *SugaredLogger {
	return nil
}

func (_ *Logger) Sync() error {
	return nil
}

func (_ *Logger) Warn(_ string, _ ...interface{}) {}

func (_ *Logger) With(_ ...interface{}) *Logger {
	return nil
}

func (_ *Logger) WithOptions(_ ...Option) *Logger {
	return nil
}

func NewProduction(_ ...Option) (*Logger, error) {
	return nil, nil
}

type Option interface{}

type SugaredLogger struct{}

func (_ *SugaredLogger) DPanic(_ ...interface{}) {}

func (_ *SugaredLogger) DPanicf(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) DPanicw(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Debug(_ ...interface{}) {}

func (_ *SugaredLogger) Debugf(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Debugw(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Desugar() *Logger {
	return nil
}

func (_ *SugaredLogger) Error(_ ...interface{}) {}

func (_ *SugaredLogger) Errorf(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Errorw(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Fatal(_ ...interface{}) {}

func (_ *SugaredLogger) Fatalf(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Fatalw(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Info(_ ...interface{}) {}

func (_ *SugaredLogger) Infof(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Infow(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Named(_ string) *SugaredLogger {
	return nil
}

func (_ *SugaredLogger) Panic(_ ...interface{}) {}

func (_ *SugaredLogger) Panicf(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Panicw(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Sync() error {
	return nil
}

func (_ *SugaredLogger) Warn(_ ...interface{}) {}

func (_ *SugaredLogger) Warnf(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) Warnw(_ string, _ ...interface{}) {}

func (_ *SugaredLogger) With(_ ...interface{}) *SugaredLogger {
	return nil
}
