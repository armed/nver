package cmd

import (
	"log"
)

func validateArgsNum(args []string, num int) {
	if len(args) < num {
		log.Fatal("Not enough arguments")
	}
}
