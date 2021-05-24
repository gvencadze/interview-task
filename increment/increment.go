package increment

import (
	"sync"
)

// Incrementor instance
type Incrementor struct {
	i     uint
	max   uint
	mutex sync.RWMutex
}

// New -- create new instance of incrementor
func New(value uint) *Incrementor {
	return &Incrementor{
		max: value,
	}
}

// GetNumber -- get current number value
func (inc *Incrementor) GetNumber() uint {
	inc.mutex.RLock()
	defer inc.mutex.RUnlock()

	return inc.i
}

// Increment -- increment current number value
func (inc *Incrementor) Increment() {
	inc.mutex.Lock()
	defer inc.mutex.Unlock()

	if inc.i+1 >= inc.max {
		inc.i = 0
	} else {
		inc.i++
	}
}

// SetMaximumValue -- set maximum possible number value
func (inc *Incrementor) SetMaximumValue(value uint) {
	inc.mutex.Lock()
	defer inc.mutex.Unlock()

	if inc.i >= value {
		inc.i = 0
	}

	inc.max = value
}
