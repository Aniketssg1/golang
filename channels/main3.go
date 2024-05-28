package main

import (
	"fmt"
	"os"
	"slices"
	"testing"
	"time"
)

func main3(t *testing.T) {
	expectedLogs := []string{
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"All backups saved!",
	}

	snapshotTicker := time.Tick(time.Millisecond * 800)
	saveAfter := time.After(time.Millisecond * 2800)
	logChan := make(chan string)

	go saveBackups(snapshotTicker, saveAfter, logChan)

	testedLogs := []string{}

	for testedLog := range logChan {
		fmt.Println(testedLog)
		testedLogs = append(testedLogs, testedLog)
	}

	if slices.Equal(testedLogs, expectedLogs) {
		var a []any = []any{sliceWithBullets(expectedLogs), sliceWithBullets(testedLogs)}
		fmt.Fprintf(os.Stdout, `Test Passed: expected: %v tested: %v`, a...)
	} else {
		t.Errorf(`Test failed: expected: %v \n tested: %v`, sliceWithBullets(expectedLogs), sliceWithBullets(testedLogs))
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "	<nil>"
	}
	if len(slice) == 0 {
		return "	[]"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}
