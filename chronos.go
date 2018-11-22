package chronos

import (
	"time"
)

// Chronos struct
type Chronos struct {
	time.Time
}

// Now get now
func Now(c *Chronos) time.Time {
	return time.Now()
}
