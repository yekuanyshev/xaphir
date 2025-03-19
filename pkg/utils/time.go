package utils

import (
	"time"
)

func WeekStartDate(t time.Time) time.Time {
	offset := (time.Monday - t.Weekday() - 7) % 7
	return toDate(t.Add(time.Duration(offset*24) * time.Hour))
}

func WeekEndDate(t time.Time) time.Time {
	offset := (7 - t.Weekday()) % 7
	return toDate(t.Add(time.Duration(offset*24) * time.Hour))
}

func WeekRange(t time.Time) (start time.Time, end time.Time) {
	start, end = WeekStartDate(t), WeekEndDate(t)
	return
}

func InCurrentWeekRange(t time.Time) bool {
	date := toDate(t)
	weekStart, weekEnd := WeekRange(time.Now())
	return (weekStart.Before(date) || weekStart.Equal(date)) && (weekEnd.After(date) || weekEnd.Equal(date))
}

func InCurrentDay(t time.Time) bool {
	return time.Now().Day() == t.Day()
}

func toDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
