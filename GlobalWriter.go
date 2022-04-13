package console

var GlobalWriter *Writer = NewWriterToStandardOutput()

func WriteToStandardOutput() {
	GlobalWriter.WriteToStandardOutput()
}

func WriteToFile(filename string) error {
	return GlobalWriter.WriteToFile(filename)
}

func CloseWritterFile() error {
	return GlobalWriter.CloseWritterFile()
}

// Print a string and a new line
func Print(S string) {
	GlobalWriter.Print(S)
}

// Print a string without a new line
func PrintInline(S string) {
	GlobalWriter.PrintInline(S)
}

// Print any object as a json, spaced and indented
func Log(Obj interface{}) {
	GlobalWriter.Log(Obj)
}

// Print any object as a single line json
func LogLine(Obj interface{}) {
	GlobalWriter.LogLine(Obj)
}
