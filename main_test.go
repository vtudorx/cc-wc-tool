package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"
	"slices"
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

	expectedContent := []string{"lorem ipsum", "ipsum lorem"}
	var byteSlice []byte
	buffer := bytes.NewBuffer(byteSlice)
	for _, str := range expectedContent {
		buffer.WriteString(str + "\n")
	}

	_, err = tmp.Write(byteSlice)
	if err != nil {
		t.Fatalf("unable to write test data %#v", err)
	}

	if err := tmp.Close(); err != nil {
		t.Fatalf("unable to close file")
	}

	t.Run("ok read", func(t *testing.T) {
		scanner := readTxtFile(tmp.Name())
		var actualContent []string
		for scanner.Scan() {
			actualContent = append(actualContent, scanner.Text())
		}
		if slices.Equal(expectedContent, actualContent) {
			t.Fatalf("expected %#v got %#v", expectedContent, actualContent)
		}
	})
}

func TestReadLines(t *testing.T) {
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

	expectedContent := []string{"lorem ipsum", "ipsum lorem"}
	var byteSlice []byte
	buffer := bytes.NewBuffer(byteSlice)
	for _, str := range expectedContent {
		buffer.WriteString(str + "\n")
	}

	byteSlice = buffer.Bytes()

	_, err = tmp.Write(byteSlice)
	if err != nil {
		t.Fatalf("unable to write test data %#v", err)
	}

	defer func() {
		if err := tmp.Close(); err != nil {
			t.Fatalf("unable to close file")
		}
	}()

	_, err = tmp.Seek(0, 0)
	if err != nil {
		t.Fatalf("unable to seek at the beginning %#v", err)
	}

	scanner := bufio.NewScanner(tmp)

	t.Run("ok read", func(t *testing.T) {
		actualLines := readLines(scanner)
		if len(expectedContent) != actualLines {
			t.Fatalf("expected %#v got %#v", len(expectedContent), actualLines)
		}
	})
}
