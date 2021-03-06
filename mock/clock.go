package mock

import (
	"sync"
	"time"

	"github.com/Patagonicus/log"
)

// Clock is an implementation of log.Clock used for mocking.
// It starts with a given time and increases that time by a fixed amount every time Now() is called.
// Clock is safe to be used concurrently.
type Clock struct {
	current time.Time
	step    time.Duration
	lock    sync.Locker
}

// Clock has to implement log.Clock
var _ log.Clock = &Clock{}

// NewClock returns a new Clock. This Clock will return start the first time Now() is called.
// Afterwards it will return start+step, start+2*step, …
func NewClock(start time.Time, step time.Duration) *Clock {
	return &Clock{start, step, new(sync.Mutex)}
}

func (c *Clock) Now() time.Time {
	c.lock.Lock()
	defer c.lock.Unlock()

	now := c.current
	c.current = c.current.Add(c.step)

	return now
}

// Next returns the time the next call to Now will return. The difference to Now is that it will not increase the time.
func (c *Clock) Next() time.Time {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.current
}
