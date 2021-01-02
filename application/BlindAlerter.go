package poker

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter temporary blind alert
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc type for implement the blind scheduler fn
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt calls the blind alerter
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter prints the blind
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
