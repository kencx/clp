package stats

import (
	"fmt"

	"github.com/kencx/clp/entry"
)

func Summary(entries entry.Entries, period string, crawlers, color bool) error {
	filtered, err := FilterByPeriod(entries, period)
	if err != nil {
		return err
	}

	if crawlers {
		filtered, err = FilterByBots(filtered)
		if err != nil {
			return err
		}
	}

	uv, err := UniqueVisitors(filtered)
	if err != nil {
		return err
	}

	uris, err := CountUris(filtered)
	if err != nil {
		return err
	}
	remoteIp, err := CountRemoteIPs(filtered)
	if err != nil {
		return err
	}
	status, err := CountStatusCodes(filtered)
	if err != nil {
		return err
	}
	userAgents, err := CountUserAgents(filtered)
	if err != nil {
		return err
	}
	avg, large, small := DurationStats(filtered)

	fmt.Println("Caddy Access Logs Summary Statistics")
	fmt.Printf("Period: %s\n", period)
	fmt.Println("")

	fmt.Printf("Unique visitors: %d\n", uv)
	fmt.Printf("Total page views: %d\n", PageViews(filtered))
	fmt.Printf("Average response time: %.3f ms\n", avg*1000)

	fmt.Println("")
	fmt.Println("URIs")
	uris.PrintTopN(5)

	fmt.Println("")
	fmt.Println("Remote IPs")
	remoteIp.PrintTopN(5)

	fmt.Println("")
	fmt.Println("Status Codes")
	status.PrintTopN(5)

	fmt.Println("")
	fmt.Println("User Agents")
	userAgents.PrintTopN(5)

	fmt.Println("")
	fmt.Printf("Response time distribution:\n")
	fmt.Printf("Max: %.3f ms\n", large*1000)
	fmt.Printf("Min: %.3f ms\n", small*1000)

	return nil
}

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

func DurationStats(entries entry.Entries) (float64, float64, float64) {
	var sum float64
	var largest float64
	var smallest float64

	for _, entry := range entries {
		sum += entry.Duration
		largest = max(entry.Duration, largest)

		if smallest > 0 {
			smallest = min(entry.Duration, smallest)
		} else {
			smallest = entry.Duration
		}

	}
	return sum / float64(len(entries)), largest, smallest
}
