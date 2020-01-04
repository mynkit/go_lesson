package main

import (
	"go_lesson/section14/config"
	"go_lesson/section14/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
