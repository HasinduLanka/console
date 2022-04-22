package console

import (
	"bufio"
	"os"
	"strings"
)

type Reader struct {
	BufReader   *bufio.Reader
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
	rd.BufReader = bufio.NewReader(strings.NewReader(s))
	return rd
}

func (rd *Reader) ScanFromStandardInput() {
	rd.CloseScannerFile()
	rd.BufReader = bufio.NewReader(os.Stdin)
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
	rd.BufReader = bufio.NewReader(file)
	return nil
}

func (rd *Reader) CloseScannerFile() error {
	rd.BufReader = nil
	if rd.ScannerFile != nil {
		f := rd.ScannerFile
		rd.ScannerFile = nil
		return f.Close()
	}
	return nil
}

// Read a line from BufReader
func (rd *Reader) ReadLine() string {

	if rd.BufReader == nil {
		rd.ScanFromStandardInput()
	}

	line, _ := rd.BufReader.ReadString('\n')

	line = strings.TrimSpace(line)
	return line
}
