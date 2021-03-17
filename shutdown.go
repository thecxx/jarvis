package jarvis

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	shut    shutdown
	signals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	}
)

type shutdown struct {
	once     sync.Once
	exited   bool
	locker   sync.Mutex
	handlers []func()
	receiver chan os.Signal
}

// init initializes shutdown.
func (s *shutdown) init() {
	s.handlers = make([]func(), 0)
	s.receiver = make(chan os.Signal)
	// Catch some signals
	signal.Notify(s.receiver, signals...)
	// Wait exit signal and shutdown
	go func() {
		<-shut.receiver
		// Shutdown
		s.shutdown()
	}()
}

// register sets a new shutdown handler.
func (s *shutdown) register(f func()) error {
	if f == nil {
		return errors.New("invalid handler")
	}
	s.locker.Lock()
	defer s.locker.Unlock()
	if s.exited {
		return errors.New("already exited")
	}
	// Register
	s.handlers = append(s.handlers, f)

	return nil
}

func (s *shutdown) shutdown() {
	s.locker.Lock()
	defer s.locker.Unlock()
	// Set exited = true
	if s.exited {
		return
	} else {
		s.exited = true
	}
	// Call one by one
	for _, fun := range shut.handlers {
		fun()
	}
	// Clean
	shut.handlers = nil
}

// Shutdown starts the shutdown logic,
// but it will never call `os.Exit()`.
func Shutdown() {
	shut.shutdown()
}

// RegisterShutdownHandler sets a new shutdown handler.
func RegisterShutdownHandler(f func()) error {
	// Initialize once
	shut.once.Do(func() {
		shut.init()
	})
	// Register shutdown handler
	return shut.register(f)
}
