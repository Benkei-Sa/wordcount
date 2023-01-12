package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	str := os.Args
	if len(str) == 1 || os.Args[1] == "" {
		fmt.Println(0)
	} else {
		mas_str := strings.Split(str[1], " ")
		fmt.Println(len(mas_str))
	}
}
