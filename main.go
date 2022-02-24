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

	line = strings.TrimSpace(line)
	return line
}

func ReadArray(sep string) []string {
	S := ReadLine()
	A := strings.Split(S, sep)
	return A
}

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
