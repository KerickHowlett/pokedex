package query_cache_config

import (
	"testing"
	"time"
)

func TestWithNow(t *testing.T) {
	setup := func() func() time.Time {
		config := &QueryCacheConfig{}
		WithNow(func() time.Time { return time.Now() })(config)
		return config.Now
	}

	t.Run("should set the Now field of the QueryCacheConfig struct", func(t *testing.T) {
		// @NOTE: There's no way to test if it's the same function given how its implemented.
		if actual := setup(); actual == nil {
			t.Error("Expected the Now field to be set, but it was nil.")
		}
	})
}

func TestWithTTL(t *testing.T) {
	setup := func() (actual time.Duration, expected time.Duration) {
		expected = 5 * time.Minute
		config := &QueryCacheConfig{}
		WithTTL(expected)(config)
		actual = config.TTL
		return actual, expected
	}

	t.Run("should set the TTL field of the QueryCacheConfig struct", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected: %v,\nActual: %v", expected, actual)
		}
	})
}
