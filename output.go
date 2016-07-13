package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const migrationTemplate = "-- +migrate Up\n\n-- +migrate Down"

var (
	countRegexp *regexp.Regexp
)

func init() {
	var err error
	countRegexp, err = regexp.Compile(`^(\d+).*\.sql$`)
	if err != nil {
		log.Fatal("fail to compile migration count regexp")
	}
}

func createFile(name, dirPath string) error {
	abs, err := filepath.Abs(dirPath)
	if err != nil {
		return err
	}
	output := filepath.Join(abs, name)
	return ioutil.WriteFile(output, []byte(migrationTemplate), 0666)
}

func fileName(format, name string) (string, error) {
	files, err := existingFiles(".")
	if err != nil {
		return "", err
	}

	count := nextMigrationCount(files)
	now := time.Now()
	return fileNameWithSystemInfo(format, name, count, now)
}

func fileNameWithSystemInfo(format, name string, count int64, now time.Time) (string, error) {
	var prefix string

	switch format {
	case "counter":
		prefix = fmt.Sprintf("%d", count)
	case "unix":
		prefix = fmt.Sprintf("%d", now.Unix())
	case "none":
		if name == "" {
			return "", errors.New("Not allowed empty name when format is none")
		}
	case "date":
		fallthrough
	default:
		prefix = now.Format("20060102150405")
	}
	return fmt.Sprintf("%s%s.sql", prefix, name), nil
}

func existingFiles(dirPath string) ([]string, error) {
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if dirInfo.IsDir() == false {
		msg := fmt.Sprintf("%s is not directory", dirPath)
		return nil, errors.New(msg)
	}

	fileInfoList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, fi := range fileInfoList {
		if fi.IsDir() == false {
			fileNames = append(fileNames, fi.Name())
		}
	}
	return fileNames, nil
}

func nextMigrationCount(fileNames []string) int64 {
	var maxCount int64
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
