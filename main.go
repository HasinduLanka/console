package console

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var Scanner *bufio.Scanner = nil

func ReadLine() string {

	if Scanner == nil {
		Scanner = bufio.NewScanner(os.Stdin)
	}

	Scanner.Scan()
	line := Scanner.Text()

	return line
}

func ReadArray(sep string) []string {
	S := ReadLine()
	A := strings.Split(S, sep)
	return A
}

func Print(S string) {
	fmt.Println(S)
}

func PrintInline(S string) {
	fmt.Print(S)
}

func Log(Obj interface{}) {
	B, JErr := json.MarshalIndent(Obj, "", "\t")

	if JErr != nil {
		fmt.Println(Obj)
		return
	}

	S := string(B)

	fmt.Println(S)
}

func LogLine(Obj interface{}) {
	B, JErr := json.Marshal(Obj)

	if JErr != nil {
		fmt.Println(Obj)
		return
	}

	S := string(B)

	fmt.Println(S)
}
