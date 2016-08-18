package pluglog

import "fmt"

type LoggerSpawner func() Logger

type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Error(string)
	Errorf(string, ...interface{})
}

type defaultLogger struct{}

func (l *defaultLogger) Trace(s string) {
	fmt.Printf("TRACE | %v\n", s)
}
func (l *defaultLogger) Tracef(s string, opts ...interface{}) {
	fmt.Printf("TRACE | %v\n", fmt.Sprintf(s, opts))
}

func (l *defaultLogger) Debug(s string) {
	fmt.Printf("DEBUG | %v\n", s)
}
func (l *defaultLogger) Debugf(s string, opts ...interface{}) {
	fmt.Printf("DEBUG | %v\n", fmt.Sprintf(s, opts))
}

func (l *defaultLogger) Info(s string) {
	fmt.Printf("INFO | %v\n", s)
}
func (l *defaultLogger) Infof(s string, opts ...interface{}) {
	fmt.Printf("INFO | %v\n", fmt.Sprintf(s, opts))
}

func (l *defaultLogger) Warn(s string) {
	fmt.Printf("WARN | %v\n", s)
}
func (l *defaultLogger) Warnf(s string, opts ...interface{}) {
	fmt.Printf("WARN | %v\n", fmt.Sprintf(s, opts))
}

func (l *defaultLogger) Error(s string) {
	fmt.Printf("ERROR | %v\n", s)
}
func (l *defaultLogger) Errorf(s string, opts ...interface{}) {
	fmt.Printf("ERROR | %v\n", fmt.Sprintf(s, opts))
}
