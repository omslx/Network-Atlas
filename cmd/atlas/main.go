package main

import (
	"fmt"
	"os"

	"Network-Atlas/internal/cli"
	"Network-Atlas/internal/receiver"
	"Network-Atlas/internal/sender"
)

func main() {
	cfg := cli.Parse()

	if cfg.Send && cfg.Receive {
		fmt.Println("ERROR: -c and -s cannot be used together")
		os.Exit(1)
	}

	switch {
	case cfg.Send:
		if err := sender.Run(cfg); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	case cfg.Receive:
		if err := receiver.Run(cfg); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Usage:")
		fmt.Println("  Network-Atlas -c   Send packets")
		fmt.Println("  Network-Atlas -s   Receive packets")
	}
}
