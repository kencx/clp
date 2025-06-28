package stats

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kencx/clp/entry"
)

var botMap = map[string]string{
	"googlebot":               "Google",
	"bingbot":                 "Bing",
	"duckduckbot":             "DuckDuckGo",
	"duckduckgo-favicons-bot": "DuckDuckGo",
	"slurp":                   "Yahoo",
	"yandexbot":               "Yandex",
	"yandexfavicons":          "Yandex",
	"baiduspider":             "Baidu",
	"exabot":                  "ExaLead",
	"semrushbot":              "SemRush",
	"bytespider":              "",
	"ccbot":                   "",
	"facebookbot":             "Facebook",
	"gptbot":                  "",
	"java":                    "Java",
	"censysinspect":           "",
	"internetmeasurement":     "Internet Measurement",
	"go-http-client":          "Go HTTP Client",
	"okhttp":                  "okhttp",
	"python-requests":         "Python Requests",
	"scrapy":                  "Scrapy",
	"dotbot":                  "DotBot",
	"paloalto":                "Palo Alto",
	"expanse":                 "Palo Alto",
	"petalbot":                "PetalBot",
	"Uptime-Kuma":             "Uptime Kuma",
	"amazonbot":               "Amazon",
}

func FilterByBots(entries entry.Entries) (entry.Entries, error) {
	var filtered entry.Entries

	for _, entry := range entries {
		userAgent, err := getStructField(entry, "UserAgent")
		if err != nil {
			return nil, err
		}

		if userAgent == "" {
			filtered = append(filtered, entry)
			continue
		}

		var match bool
		userAgent = strings.ToLower(userAgent)
		for k := range botMap {
			match = strings.Contains(userAgent, strings.ToLower(k))
			if match {
				break
			}
		}

		if !match {
			filtered = append(filtered, entry)
		}
	}

	return filtered, nil
}

func Filter404(entries entry.Entries) entry.Entries {
	var filtered entry.Entries

	for _, entry := range entries {
		if entry.Status != 404 {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}

func FilterByPeriod(entries entry.Entries, period string) (entry.Entries, error) {
	p, err := parseDuration(period)
	if err != nil {
		return nil, err
	}

	var filtered entry.Entries

	for _, entry := range entries {
		key, err := getStructField(entry, "Timestamp")
		if err != nil {
			return nil, err
		}

		floatKey, err := strconv.ParseFloat(key, 64)
		if err != nil {
			return nil, err
		}

		t := time.Unix(int64(floatKey), 0)
		if !time.Now().After(t.Add(p)) {
			filtered = append(filtered, entry)
		}
	}

	return filtered, nil
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
