package pluglog

import (
	"runtime"
	"strings"
	"sync"
)

type Container struct {
	loggers map[string]*dispatcher
	logMu   sync.RWMutex

	spawner LoggerSpawner
	skip    int

	lvl *levelGetter
}

func GetContainer(spawner LoggerSpawner, skip int) *Container {
	// use defaults if empty
	if spawner == nil {
		spawner = func() Logger {
			return &defaultLogger{}
		}
	}
	return &Container{
		loggers: map[string]*dispatcher{},
		spawner: spawner,
		skip:    skip,
		lvl:     &levelGetter{},
	}
}

func (c *Container) SetLevel(l Level) {
	c.lvl.setMinLevel(l)
}

func (c *Container) GetLogger() Logger {
	pc := make([]uintptr, 10)
	runtime.Callers(2+c.skip, pc)
	f := runtime.FuncForPC(pc[0])
	lindex := strings.LastIndex(f.Name(), ".")
	packName := f.Name()[:lindex]

	c.logMu.Lock()
	defer c.logMu.Unlock()

	logger, ok := c.loggers[packName]
	if !ok {
		logger = &dispatcher{
			logger: c.spawner(),
		}
		c.loggers[packName] = logger
	}

	return &levelLimiter{
		logger: logger,
		lvl:    c.lvl,
	}
}

func (c *Container) GetLoggersList() ([]string, []bool) {
	c.logMu.RLock()
	defer c.logMu.RUnlock()

	ret := make([]string, 0, len(c.loggers))
	enabled := make([]bool, 0, len(c.loggers))

	for key, disp := range c.loggers {
		ret = append(ret, key)
		enabled = append(enabled, disp.active)
	}
	return ret, enabled
}

func (c *Container) ToggleLogger(name string, isActive bool) {
	c.logMu.RLock()
	defer c.logMu.RUnlock()

	logger, ok := c.loggers[name]
	if !ok {
		return
	}

	logger.setActive(isActive)
}
