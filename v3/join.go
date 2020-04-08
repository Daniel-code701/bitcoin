package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"Hello", "world"}
	//res := strings.Join(str1, "+")
	res := strings.Join(str1, "")
	fmt.Printf("%s\n", res)

	//res1 := bytes.Join([][]byte{[]byte("Hello"), []byte("World")}, []byte("+"))
	res1 := bytes.Join([][]byte{[]byte("Hello"), []byte("World")}, []byte(""))
	fmt.Printf("%s\n", res1)
}
