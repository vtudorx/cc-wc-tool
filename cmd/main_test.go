package main

import (
	"flag"
	"os"
	"testing"
)

func TestReadFlags(t *testing.T) {
	var tests = []struct {
		name     string
		input    []string
		expected Flags
	}{
		{"print lines", []string{"test", "-lines"}, Flags{PrintLines: true}},
		{"print no lines", []string{"test"}, Flags{PrintLines: false}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			originalArgs := os.Args
			os.Args = tc.input
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			actual := readFlags()
			if actual != tc.expected {
				t.Errorf("expected %#v, actual %#v", tc.expected, actual)
			}

			os.Args = originalArgs
		})
	}
}
