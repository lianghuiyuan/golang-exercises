package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if strings.ToLower(text) == "exit" {
			fmt.Println("收到退出命令!")
			break
		}
		counts[text]++
		fmt.Println("输入的行内容为：", text)
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取输入时发生错误：", err)
	}
	for k, v := range counts {
		if v > 1 {
			fmt.Println("重复行：", k, "数量：", v)
		}
	}
}
