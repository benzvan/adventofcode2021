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
		stdin      string
		wantResult []int
		wantError  string
	}{
		{
			name:       "two numbers",
			stdin:      "5\n6",
			wantResult: []int{5, 6},
		},
		{
			name:      "two numbers and one string",
			stdin:     "5\n6\nfoo",
			wantError: "strconv.Atoi: parsing \"foo\": invalid syntax",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result, err := parseInput(strings.NewReader(tc.stdin))

			// then
			if len(tc.wantError) > 0 {
				assert.Equal(tc.wantError, err.Error())
			} else {

				assert.Equal(tc.wantResult, result)
			}
		})
	}
}
func TestCountIncreases(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		list       []int
		wantResult int
	}{
		{
			name:       "two increasing values",
			list:       []int{5, 6},
			wantResult: 1,
		},
		{
			name:       "two decreasing values",
			list:       []int{5, 4},
			wantResult: 0,
		},
		{
			name:       "three increasing values",
			list:       []int{5, 6, 7},
			wantResult: 2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result := countIncreases(tc.list)

			// then
			assert.Equal(tc.wantResult, result)
		})
	}
}
