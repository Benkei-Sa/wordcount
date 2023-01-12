package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	mas_str := strings.Split(str, " ")
	fmt.Println(len(mas_str))
}
