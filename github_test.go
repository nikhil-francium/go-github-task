package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetGithubRequestURL(t *testing.T) {
	repo := "test"
	currentTime := time.Now()
	previousTime := currentTime.AddDate(0, -3, 0)

	expectedURL := fmt.Sprintf("https://api.github.com/repos/%s/commits?since=%s&until=%s", repo, previousTime.Format(time.RFC3339), currentTime.Format(time.RFC3339))
	actualURL := getGithubRequestURL(repo, previousTime, currentTime)

	if actualURL != expectedURL {
		t.Errorf("Invalid URL")
	}
}

func TestGetCommits(t *testing.T) {
	expectedResult := []Commits{
		{Commit: Commit{Committer: CommitDetails{Email: "nikhil@@francium.tech", CommitDate: "2006-01-02T15:04:05Z07:00"}}},
	}
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, _ := json.Marshal(expectedResult)
		rw.Write(body)
	}))
	defer server.Close()

	actualResult := getCommits(server.URL, server.Client())

	expectedByteResult, _ := json.Marshal(expectedResult)
	actualByteResult, _ := json.Marshal(getGithubCommits(actualResult))

	res := bytes.Compare(actualByteResult, expectedByteResult)
	if res != 0 {
		t.Errorf("Invalid github commits response")
	}
}

func TestGetURLFromHeader(t *testing.T) {
	sampleHeader := `<https://api.github.com/repositories/31792824/commits?since=2019-02-20T17%3A17%3A08+05%3A30&until=2020-02-24T17%3A17%3A08+05%3A30&page=2>; rel="next", <https://api.github.com/repositories/31792824/commits?since=2019-02-20T17%3A17%3A08+05%3A30&until=2020-02-24T17%3A17%3A08+05%3A30&page=150>; rel="last"`

	expectedResult := `https://api.github.com/repositories/31792824/commits?since=2019-02-20T17%3A17%3A08+05%3A30&until=2020-02-24T17%3A17%3A08+05%3A30&page=2`

	actualResult := getURLFromHeader(sampleHeader)

	if actualResult != expectedResult {
		t.Errorf("Invalid URL from header")
	}
}

func TestSortDetails(t *testing.T) {
	cm := initializeCommitDayWiseMap()
	sortOrder := "asc"
	cm.sortDetails(sortOrder)
	sortOrder = "desc"
	cm = initializeCommitDayWiseMap()
	cm["Monday"]++
	cm["Tuesday"]++
	cm.sortDetails(sortOrder)
	sortOrder = ""
	cm = initializeCommitDayWiseMap()
	cm.sortDetails(sortOrder)
}

func TestIncrementDayCount(t *testing.T) {
	sampleDate := "2006-01-02T15:04:05Z07:00"
	cm := initializeCommitDayWiseMap()
	cm.incrementDayCount(sampleDate)
	expectedResult := []Commits{
		{Commit: Commit{Committer: CommitDetails{Email: "nikhil@@francium.tech", CommitDate: "2006-01-02T15:04:05Z07:00"}}},
	}
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, _ := json.Marshal(expectedResult)
		rw.Write(body)
	}))
	defer server.Close()
	fmt.Println(calculateCommitDayCount(10, "flutter/flutter", server.Client()))
}
