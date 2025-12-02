package main

import (
	"fmt"
	"log"

	"github.com/Zuppah1/Blog-Aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("mantvydas")
	if err != nil {
		log.Fatalf("could not set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
