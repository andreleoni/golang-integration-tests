package main

import (
	"build-tags-integration-tests/internal/request"
	"fmt"
)

func main() {
	fmt.Println("result: ", request.NewRequest().KeepAlive("https://www.google.com.br"))
}
