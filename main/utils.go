package main

import (
	"fmt"
	"os"
)

func PrintCommands() {
	fmt.Println("1: Download file \n" +
		"2: See current downloads \n" +
		"3: Set default path \n" +
		"exit: exit the program")
}

func SetPath(path string) string {
	_ = os.Mkdir(path)
	os.Setenv("GODOW_FILE_PATH", path)
	return path
}

func GetPath() string {
	file_path := os.Getenv("GODOW_FILE_PATH")
	if file_path == "" {
		file_path = SetPath("~/Downloads")
	}
	return file_path
}

func ReadAndTrim(
