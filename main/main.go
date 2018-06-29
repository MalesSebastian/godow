package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var FILE_PATH = GetPath()

func main() {
	fmt.Println("Welcome to Godow! A rudimentary download manager written in golang!")
	PrintCommands()
	reader := bufio.NewReader(os.Stdin)
	input := ""
	for input != "exit" {
		fmt.Print("> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		switch input {
		case "3":

		}
	}
}
