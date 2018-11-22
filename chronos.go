package chronos

import (
	"time"
)

// Chronos chronos struct
type Chronos struct {
	time.Time
}

// New initialize Chronos with time
func New(t time.Time) *Chronos {
	return &Chronos{t}
}
