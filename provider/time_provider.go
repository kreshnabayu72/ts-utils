package provider

import (
	"time"
)

// TimeProviderer provides a way to pass time.Time as easy access for manipulating time.
type TimeProviderer interface {
	Init()
	Now() time.Time
	Set(t time.Time)
	Close()
}

// TimeProvider implements TimeProviderer.
//
// It is used to get time and safely manipulate time.
type TimeProvider struct {
	Time time.Time
}

// Init initializes time provider. Currently only passes execution.
func (tp *TimeProvider) Init() {
	tp.Time = time.Now()
}

// Now returns the current machine time.
func (tp *TimeProvider) Now() time.Time {
	return time.Now()
}

// Set overrides the currently assigned time for the provider.
func (tp *TimeProvider) Set(t time.Time) {
	tp.Time = t
}

// Close sets Time to zero value.
func (tp *TimeProvider) Close() {
	tp.Time = time.Time{}
}
