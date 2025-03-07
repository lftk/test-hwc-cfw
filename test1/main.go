package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func init() {
	fmt.Println(runtime.Version())
}

func main() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	resp.Header.Write(os.Stdout)
}
