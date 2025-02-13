package provider

import "time"

type TimeProvider interface {
	Now() time.Time
}

type TimeProviderImpl struct{}

func (tp TimeProviderImpl) Now() time.Time {
	return time.Now()
}

type MockTimeProvider struct {
	Time time.Time
}

func (tp *MockTimeProvider) Now() time.Time {
	return tp.Time
}
