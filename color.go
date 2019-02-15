package main

import (
	"fmt"
	"strconv"
)

func hex2int(str string) int {
	result, _ := strconv.ParseInt(str, 16, 64)
	return int(result)
}

func colorize(text string, h string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", hex2int(h[1:3]), hex2int(h[3:5]), hex2int(h[5:7]), text)
}
