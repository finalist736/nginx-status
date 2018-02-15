package main

import "testing"

type testdata struct {
	result bool
	slice []string
	value string
}

var tests = []testdata {
	{value: "1", result: true, slice: []string{"1", "3"}},
	{value: "15test", result: true, slice: []string{"1", "3", "15test"}},
	{value: "14test", result: false, slice: []string{"1", "3", "15test"}},
	{value: "3", result: false, slice: []string{"1", "15test", "55", "66", "abc"}},
}

func TestIsValueInList(t *testing.T) {
	for _, iterator := range tests {
		if IsValueInList(iterator.value, iterator.slice) != iterator.result {
			t.Error("incorrect result")
		}
	}
}
