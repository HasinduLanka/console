package console

import (
	"strconv"
	"strings"
)

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

// Read pairs of key/value in format: key1 111 key2 222
//
// sep : separator
// skip : skip this many entries before reading key/value pairs
func (rd *Reader) ReadStringIntPairs(sep string, skip int) (map[string]int, error) {
	A := rd.ReadArrayClean(sep)

	Out := make(map[string]int, (len(A)-skip)/2)

	lenA := len(A)

	if (len(A)-skip)%2 == 1 {
		// odd number of entries
		lenA--

		Out[A[lenA]] = 0
	}

	for i := skip; i < lenA; i += 2 {
		n, err := strconv.Atoi(A[i+1])
		if err != nil {
			return nil, err
		}
		Out[A[i]] = n
	}

	return Out, nil
}

func (rd *Reader) ReadBoolean(possitiveResponse string) bool {
	S := strings.ToLower(rd.ReadLine())
	possitiveResponse = strings.ToLower(possitiveResponse)

	return S == possitiveResponse
}

// Read one integer in one line
func (rd *Reader) ReadInt() (int, error) {
	S := rd.ReadLine()
	I, err := strconv.Atoi(S)
	return I, err
}

// Read a list of integers separated by sep
func (rd *Reader) ReadIntArray(sep string) ([]int, error) {
	A := rd.ReadArrayClean(sep)
	Out := make([]int, len(A))

	for i, S := range A {
		n, err := strconv.Atoi(S)
		if err != nil {
			return nil, err
		}
		Out[i] = n
	}

	return Out, nil
}
