package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeekStartDate(t *testing.T) {
	testCases := []struct {
		current  time.Time
		expected time.Time
	}{
		{current: date(2025, time.March, 19), expected: date(2025, time.March, 17)},
		{current: date(2025, time.March, 20), expected: date(2025, time.March, 17)},
		{current: date(2025, time.March, 23), expected: date(2025, time.March, 17)},
		{current: date(2025, time.March, 17), expected: date(2025, time.March, 17)},
		{current: date(2025, time.March, 24), expected: date(2025, time.March, 24)},
		{current: date(2025, time.March, 30), expected: date(2025, time.March, 24)},
		{current: date(2025, time.March, 1), expected: date(2025, time.February, 24)},
		{current: date(2025, time.April, 6), expected: date(2025, time.March, 31)},
	}

	for _, tc := range testCases {
		got := WeekStartDate(tc.current)
		assert.Equal(t, tc.expected, got)
	}
}

func TestWeekEndDate(t *testing.T) {
	testCases := []struct {
		current  time.Time
		expected time.Time
	}{
		{current: date(2025, time.March, 19), expected: date(2025, time.March, 23)},
		{current: date(2025, time.March, 20), expected: date(2025, time.March, 23)},
		{current: date(2025, time.March, 23), expected: date(2025, time.March, 23)},
		{current: date(2025, time.March, 17), expected: date(2025, time.March, 23)},
		{current: date(2025, time.March, 24), expected: date(2025, time.March, 30)},
		{current: date(2025, time.March, 30), expected: date(2025, time.March, 30)},
		{current: date(2025, time.March, 1), expected: date(2025, time.March, 2)},
		{current: date(2025, time.March, 31), expected: date(2025, time.April, 6)},
	}

	for _, tc := range testCases {
		got := WeekEndDate(tc.current)
		assert.Equal(t, tc.expected, got)
	}
}

func TestWeekRange(t *testing.T) {
	testCases := []struct {
		current       time.Time
		expectedStart time.Time
		expectedEnd   time.Time
	}{
		{current: date(2025, time.March, 19), expectedStart: date(2025, time.March, 17), expectedEnd: date(2025, time.March, 23)},
		{current: date(2025, time.March, 23), expectedStart: date(2025, time.March, 17), expectedEnd: date(2025, time.March, 23)},
		{current: date(2025, time.April, 1), expectedStart: date(2025, time.March, 31), expectedEnd: date(2025, time.April, 6)},
		{current: date(2025, time.March, 1), expectedStart: date(2025, time.February, 24), expectedEnd: date(2025, time.March, 2)},
	}

	for _, tc := range testCases {
		actualStart, actualEnd := WeekRange(tc.current)
		assert.Equal(t, tc.expectedStart, actualStart)
		assert.Equal(t, tc.expectedEnd, actualEnd)
	}
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
