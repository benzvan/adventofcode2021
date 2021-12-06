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
		wantResult [][]int
	}{
		{
			name:  "two short lines",
			input: "10\n01",
			wantResult: [][]int{
					[]int{1, 0},
					[]int{0, 1},
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
