package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		userInputSlice := strings.Fields(userInput[:len(userInput)-1])
		command := userInputSlice[0]
		arguments := userInputSlice[1:]

		builtinCommands := []string{"pwd", "cd", "echo", "type", "exit"}

		switch command {
		case "echo":
			fmt.Println(strings.Join(arguments, " "))
		case "type":
			if isBuiltin(arguments[0], builtinCommands) {
				fmt.Println(arguments[0] + " is a shell builtin")
			} else if path, err := exec.LookPath(arguments[0]); err == nil {
				fmt.Println(arguments[0] + " is " + path)
			} else {
				fmt.Println(arguments[0] + ": not found")
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println(dir)
		case "cd":
			path := arguments[0]

			if path == "~" {
				if homePath, err := os.UserHomeDir(); err != nil {
					fmt.Println("Error getting home directory:", err)
				} else {
					path = homePath
				}
			}

			if err := os.Chdir(path); err != nil {
				fmt.Println("cd: " + arguments[0] + ": No such file or directory")
			}
		case "exit":
			os.Exit(0)
		default:
			cmd := exec.Command(command, arguments...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}

}

func isBuiltin(command string, builtinCommands []string) bool {
	isFound := false
	for _, builtinCommand := range builtinCommands {
		if command == builtinCommand {
			isFound = true
			break
		}
	}
	return isFound
}
