package ttl

import "time"

// TTL (Time-To-Live) presets for cache expiration.
const (
	// Disable represents the time-to-live (TTL) duration of zero seconds.
	Disable = 0 * time.Second
	// OneDay represents the time-to-live (TTL) duration of one day.
	OneDay = 24 * time.Hour
)
