package main

import (
	"fmt"

	"github.com/dipan-ck/redis-clone-golang.git/internal/resp"
)

func main() {
	raw := []byte("*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$5\r\ndipan\r\n")

	parsed, err := resp.RespParsing(raw)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsed)
}
