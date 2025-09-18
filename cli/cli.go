package cli

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kencx/clp/entry"
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
		path     string
		period   string
		number   int
		crawlers bool
		notFound bool
		color    bool
		version  bool
		help     bool
	)

	flag.StringVar(&path, "file", "access.log", "path to access log")
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

	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	var entries entry.Entries

	if fi.IsDir() {
		if err := filepath.Walk(path, func(filename string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// ignore subdirectories
			if info.IsDir() {
				return nil
			}

			entry, err := process(filename)
			if err != nil {
				return err
			}

			entries = append(entries, entry...)
			return nil
		}); err != nil {
			fmt.Println(err)
			return 1
		}
	} else {
		entries, err = process(path)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}

	if err = stats.Summary(entries, number, period, crawlers, notFound, color); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func process(filename string) (entry.Entries, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var entries entry.Entries

	if strings.Contains(f.Name(), ".gz") {
		entries, err = file.DecodeGz(f)
		if err != nil {
			return nil, err
		}
	} else {
		entries, err = file.Decode(f)
		if err != nil {
			return nil, err
		}
	}

	return entries, nil
}
