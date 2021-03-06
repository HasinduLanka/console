package console

import (
	"bufio"
	"os"
	"strings"
)

type Writer struct {
	Buff       *bufio.Writer
	WriterFile *os.File
}

type ioWriter struct {
	S *strings.Builder
}

func (iow *ioWriter) Write(p []byte) (n int, err error) {
	iow.S.WriteString(string(p))
	return len(p), nil
}

func NewWriterToStandardOutput() *Writer {
	wr := new(Writer)
	wr.WriteToStandardOutput()
	return wr
}

func NewWriterToFile(filename string) *Writer {
	wr := new(Writer)
	wr.WriteToFile(filename)
	return wr
}

func NewWriterToString() (*Writer, func() string) {

	iow := new(ioWriter)

	wr := new(Writer)
	wr.Buff = bufio.NewWriter(iow)

	return wr, func() string {
		wr.Buff.Flush()
		return iow.S.String()
	}
}

func (wr *Writer) WriteToStandardOutput() {
	wr.CloseWritterFile()
	wr.Buff = bufio.NewWriter(os.Stdout)
}

func (wr *Writer) WriteToFile(filename string) error {

	if wr.WriterFile != nil {
		wr.WriterFile.Close()
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	wr.WriterFile = file
	wr.Buff = bufio.NewWriter(file)
	return nil

}

func (wr *Writer) CloseWritterFile() error {
	wr.Buff = nil
	if wr.WriterFile != nil {
		f := wr.WriterFile
		wr.WriterFile = nil
		return f.Close()
	}
	return nil
}

// Print a string without a new line
func (wr *Writer) PrintInline(S string) {
	if wr.Buff == nil {
		wr.WriteToStandardOutput()
	}

	wr.Buff.WriteString(S)
	wr.Buff.Flush()
}
