package logs

import (
	"fmt"
	"sync"
)

type Logs struct {
	level LogLevel
	mu    sync.RWMutex
}

var (
	instance *Logs
	once     sync.Once
)

func GetLog() *Logs {
	once.Do(func() {
		instance = &Logs{
			level: DEBUG,
		}
	})

	return instance
}

func (l *Logs) SetLevel(lvl LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = lvl
}

func (l *Logs) Debug(msg string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.level <= DEBUG {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}

func (l *Logs) Info(msg string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.level <= INFO {
		fmt.Printf("[INFO] %s\n", msg)
	}
}

func (l *Logs) Warn(msg string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.level <= WARN {
		fmt.Printf("[WARN] %s\n", msg)
	}
}

func (l *Logs) Error(msg string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.level <= ERROR {
		fmt.Printf("[ERROR] %s\n", msg)
	}
}
