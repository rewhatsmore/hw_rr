package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reverse := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reverse)
}
