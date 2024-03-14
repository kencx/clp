package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/kencx/clp/file"
	"github.com/kencx/clp/stats"
)

var tag = "v0.1.0"

const (
	helpText = `usage: clp [flags]

  Flags:
    -file        path to access log
    -time        time period to filter
	-number      print top N
    -no-crawler  exclude crawlers
	-404         exclude 404 status
    -version     print version
    -color       print with color
    -help        print help
`
)

func Run() int {
	var (
		filename string
		period   string
		number   int
		crawlers bool
		notFound bool
		color    bool
		version  bool
		help     bool
	)

	flag.StringVar(&filename, "file", "access.log", "path to access log")
	flag.StringVar(&period, "time", "30d", "time period to filter")
	flag.IntVar(&number, "number", 5, "top N")
	flag.BoolVar(&crawlers, "no-crawler", false, "exclude crawlers")
	flag.BoolVar(&notFound, "404", false, "filter 404 status codes")
	flag.BoolVar(&color, "color", false, "print with color")
	flag.BoolVar(&version, "version", false, "print version")
	flag.BoolVar(&help, "help", false, "print help")

	flag.Usage = func() { os.Stdout.Write([]byte(helpText)) }
	flag.Parse()

	if help {
		fmt.Print(helpText)
		return 0
	}

	if version {
		fmt.Println(tag)
		return 0
	}

	if len(os.Args) < 1 {
		fmt.Print(helpText)
		return 1
	}

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

	if err = stats.Summary(entries, number, period, crawlers, notFound, color); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
