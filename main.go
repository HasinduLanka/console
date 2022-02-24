package console

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

var Scanner *bufio.Scanner = nil
var ScannerFile *os.File = nil

var Writter *bufio.Writer = nil
var WritterFile *os.File = nil

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
	Scanner = bufio.NewScanner(os.Stdin)
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
		Scanner = bufio.NewScanner(os.Stdin)
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

func WriteToFile(filename string) error {

	if WritterFile != nil {
		WritterFile.Close()
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	WritterFile = file
	Writter = bufio.NewWriter(file)
	return nil

}

func CloneWritterFile() error {
	Writter = bufio.NewWriter(os.Stdout)
	if WritterFile != nil {
		f := WritterFile
		WritterFile = nil
		return f.Close()
	}
	return nil
}

// Print a string and a new line
func Print(S string) {
	PrintInline(S + "\n")
}

// Print a string without a new line
func PrintInline(S string) {
	if Writter == nil {
		Writter = bufio.NewWriter(os.Stdout)
	}

	Writter.WriteString(S)
	Writter.Flush()
}

// Print any object as a json, spaced and indented
func Log(Obj interface{}) {
	B, JErr := json.MarshalIndent(Obj, "", "\t")

	if JErr != nil {
		return
	}

	S := string(B)

	Print(S)
}

// Print any object as a single line json
func LogLine(Obj interface{}) {
	B, JErr := json.Marshal(Obj)

	if JErr != nil {
		return
	}

	S := string(B)

	Print(S)
}
