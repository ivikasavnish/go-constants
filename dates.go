package constants

import "time"

const (
	Duration = time.Duration(1)
	Second   = Duration * time.Second
	Minute   = 60 * Second
	Hour     = 60 * Minute
	Day      = 24 * Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 365 * Day
	Ms       = Duration * time.Millisecond
	Us       = Duration * time.Microsecond
	Ns       = Duration * time.Nanosecond
)
