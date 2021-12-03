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
		wantResult [][]uint
	}{
		{
			name:  "two short lines",
			input: "10\n01",
			wantResult: [][]uint{
				{1, 0},
				{0, 1},
			},
		},
		{
			name:  "six short lines",
			input: "100\n011\n001\n110\n101\n111",
			wantResult: [][]uint{
				{1, 0, 0, 1, 1, 1},
				{0, 1, 0, 1, 0, 1},
				{0, 1, 1, 0, 1, 1},
			},
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

func TestGetCommons(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		input      []uint
		wantResult []uint
	}{
		{
			name:       "all ones",
			input:      []uint{1, 1, 1, 1, 1},
			wantResult: []uint{1, 0},
		},
		{
			name:       "more than half ones",
			input:      []uint{1, 0, 1, 0, 1},
			wantResult: []uint{1, 0},
		},
		{
			name:       "less than half ones",
			input:      []uint{1, 0, 1, 0, 0},
			wantResult: []uint{0, 1},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result := getCommons(tc.input)

			// then
			assert.Equal(tc.wantResult, result)
		})
	}
}
