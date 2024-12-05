package main

import (
	"bytes"
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
		{"print lines", []string{"test", "-l"}, Flags{
			PrintLines:      true,
			PrintWords:      false,
			PrintCharacters: false,
			PrintBytes:      false,
		}},
		{"print words", []string{"test", "-w"}, Flags{
			PrintLines:      false,
			PrintWords:      true,
			PrintCharacters: false,
			PrintBytes:      false,
		}},
		{"print bytes", []string{"test", "-c"}, Flags{
			PrintLines:      false,
			PrintWords:      false,
			PrintCharacters: false,
			PrintBytes:      true,
		}},
		{"print characters", []string{"test", "-m"}, Flags{
			PrintLines:      false,
			PrintWords:      false,
			PrintCharacters: true,
			PrintBytes:      false,
		}},
		{"print all", []string{"test"}, Flags{
			PrintLines:      true,
			PrintWords:      true,
			PrintCharacters: true,
			PrintBytes:      true,
		}},
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

func TestReadTxtFile(t *testing.T) {
	tmp, err := os.CreateTemp("", "file.txt")
	if err != nil {
		t.Fatalf("unable to create file %#v", err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("unable to remove tmp %v", err)
		}
	}(tmp.Name())

	content := []byte("lorem ipsum")

	_, err = tmp.Write(content)
	if err != nil {
		t.Fatalf("unable to write test data %#v", err)
	}

	if err := tmp.Close(); err != nil {
		t.Fatalf("unable to close file")
	}

	t.Run("ok read", func(t *testing.T) {
		results := readTxtFile(tmp.Name())
		if !bytes.Equal(content, results) {
			t.Fatalf("expected %#v got %#v", content, results)
		}
	})
}
