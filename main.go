package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 1 {
		str := os.Args[1]
		mas_str := strings.Split(str, " ")
		fmt.Println(len(mas_str))
	} else {
		fmt.Println("0")
	}

}
