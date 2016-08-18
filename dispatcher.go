package pluglog

import "sync"

type dispatcher struct {
	active bool
	logger Logger
	mu     sync.RWMutex
}

func (d *dispatcher) setActive(a bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.active = a
}

func (d *dispatcher) Trace(s string) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Trace(s)
}
func (d *dispatcher) Tracef(s string, opts ...interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Tracef(s, opts...)
}

func (d *dispatcher) Debug(s string) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Debug(s)
}
func (d *dispatcher) Debugf(s string, opts ...interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Debugf(s, opts...)
}

func (d *dispatcher) Info(s string) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Info(s)
}
func (d *dispatcher) Infof(s string, opts ...interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Infof(s, opts...)
}

func (d *dispatcher) Warn(s string) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Warn(s)
}
func (d *dispatcher) Warnf(s string, opts ...interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Warnf(s, opts...)
}

func (d *dispatcher) Error(s string) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Error(s)
}
func (d *dispatcher) Errorf(s string, opts ...interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if !d.active {
		return
	}
	d.logger.Errorf(s, opts...)
}
