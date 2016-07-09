package main

import "testing"

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
