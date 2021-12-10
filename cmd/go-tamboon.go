package cmd

import (
	"fmt"
	"log"
	"opn-challenge/internal/config"
)

func Execute() {
	cfg, err := config.TryGetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
}
