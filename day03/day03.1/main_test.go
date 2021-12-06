package main

import (
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		input      string
		wantResult []int64
	}{
		{
			name:       "two short lines",
			input:      "10\n01",
			wantResult: []int64{0, 0},
		},
		{
			name:       "two short lines",
			input:      "10\n01",
			wantResult: []int64{0, 0},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result, _ := parseInput(strings.NewReader(tc.input))

			// then
			assert.Equal(tc.wantResult, result)
		})
	}
}

func TestGetGamma(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		input      []int64
		wantResult uint64
		wantError  string
	}{
		{
			name:       "single test",
			input:      []int64{-2, -9, 5, 3},
			wantResult: 3 * 12,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result := getGamma(tc.input)

			// then
			assert.Equal(tc.wantResult, result)
		})
	}
}
