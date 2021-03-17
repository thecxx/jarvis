# Jarvis

> It helps us to manage some of scattered jobs of project.

## Functions
- jarvis.InitObject(obj interface{})
- jarvis.Shutdown()
- jarvis.RegisterShutdownHandler(f func()) error

## Samples

### Use `RegisterShutdownHandler`
```
package main

import (
	"fmt"
	"time"

	"github.com/thecxx/jarvis"
)

func main() {
	// Exit without exception,
	// when the process exits abnormally, 
	// it will catch the exit signal and automatically call `shutdown`.
	defer jarvis.Shutdown()

	jarvis.RegisterShutdownHandler(func() {
		fmt.Printf("exited 1\n")
	})
	jarvis.RegisterShutdownHandler(func() {
		fmt.Printf("exited 2\n")
	})
	jarvis.RegisterShutdownHandler(func() {
		fmt.Printf("exited 3\n")
	})

}
```

### Use `struct` like an object
```
package main

import (
	"fmt"

	"github.com/thecxx/jarvis"
)

type sample struct {
}

type Sample struct {
	*sample
}

// NewSample returns a Sample.
func NewSample() *Sample {
	// Private object
	s := &sample{}
	// Public object
	obj := &Sample{s}
	{
		jarvis.InitObject(obj)
	}
	return obj
}

func (s *sample) Init() {
	fmt.Printf("initialized\n")
}

func (s *sample) Delete() {
	fmt.Printf("deleted\n")
}
```
