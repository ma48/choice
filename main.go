package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 乱数のタネを初期化
	rand.Seed(time.Now().UnixNano())

	// コマンドライン引数の数をカウント。設定されていなければエラーに
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s choice1 choice2...", os.Args[0])
		os.Exit(1)
	}

	// 乱数で洗濯したものを表示
	fmt.Printf(os.Args[rand.Intn(c)+1])
}
