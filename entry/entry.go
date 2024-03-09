package entry

type Entry struct {
	Level           string   `json:"level"`
	Timestamp       float64  `json:"ts"`
	Logger          string   `json:"logger"`
	Msg             string   `json:"msg"`
	Request         Request  `json:"request"`
	BytesRead       int      `json:"bytes_read"`
	UserId          string   `json:"user_id"`
	Duration        float64  `json:"duration"`
	Size            int      `json:"size"`
	Status          int      `json:"status"`
	ResponseHeaders Response `json:"resp_headers"`
}

type Request struct {
	RemoteIP   string `json:"remote_ip"`
	RemotePort string `json:"remote_port"`
	ClientIP   string `json:"client_ip"`
	Proto      string `json:"proto"`
	Method     string `json:"method"`
	Host       string `json:"host"`
	Uri        string `json:"uri"`
	Headers    struct {
		UserAgent []string `json:"User-Agent"`
	} `json:"headers"`
	TLS struct {
		Resumed     bool
		Version     int
		CipherSuite int `json:"ciper_suite"`
		Proto       string
		ServerName  string `json:"server_name"`
	} `json:"tls"`
}

type Response struct {
}

type Entries []Entry
