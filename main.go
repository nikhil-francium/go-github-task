package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	weeks, repo, sortOrder := getInputs()

	godotenv.Load()

	client := &http.Client{}

	commitsDayWiseCountMap := calculateCommitDayCount(weeks, repo, client)

	commitsDayWiseCountMap.sortDetails(sortOrder)

}
