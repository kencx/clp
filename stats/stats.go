package stats

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/kencx/clp/entry"
)

func CountUris(entries entry.Entries) (Counter, error) {
	return NewCounter(entries, "Uri")
}

func CountRemoteIPs(entries entry.Entries) (Counter, error) {
	return NewCounter(entries, "RemoteIP")
}

func CountUserAgents(entries entry.Entries) (Counter, error) {
	return NewCounter(entries, "UserAgent")
}

func CountStatusCodes(entries entry.Entries) (Counter, error) {
	return NewCounter(entries, "Status")
}

func UniqueVisitors(entries entry.Entries) (int, error) {
	counter, err := CountRemoteIPs(entries)
	if err != nil {
		return -1, nil
	}
	return len(counter), nil
}

func PageViews(entries entry.Entries) int {
	return len(entries)
}

func AverageDuration(entries entry.Entries) float64 {
	var res float64

	for _, entry := range entries {
		res += entry.Duration
	}
	return res / float64(len(entries))
}

func Summary(entries entry.Entries, duration string, crawlers bool) error {
	_, err := parseDuration(duration)
	if err != nil {
		return err
	}

	uv, err := UniqueVisitors(entries)
	if err != nil {
		return err
	}

	fmt.Printf("Unique visitors: %d\n", uv)
	fmt.Printf("Total page views: %d\n", PageViews(entries))
	fmt.Printf("Average response time: %.3f ms\n", AverageDuration(entries)*1000)

	return nil
}

func parseDuration(duration string) (time.Duration, error) {
	dMap := map[string]time.Duration{
		"h": time.Hour,
		"d": 24 * time.Hour,
		"w": 7 * 24 * time.Hour,
		"m": 30 * 24 * time.Hour,
		"y": 365 * 24 * time.Hour,
	}

	rx := regexp.MustCompile(`(\d+)(\w)$`)
	matches := rx.FindAllStringSubmatch(duration, -1)

	var number, unitString string
	for _, match := range matches {
		number = match[1]
		unitString = match[2]

	}

	multiple, err := strconv.Atoi(number)
	if err != nil {
		return -1, err
	}

	unit, ok := dMap[unitString]
	if !ok {
		return -1, fmt.Errorf("invalid unit: %v", unitString)
	}

	return time.Duration(multiple) * unit, nil
}
