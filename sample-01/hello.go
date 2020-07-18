package main

import "fmt"

/*
실행 명령어
	$ go run hello.go

빌드 & 실행 명령어
	$ go build hello.go
	$ ./hello
*/
func main() {
	fmt.Println("hello world")

	a := 2
	b := &a
	a = 3

	fmt.Println("a: ", a)
	fmt.Println("b: ", *b)
}
