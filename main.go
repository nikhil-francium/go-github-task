package main

import (
	"github.com/joho/godotenv"
)

func main() {

	weeks, repo, sortOrder := getInputs()

	godotenv.Load()

	commitsDayWiseCountMap := calculateCommitDayCount(weeks, repo)

	commitsDayWiseCountMap.sortDetails(sortOrder)

}
