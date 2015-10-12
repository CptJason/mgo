package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//
// 维基解释rot13:
// https://zh.wikipedia.org/wiki/ROT13
//
// Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader,
// modifying the stream in some way.
//
// For example, the gzip.NewReader function takes an io.Reader
// (a stream of compressed data) and returns a *gzip.Reader that
// also implements io.Reader (a stream of the decompressed data).
//
// The rot13Reader type is provided for you.
// Make it an io.Reader by implementing its Read method.
//
// 练习：rot13Reader
// 通常是用 io.Reader 封装另一个io.Reader的方式,
// 来修改io流。
//
// 例如，gzip.NewReader 函数接受 io.Reader（压缩的数据流）并且返回同样实现了
// io.Reader 的 *gzip.Reader（解压缩后的数据流）。
//
// 编写一个实现了 io.Reader 的 rot13Reader， 并从一个 io.Reader 读取，
// 利用 rot13 代换密码对数据流进行修改。
//
// 已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 io.Reader。
//
type rot13Reader struct {
	r io.Reader
}

// Implement a rot13Reader that implements io.Reader and
// reads from an io.Reader,
// modifying the stream by applying the rot13 substitution cipher
// to all alphabetical characters.
func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	// for each letter read from io.Reader
	for i := 0; i < len(p); i++ {
		// if the letter's index is between A -N, add 13 to its index
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			fmt.Println(string(p[i]))
			p[i] += 13
			// if the letter's index is between M - Z, substract 13 from its index
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] < 'z') {
			fmt.Println(string(p[i]))
			//			fmt.Println("##################")
			p[i] -= 13
		}
	}
	return
}

func main() {
	// use strings initial NewReader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//
	r := rot13Reader{s}
	//
	io.Copy(os.Stdout, &r)
}
