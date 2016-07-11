package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestExistingFiles(t *testing.T) {
	fileNames := map[string]bool{
		"1.sql":              true,
		"hoge.sql":           true,
		"20060102150405.sql": true,
		".sql":               true,
	}

	dir, err := ioutil.TempDir("", "migrations")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	for n, _ := range fileNames {
		path := fmt.Sprintf("%s/%s", dir, n)
		ioutil.WriteFile(path, []byte{}, 0600)
	}

	files, err := existingFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range files {
		_, exists := fileNames[n]
		if !exists {
			t.Errorf("Unexpected file name %s", n)
			continue
		}
		fileNames[n] = false
	}
	for n, leaked := range fileNames {
		if leaked {
			t.Errorf("return value does't include %s", n)
		}
	}
}

func TestNextMigrationCount(t *testing.T) {
	tests := []struct {
		names    []string
		expected int64
	}{
		{[]string{"foo.sql"}, 1},
		{[]string{"1.txt"}, 1},
		{[]string{"1.sql"}, 2},
		{[]string{"1foo.sql"}, 2},
		{[]string{"1.sql", "2.sql"}, 3},
		{[]string{"1.sql", "2.sql", "4.sql"}, 5},
		{[]string{"1.sql", "20060102150405.sql"}, 20060102150406},
		{[]string{"1.sql", "4foo.sql", "bar.sql", "20060102150405.txt"}, 5},
	}

	for _, tt := range tests {
		count := nextMigrationCount(tt.names)
		if count != tt.expected {
			t.Errorf("Expected count %d, actual %d, when %v", tt.expected, count, tt.names)
		}
	}
}
