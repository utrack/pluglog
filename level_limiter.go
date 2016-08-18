package pluglog

import "sync"

type Level uint

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

type levelGetter struct {
	minLevel Level
	mu       sync.RWMutex
}

func (g *levelGetter) getMinLevel() Level {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.minLevel
}

func (g *levelGetter) setMinLevel(l Level) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.minLevel = l
}

type levelLimiter struct {
	lvl    *levelGetter
	logger Logger
}

func (l *levelLimiter) Trace(s string) {
	if l.lvl.getMinLevel() > LevelTrace {
		return
	}
	l.logger.Trace(s)
}
func (l *levelLimiter) Tracef(s string, opts ...interface{}) {
	if l.lvl.getMinLevel() > LevelTrace {
		return
	}
	l.logger.Tracef(s, opts...)
}

func (l *levelLimiter) Debug(s string) {
	if l.lvl.getMinLevel() > LevelDebug {
		return
	}
	l.logger.Debug(s)
}
func (l *levelLimiter) Debugf(s string, opts ...interface{}) {
	if l.lvl.getMinLevel() > LevelDebug {
		return
	}
	l.logger.Debugf(s, opts...)
}

func (l *levelLimiter) Info(s string) {
	if l.lvl.getMinLevel() > LevelInfo {
		return
	}
	l.logger.Info(s)
}
func (l *levelLimiter) Infof(s string, opts ...interface{}) {
	if l.lvl.getMinLevel() > LevelInfo {
		return
	}
	l.logger.Infof(s, opts...)
}

func (l *levelLimiter) Warn(s string) {
	if l.lvl.getMinLevel() > LevelWarn {
		return
	}
	l.logger.Warn(s)
}
func (l *levelLimiter) Warnf(s string, opts ...interface{}) {
	if l.lvl.getMinLevel() > LevelWarn {
		return
	}
	l.logger.Warnf(s, opts...)
}

func (l *levelLimiter) Error(s string) {
	if l.lvl.getMinLevel() > LevelError {
		return
	}
	l.logger.Error(s)
}
func (l *levelLimiter) Errorf(s string, opts ...interface{}) {
	if l.lvl.getMinLevel() > LevelError {
		return
	}
	l.logger.Errorf(s, opts...)
}
