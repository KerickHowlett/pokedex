package ttl

import "time"

// TTL (Time-To-Live) presets for cache expiration.
const (
	// OneDayTTL represents the time-to-live (TTL) duration of one day.
	OneDayTTL = 24 * time.Hour
)
