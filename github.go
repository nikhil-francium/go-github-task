package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

//Commits contains list of commits
type Commits struct {
	Commit Commit `json:"commit"`
}

//Commit contains commit details
type Commit struct {
	Committer CommitDetails `json:"committer"`
}

//CommitDetails contains github commit details
type CommitDetails struct {
	Email      string `json:"email"`
	CommitDate string `json:"date"`
}

type commitMap map[string]int

func getGithubRequestURL(repo string, previousTime time.Time, currentTime time.Time) string {

	url := fmt.Sprintf("https://api.github.com/repos/%s/commits?since=%s&until=%s", repo, previousTime.Format(time.RFC3339), currentTime.Format(time.RFC3339))
	return url

}

func getCommits(url string) *http.Response {
	var bearer = "Bearer " + os.Getenv("GITHUB_API_KEY")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	res, _ := client.Do(req)
	return res
}

func getGithubCommits(res *http.Response) []Commits {
	var result []Commits
	body, e := ioutil.ReadAll(res.Body)
	if e == nil {
		json.Unmarshal(body, &result)
	}
	return result
}

func initializeCommitDayWiseMap() commitMap {
	commitCountMap := map[string]int{
		"Sunday":    0,
		"Monday":    0,
		"Tuesday":   0,
		"Wednesday": 0,
		"Thursday":  0,
		"Friday":    0,
		"Saturday":  0,
	}

	return commitCountMap
}

func (cm commitMap) incrementDayCount(commitDate string) {
	ti, _ := time.Parse(time.RFC3339, commitDate)
	cm[ti.Weekday().String()]++
}

func calculateCommitDayCount(weeks int, repo string) commitMap {

	var url string
	var res *http.Response
	var linkHeaders string

	cm := initializeCommitDayWiseMap()

	for {
		if linkHeaders == "" {
			previousTime, currentTime := getTimeRange(weeks)
			url = getGithubRequestURL(repo, previousTime, currentTime)
		} else {
			url = getURLFromHeader(linkHeaders)
			if url == "" {
				break
			}
		}
		res = getCommits(url)
		linkHeaders = res.Header.Get("Link")
		data := getGithubCommits(res)

		for _, commitData := range data {
			cm.incrementDayCount(commitData.Commit.Committer.CommitDate)
		}
		if linkHeaders == "" {
			break
		}
	}

	return cm
}

func getURLFromHeader(headers string) string {
	var links = map[string]string{}
	rels := strings.Split(headers, ",")
	for _, data := range rels {
		ops := strings.Split(data, ";")
		links[getKeyFromRel(ops[1])] = strings.Trim(ops[0], "<>")
	}
	if val, ok := links["next"]; ok {
		return val
	}
	return ""
}

func getKeyFromRel(key string) string {
	values := strings.Split(strings.TrimSpace(key), "=")
	return strings.Trim(values[1], `""`)
}

func (cm commitMap) printDetails() {
	for key, val := range cm {
		fmt.Println(key + " : " + strconv.Itoa(val))
	}
}

func (cm commitMap) sortDetails(sortOrder string) {
	if sortOrder == "" {
		cm.printDetails()
		return
	}
	countMap := map[int][]string{}
	keysList := []int{}

	for key, val := range cm {
		countMap[val] = append(countMap[val], key)
		keysList = append(keysList, val)
	}

	if sortOrder == "desc" {
		sort.Sort(sort.Reverse(sort.IntSlice(keysList)))
	} else {
		sort.Ints(keysList)
	}

	keysList = removeDuplicates(keysList)

	for _, key := range keysList {
		for _, val := range countMap[key] {
			fmt.Println(fmt.Sprintf("%s : %s", val, strconv.Itoa(key)))
		}
	}

}

func removeDuplicates(values []int) []int {
	result := []int{}
	isPresent := false
	for _, val := range values {
		if len(result) == 0 {
			result = append(result, val)
		} else {
			for _, item := range result {
				if item == val {
					isPresent = true
					break
				}
			}
			if isPresent == false {
				result = append(result, val)
			}
		}
		isPresent = false
	}

	return result
}
