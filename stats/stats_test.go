package stats

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration string
		want     time.Duration
		err      bool
	}{{
		name:     "3 months",
		duration: "3m",
		want:     3 * 30 * 24 * time.Hour,
		err:      false,
	}, {
		name:     "20 days",
		duration: "20d",
		want:     20 * 24 * time.Hour,
		err:      false,
	}, {
		name:     "100 weeks",
		duration: "100w",
		want:     100 * 7 * 24 * time.Hour,
		err:      false,
	}, {
		name:     "4 years",
		duration: "4y",
		want:     4 * 365 * 24 * time.Hour,
		err:      false,
	}, {
		name:     "invalid unit",
		duration: "10s",
		want:     -1,
		err:      true,
	}, {
		name:     "invalid number",
		duration: "fooh",
		want:     -1,
		err:      true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDuration(tt.duration)
			if err != nil && !tt.err {
				t.Fatalf("expected no err, got %v", err)
			}

			if err == nil && tt.err {
				t.Fatalf("expected err, got nil")
			}

			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
