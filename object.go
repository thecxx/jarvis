package jarvis

import (
	"runtime"
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
		runtime.SetFinalizer(ob, func(ob Finalizer) {
			ob.Delete()
		})
	}
}
