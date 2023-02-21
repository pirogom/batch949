package main

import (
	"fmt"

	"github.com/pirogom/batch949"
)

func exproc() {
	str, _ := batch949.Output("ping", "127.0.0.1")

	fmt.Println(str)

	batch949.Run("copy", "공백 있는 파일 이름.txt", "공백 있는 파일 이름2.txt")

	batch949.Run("del", "공백 있는 파일 이름.txt")

	batch949.Start("copy", "공백 있는 파일 이름2.txt", "공백 있는 파일 이름.txt")

	batch949.Start("del", "공백 있는 파일 이름2.txt")
}

func main() {
	exproc()
}
