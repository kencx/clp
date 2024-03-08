package main

func CountUris(entries Entries) (Counter, error) {
	return NewCounter(entries, "Uri")
}

func CountRemoteIPs(entries Entries) (Counter, error) {
	return NewCounter(entries, "RemoteIP")
}

func CountUserAgents(entries Entries) (Counter, error) {
	return NewCounter(entries, "UserAgent")
}

func CountStatusCodes(entries Entries) (Counter, error) {
	return NewCounter(entries, "Status")
}

func UniqueVisitors(entries Entries) (int, error) {
	counter, err := CountRemoteIPs(entries)
	if err != nil {
		return -1, nil
	}
	return len(counter), nil
}

func PageViews(entries Entries) (int, error) {
	counter, err := CountRemoteIPs(entries)
	if err != nil {
		return -1, nil
	}

	var result int
	for _, v := range counter {
		result += v
	}
	return result, nil
}
