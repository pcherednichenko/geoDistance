package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithNormalLines(t *testing.T) {
	lines := [][]string{
		{"id", "lat", "lon"},
		{"382582", "37.1768672", "-3.608897"},
		{"482365", "52.36461880000000235", "4.93169289999999982"},
	}
	result, err := readFromLines(lines)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 37.1768672, result[382582][0])
	assert.Equal(t, -3.608897, result[382582][1])
	assert.Equal(t, 52.36461880000000235, result[482365][0])
	assert.Equal(t, 4.93169289999999982, result[482365][1])
}

func TestWithBadLines(t *testing.T) {
	lines := [][]string{
		{"id", "lat", "lon"},
		{"382582", "test", "-3.608897"},
		{"482365", "52.36461880000000235", "4.93169289999999982"},
	}
	_, err := readFromLines(lines)
	assert.Error(t, err)
}
