package console

import (
	"bufio"
	"os"
	"strings"
)

var Scanner *bufio.Scanner = nil
var ScannerFile *os.File = nil

func ScanFromStandardInput() {
	CloneScannerFile()
	Scanner = bufio.NewScanner(os.Stdin)
}

func ScanFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	if ScannerFile != nil {
		ScannerFile.Close()
	}

	ScannerFile = file
	Scanner = bufio.NewScanner(file)
	return nil
}

func CloneScannerFile() error {
	Scanner = nil
	if ScannerFile != nil {
		f := ScannerFile
		ScannerFile = nil
		return f.Close()
	}
	return nil
}

// Read a line from Scanner
func ReadLine() string {

	if Scanner == nil {
		ScanFromStandardInput()
	}

	Scanner.Scan()
	line := Scanner.Text()

	line = strings.TrimSpace(line)
	return line
}

// Read a list of strings separated by sep
func ReadArray(sep string) []string {
	S := ReadLine()
	A := strings.Split(S, sep)
	return A
}

// Read a list of strings seperated by sep and clean the strings by removing empty/sep values
func ReadArrayClean(sep string) []string {
	A := ReadArray(sep)
	Out := make([]string, 0, len(A))

	for _, S := range A {
		if (len(S) > 0) && (S != sep) {
			Out = append(Out, S)
		}
	}

	return Out
}

// Read a list of strings seperated by sep, and return a HashSet of the strings
func ReadHashset(sep string) map[string]struct{} {
	A := ReadArray(sep)
	Out := make(map[string]struct{}, len(A))

	for _, S := range A {
		Out[S] = struct{}{}
	}

	delete(Out, sep)
	delete(Out, "")

	return Out
}

// Read pairs of key/value in format: key1 value1 key2 value2
//
// sep : separator
// skip : skip this many entries before reading key/value pairs
func ReadPairs(sep string, skip int) map[string]string {
	A := ReadArrayClean(sep)

	Out := make(map[string]string, (len(A)-skip)/2)

	for i := skip; i < len(A); i += 2 {
		Out[A[i]] = A[i+1]
	}

	return Out
}
