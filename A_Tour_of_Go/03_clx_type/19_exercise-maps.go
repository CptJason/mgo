package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	w := strings.Fields(s)
	fmt.Println(w)
	for _, v := range w {
		m[v] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
