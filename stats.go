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
