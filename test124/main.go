package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// GOOS=linux GOARCH=amd64 go build -o ../bin/test124_124 .

func init() {
	fmt.Printf("Go Version: %s(%s)\n", runtime.Version(), "1.24")
}

func main() {
	data := "hello"
	endpoint := "https://bucket.oss-cn-hangzhou.aliyuncs.com"
	req, err := http.NewRequest(http.MethodPut, endpoint, strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)
	_, _ = io.Copy(os.Stdout, resp.Body)
}
