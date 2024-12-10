package main

import (
	"fmt"
	"os"
	"time"
	_ "time/tzdata"
)

func main() {
	var timeFormat string
	if format, ok := os.LookupEnv("TIME_FORMAT"); ok {
		timeFormat = format
	} else {
		timeFormat = time.UnixDate
	}

	var location *time.Location
	if tz, ok := os.LookupEnv("TZ"); ok {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in time.LoadLocation: %v\n", err)
			location = time.Local
		} else {
			location = loc
		}
	} else {
		location = time.Local
	}

	for {
		now := time.Now()

		fmt.Println(now.Format(timeFormat))

		nextMin := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()+1, 0, 0, location)
		time.Sleep(nextMin.Sub(now))
	}
}
