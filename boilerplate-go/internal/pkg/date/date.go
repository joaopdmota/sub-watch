package date

import "time"

type DateProvider interface {
	Now() time.Time
}

type SystemDateProvider struct{}

func NewDateProvider() DateProvider {
	return &SystemDateProvider{}
}

func (d *SystemDateProvider) Now() time.Time {
	return time.Now()
}