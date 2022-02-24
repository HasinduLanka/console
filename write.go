package console

import (
	"bufio"
	"encoding/json"
	"os"
)

var Writter *bufio.Writer = nil
var WritterFile *os.File = nil

func WriteToStandardOutput() {
	CloseWritterFile()
	Writter = bufio.NewWriter(os.Stdout)
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

func CloseWritterFile() error {
	Writter = nil
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
		WriteToStandardOutput()
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
