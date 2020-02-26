package main

import (
	"flag"
	"time"
)

func getInputs() (int, string, string) {

	var weeks int
	flag.IntVar(&weeks, "weeks", 52, "no of weeks prior")
	var repo string
	flag.StringVar(&repo, "repo", "", "repo details")
	var sort string
	flag.StringVar(&sort,"sort","","commit count sort daywise")

	flag.Parse()

	return weeks, repo, sort
}

func getTimeRange(weeks int) (time.Time, time.Time) {

	currentTime := time.Now()
	totalDays := weeks * 7
	years := totalDays / 365
	months := (totalDays - years*365) / 30
	days := (totalDays - years*365) - months*30
	previousTime := currentTime.AddDate(-years, -months, -days)

	return previousTime, currentTime
}
