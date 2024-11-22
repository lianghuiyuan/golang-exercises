package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "从 %v 获取数据失败: %v\n", url, err)
			continue //继续下一个循环
			// os.Exit(1)  //直接退出程序
		}
		fmt.Println("11111111111111111: %s", url)
		b, err := io.ReadAll(resp.Body)   
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "解析 body 数据失败: %v\n", err)
			continue // 继续下一个循环
			// os.Exit(1) // 直接退出程序
		}
		fmt.Println(b)
		// fmt.Printf("%s", b)
	}
}
