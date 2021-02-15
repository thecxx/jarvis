# Jarvis

> It helps us to manage some of scattered jobs of project.

## Samples

### Use `struct` like an object
```
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
		InitObject(obj)
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
