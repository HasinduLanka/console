package console

import (
	"bufio"
	"os"
	"strings"
)

type Reader struct {
	Scanner     *bufio.Scanner
	ScannerFile *os.File
}

func NewReaderFromStandardInput() *Reader {
	rd := new(Reader)
	rd.ScanFromStandardInput()
	return rd
}

func NewReaderFromFile(filename string) *Reader {
	rd := new(Reader)
	rd.ScanFromFile(filename)
	return rd
}

func NewReaderFromString(s string) *Reader {
	rd := new(Reader)
	rd.Scanner = bufio.NewScanner(strings.NewReader(s))
	return rd
}

func (rd *Reader) ScanFromStandardInput() {
	rd.CloseScannerFile()
	rd.Scanner = bufio.NewScanner(os.Stdin)
}

func (rd *Reader) ScanFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	if rd.ScannerFile != nil {
		rd.ScannerFile.Close()
	}

	rd.ScannerFile = file
	rd.Scanner = bufio.NewScanner(file)
	return nil
}

func (rd *Reader) CloseScannerFile() error {
	rd.Scanner = nil
	if rd.ScannerFile != nil {
		f := rd.ScannerFile
		rd.ScannerFile = nil
		return f.Close()
	}
	return nil
}

// Read a line from Scanner
func (rd *Reader) ReadLine() string {

	if rd.Scanner == nil {
		rd.ScanFromStandardInput()
	}

	rd.Scanner.Scan()
	line := rd.Scanner.Text()

	line = strings.TrimSpace(line)
	return line
}

// Read a list of strings separated by sep
func (rd *Reader) ReadArray(sep string) []string {
	S := rd.ReadLine()
	A := strings.Split(S, sep)
	return A
}

// Read a list of strings seperated by sep and clean the strings by removing empty/sep values
func (rd *Reader) ReadArrayClean(sep string) []string {
	A := rd.ReadArray(sep)
	Out := make([]string, 0, len(A))

	for _, S := range A {
		if (len(S) > 0) && (S != sep) {
			Out = append(Out, S)
		}
	}

	return Out
}

// Read a list of strings seperated by sep, and return a HashSet of the strings
func (rd *Reader) ReadHashset(sep string) map[string]struct{} {
	A := rd.ReadArray(sep)
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
func (rd *Reader) ReadPairs(sep string, skip int) map[string]string {
	A := rd.ReadArrayClean(sep)

	Out := make(map[string]string, (len(A)-skip)/2)

	lenA := len(A)

	if (len(A)-skip)%2 == 1 {
		// odd number of entries
		lenA--

		Out[A[lenA]] = ""
	}

	for i := skip; i < lenA; i += 2 {
		Out[A[i]] = A[i+1]
	}

	return Out
}

func (rd *Reader) ReadBoolean(possitiveResponse string) bool {
	S := strings.ToLower(rd.ReadLine())
	possitiveResponse = strings.ToLower(possitiveResponse)

	return S == possitiveResponse
}
