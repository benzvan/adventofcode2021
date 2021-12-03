package main

import (
	"github.com/stretchr/testify/assert"

	//"strings"
	"testing"
)

func TestReadDiagnostic(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		input      []int
		wantResult []int
	}{
		{
			name:       "one value",
			input:      []int{00100},
			wantResult: []int{00100, 11011},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result := readDiagnostic(tc.input)

			// then
			assert.Equal(tc.wantResult, result)
		})
	}
}

func TestGetCommons(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		name       string
		input      []uint8
		wantResult []uint8
	}{
		{
			name: "all ones",
			input: []uint8{1,1,1,1,1},
			wantResult: []uint8{1,0},
		},
		{
			name: "more than half ones",
			input: []uint8{1,0,1,0,1},
			wantResult: []uint8{1,0},
		},
		{
			name: "less than half ones",
			input: []uint8{1,0,1,0,0},
			wantResult: []uint8{0,1},
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
