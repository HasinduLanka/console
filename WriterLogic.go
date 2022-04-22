package console

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Print a string and a new line
func (wr *Writer) Print(S string) {
	wr.PrintInline(S + "\n")
}

// Print an integer and a new line
func (wr *Writer) PrintInt(I int) {
	wr.Print(strconv.Itoa(I))
}

// Print an integer without a new line
func (wr *Writer) PrintIntInline(I int) {
	wr.PrintInline(strconv.Itoa(I))
}

// Print an array of strings, separated by sep
func (wr *Writer) PrintArray(A []string, sep string) {
	wr.Print(strings.Join(A, sep))
}

// Print an array of integers, separated by sep
func (wr *Writer) PrintIntArray(A []int, sep string) {
	SA := make([]string, len(A))

	for i, v := range A {
		SA[i] = strconv.Itoa(v)
	}

	wr.PrintArray(SA, sep)
}

// Print any object as a json, spaced and indented
func (wr *Writer) Log(Obj interface{}) {
	B, JErr := json.MarshalIndent(Obj, "", "\t")

	if JErr != nil {
		return
	}

	S := string(B)

	wr.Print(S)
}

// Print any object as a single line json
func (wr *Writer) LogLine(Obj interface{}) {
	B, JErr := json.Marshal(Obj)

	if JErr != nil {
		return
	}

	S := string(B)

	wr.Print(S)
}
