package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getPath() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v -> ", wd)
}

func main() {
	// create a scanner and get the path of the current directory
	sc := bufio.NewScanner(os.Stdin)
	getPath()

	// read commands from the console
	for sc.Scan() {
		text := sc.Text()
		if text == `\quit` {
			os.Exit(0)
		}
		// check if the input from the console is a command and execute it
		parsePipeline(os.Stdout, text)
		// refresh the current directory path
		getPath()
	}

}

func parsePipeline(stdout *os.File, str string) {
	// split the input string into individual commands separated by '|'
	commands := strings.Split(str, "|")
	var result string
	var err error

	// execute the specified commands
	for _, c := range commands {

		result, err = parseCommand(result, c)
		if err != nil {
			fmt.Fprintln(stdout, err)
		}
	}
	// print the result if the command exists and is executed successfully
	if result != "" {
		fmt.Fprintln(stdout, result)
	}
}

func parseCommand(input string, c string) (string, error) {
	// split the command string into fields
	str := strings.Fields(c)

	// check if there are no commands (empty command)
	if len(str) == 0 {
		return "", errors.New("empty command")
	}

	// append the input to the command unless it's a '{'
	if input != "{" {
		str = append(str, input)
	}

	// execute the specified command
	switch str[0] {
	case "cd":
		// if the number of parameters is not 3(2) - 'cd'
		//  and the directory to navigate to, throw an error.
		if len(str) != 3 {
			fmt.Println(str, len(str))
			return "", errors.New("cd: requires 1 parameter")
		}

		// change to the specified directory
		if err := os.Chdir(str[1]); err != nil {
			return "", err
		}

	// if the number of parameters is not 2, throw an error
	case "pwd":
		if len(str) != 2 {
			return "", errors.New("pwd: unused parameter")
		}

		// get the current working directory and return it
		result, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return result, nil

	case "echo":
		// if the number of parameters is not 3, throw an error
		if len(str) != 3 {
			return "", errors.New("echo: requires 1 parameter")
		}
		return str[1], nil

	case "ps":
		// get the current processes
		processes, err := ps.Processes()
		if err != nil {
			return "", err
		}

		// build a string representation of processes
		var strForProc strings.Builder
		strForProc.WriteString("\tPID\tCMD\n")

		for _, proc := range processes {
			strForProc.WriteString(fmt.Sprintf("\t%v\t%v\n", proc.Pid(), proc.Executable()))
		}
		return strForProc.String(), nil
	}

	return "", nil
}
