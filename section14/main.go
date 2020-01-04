package main

import (
	"fmt"
	"go_lesson/section14/config"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
