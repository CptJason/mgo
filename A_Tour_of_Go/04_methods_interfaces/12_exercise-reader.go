package main

import "golang.org/x/tour/reader"

//
// Exercise: Readers
//
// Implement a Reader type that emits
// an infinite stream of the ASCII character 'A'.
//
// type MyReader struct{
// 	s        string
// 	i        int64 // current reading index
// 	prevRune int   // index of previous rune; or < 0
// }
//
// func NewReader(s string) *MyReader {
// 	return &MyReader{s, 0, -1}
// }
//
// // TODO: Add a Read([]byte) (int, error) method to MyReader.
// func (mr *MyReader) Read([]byte) (int, error) {
// 	mr := MyReader{'A'}
// 	b :=make([]byte,8)
// 	n, err := mr.Read(b)
// 	return n, err
// }
//
// func main() {
// 	reader.Validate(MyReader{})
// }

type MyReader struct{}

// object of MyReader read string from itself, and
// into b of slice.

func (r MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
