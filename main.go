package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
)

const defaultFormat = "2006-01-02T15:04:05.999Z07:00"

var (
	argFormat = kingpin.Flag("format", "A custom date/time format to use. See README.md.").Short('f').Default(defaultFormat).String()
	argLevel  = kingpin.Flag("level", "An optional log level to add in addition to the timestamp.").Short('l').String()
	argNoUtc  = kingpin.Flag("no-utc", "Disables converting the time to UTC first.").Bool()
)

func now() string {
	if *argNoUtc {
		return time.Now().Format(*argFormat)
	} else {
		return time.Now().UTC().Format(*argFormat)
	}
}

func main() {
	kingpin.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if *argLevel == "" {
			fmt.Println(now(), scanner.Text())
		} else {
			fmt.Println(now(), *argLevel, scanner.Text())
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "logstamper error:", scanner.Err())
		os.Exit(1)
	}
}
