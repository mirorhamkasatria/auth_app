package utils

import (
	"testing"
	"time"
)

func TestGetTimeNow(t *testing.T) {
	tests := []struct {
		name            string
		acceptableDelta time.Duration
	}{
		// Test case 1: Check if the time obtained is within a reasonable range
		{
			name:            "Current Time",
			acceptableDelta: 2 * time.Second, // You can adjust the acceptable delta based on your needs
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			got := GetTimeNow()
			endTime := time.Now()

			if got.Before(startTime) || got.After(endTime) {
				t.Errorf("GetTimeNow() result %v is not within the expected range [%v, %v]", got, startTime, endTime)
			}

			elapsedTime := endTime.Sub(startTime)
			if elapsedTime > tt.acceptableDelta {
				t.Errorf("GetTimeNow() took longer than the acceptable delta. Elapsed time: %v, Acceptable delta: %v", elapsedTime, tt.acceptableDelta)
			}
		})
	}
}
