package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "从 %v 获取数据失败: %v\n", url, err)
			continue
		}
		fmt.Printf("状态码: %s\n", resp.Status)
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "读取数据失败: %v\n", err)
			continue
		}
		fmt.Printf("%s", body)
	}
}
