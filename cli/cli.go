package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/kencx/clp/file"
	"github.com/kencx/clp/stats"
)

func Run() int {
	var (
		filename string
		period   string
		crawlers bool
	)

	flag.StringVar(&filename, "file", "access.log", "path to access log")
	flag.StringVar(&period, "time", "30d", "time period to filter")
	flag.BoolVar(&crawlers, "exclude-crawler", false, "exclude crawlers")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer f.Close()

	entries, err := file.Decode(f)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	if err = stats.Summary(entries, period, crawlers); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
