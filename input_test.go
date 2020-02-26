package main

import (
	"os"
	"testing"
)

func TestGetInputs(t *testing.T) {
	//Let the weeks=20, repo=test, sort=asc
	os.Args = []string{"cmd", "-weeks=20", "-repo=test", "-sort=asc"}
	weeks, repo, sortOrder := getInputs()
	if weeks != 20 || repo != "test" || sortOrder != "asc" {
		t.Errorf("Invalid command line arguments")
	}
}
