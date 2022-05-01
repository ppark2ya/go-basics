package main

import "fmt"

func main() {
	x := 40120312031

	fmt.Printf("%b\n", x) // binary
	fmt.Printf("%o\n", x) // 8진법
	fmt.Printf("%x\n", x) // 16진법
	fmt.Printf("%U\n", x) // unicode
	xAsBinary := fmt.Sprintf("%b\n", x) // format된 string return
	fmt.Println(x, xAsBinary)
}
