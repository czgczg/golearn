package main

import (
	"fmt"
	"reflect"
	"test"
)

func main() {
	hello := test.SayHello()
	count := 18
	fmt.Println(hello)

	fmt.Println(reflect.TypeOf(hello))
	fmt.Println(&hello)
	fmt.Println(reflect.TypeOf(count))

	const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009693147180559945309417232121458176568075500134360255254120680009;
	var a = float32(Ln2)

	fmt.Println(Ln2)
	fmt.Println(a)
}
