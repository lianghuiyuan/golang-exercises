package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	for _, args := range os.Args[1:] {
		data, err := os.ReadFile(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, "dup3 读文件%v 失败: %v \n", args, err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if counts[args] == nil {
				counts[args] = make(map[string]int)
			}
			counts[args][line]++
		}
	}
	for file, v := range counts {
		for text, count := range v {
			if count > 1 {
				fmt.Println("文件：", file, "重复行：", text, "数量：", count)
			}
		}
	}
}
