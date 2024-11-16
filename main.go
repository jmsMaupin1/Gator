package main

import (
	"fmt"
	"github.com/jmsMaupin1/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	if err := cfg.SetUser("JT"); err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)
}
