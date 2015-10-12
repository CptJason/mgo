package main

import (
	"fmt"
	"github.com/fumiyas/go-tty"
	"os"
)

func main() {
	da1, err := tty.GetDeviceAttributes1(os.Stdout)
	if err == nil && da1[tty.DA1_SIXEL] {
		fmt.Printf("\x1BPq#0;2;0;0;0#1;2;100;100;0#2;2;0;0;100#1~~@@vv@@~~@@~~$#2??}}GG}}??}}??-#1!14@\x1B\\")
	} else {
		fmt.Printf("Hi\n")
	}
}
