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
		wantResult []move
		wantError  string
	}{
		{
			name:  "forward 9",
			stdin: "forward 9",
			wantResult: []move{
				{
					direction: "forward",
					distance:  9,
				},
			},
		},
		{
			name:  "3 lines",
			stdin: "forward 9\ndown 7\nup 2",
			wantResult: []move{
				{
					direction: "forward",
					distance:  9,
				},
				{
					direction: "down",
					distance:  7,
				},
				{
					direction: "up",
					distance:  2,
				},
			},
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
func TestMove(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		direction  string
		distance   int
		sub        submarine
		wantResult submarine
	}{
		{
			name:      "forward 5",
			direction: "forward",
			distance:  5,
			wantResult: submarine{
				distance: 5,
			},
		},
		{
			name:      "up 5",
			direction: "up",
			distance:  5,
			wantResult: submarine{
				aim: -5,
			},
		},
		{
			name:      "down 5",
			direction: "down",
			distance:  5,
			wantResult: submarine{
				aim: 5,
			},
		},
		{
			name:      "forward 5 at aim 3",
			direction: "forward",
			distance:  5,
			sub: submarine{
				aim: 3,
			},
			wantResult: submarine{
				aim:      3,
				distance: 5,
				depth:    15,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// given
			sub := tc.sub

			// when
			sub.move(tc.direction, tc.distance)

			// then
			assert.Equal(tc.wantResult, sub)
		})
	}
}
