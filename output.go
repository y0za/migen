package main

import (
	"log"
	"regexp"
	"strconv"
)

var (
	// currentTime = time.Now
	countRegexp *regexp.Regexp
)

func init() {
	var err error
	countRegexp, err = regexp.Compile(`^(\d+).*\.sql$`)
	if err != nil {
		log.Fatal("fail to compile migration count regexp")
	}
}

func nextMigrationCount(fileNames []string) int64 {
	var maxCount int64 = 0
	for _, name := range fileNames {
		matches := countRegexp.FindStringSubmatch(name)
		if len(matches) < 2 {
			continue
		}
		c, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			continue
		}
		if c > maxCount {
			maxCount = c
		}
	}
	return maxCount + 1
}
