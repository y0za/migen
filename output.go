package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func existingFiles(dirPath string) ([]string, error) {
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if dirInfo.IsDir() == false {
		msg := fmt.Sprintf("%d is not directory", dirPath)
		return nil, errors.New(msg)
	}

	fileInfoList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	fileNames := make([]string, 0)
	for _, fi := range fileInfoList {
		if fi.IsDir() == false {
			fileNames = append(fileNames, fi.Name())
		}
	}
	return fileNames, nil
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
