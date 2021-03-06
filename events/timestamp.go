package events

import "time"

const (
	// TimestampLayout is used for package-wide timestamping of events
	TimestampLayout = "2006-01-02T15:04:05.999999-07:00"
)

// FormatedTimestamp creates a string formated timestamp of the current time and date(time.Now())
func FormatedTimestamp() string {
	return time.Now().Format(TimestampLayout)
}

// FormatTimestamp allows to format any given time to a string
func FormatTimestamp(t time.Time) string {
	return t.Format(TimestampLayout)
}

// ParseTimestamp allows to parse the timestamp according to the package-wide
// timestamp layout that is saved in the TimestampLayout variable
func ParseTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(TimestampLayout, timestamp)
}
