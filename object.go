package jarvis

import (
	"runtime"
	"sync"
)

var (
	objOnce sync.Once
)

type Initializer interface {
	Init()
}

type Finalizer interface {
	Delete()
}

// Sample: `InitObject(obj)`
func InitObject(obj interface{}) {
	// Initializer
	if ob, ok := obj.(Initializer); ok {
		ob.Init()
	}
	// Finalizer
	if ob, ok := obj.(Finalizer); ok {
		// Only register the shutdown handler once
		objOnce.Do(func() {
			RegisterShutdownHandler(runtime.GC)
		})
		runtime.SetFinalizer(ob, func(ob Finalizer) {
			ob.Delete()
		})
	}
}
