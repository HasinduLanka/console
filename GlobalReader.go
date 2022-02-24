package console

var GlobalReader *Reader

func ScanFromStandardInput() {
	GlobalReader.ScanFromStandardInput()
}

func ScanFromFile(filename string) error {
	return GlobalReader.ScanFromFile(filename)
}

func CloneScannerFile() error {
	return GlobalReader.CloseScannerFile()
}

// Read a line from Scanner
func ReadLine() string {
	return GlobalReader.ReadLine()
}

// Read a list of strings separated by sep
func ReadArray(sep string) []string {
	return GlobalReader.ReadArray(sep)
}

// Read a list of strings seperated by sep and clean the strings by removing empty/sep values
func ReadArrayClean(sep string) []string {
	return GlobalReader.ReadArrayClean(sep)
}

// Read a list of strings seperated by sep, and return a HashSet of the strings
func ReadHashset(sep string) map[string]struct{} {
	return GlobalReader.ReadHashset(sep)
}

// Read pairs of key/value in format: key1 value1 key2 value2
//
// sep : separator
// skip : skip this many entries before reading key/value pairs
func ReadPairs(sep string, skip int) map[string]string {
	return GlobalReader.ReadPairs(sep, skip)
}
