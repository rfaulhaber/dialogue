package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"
)

const delay = 50 * time.Millisecond

func main() {
	args := os.Args[1:]

	u, err := user.Current()

	if err != nil {
		log.Fatalln("could not get current user")
	}

	name := u.Name

	fmt.Println(name + ":")

	if len(args) > 1 {
		for _, str := range args {
			printStr(str)
			fmt.Println()
		}
	} else {
		// figure out how to print just first arg
		printStr(args[0])
	}
}

func printStr(s string) {
	for _, c := range s {
		fmt.Print(string(c))
		time.Sleep(delay)
	}
}
