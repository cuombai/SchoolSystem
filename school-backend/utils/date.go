package utils

import "time"

func FormatDate(t time.Time) string {
    return t.Format("2006-01-02")
}

func FormatDateVerbose(t time.Time) string {
    return t.Format("Monday, 02 January 2006")
}
