package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().UTC()

	fmt.Println(now)

	format := now.Format("2006-01-02 15:04:05")

	const i=5
	fmt.Println(&10)
	fmt.Println(&i)
	fmt.Println(format)
}
