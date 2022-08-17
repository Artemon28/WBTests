package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	shell()
}

func shell() {
	for {
		var command string
		fmt.Scan(&command)
		switch command {
		case "\\q":
			os.Exit(0)
		case "cd":
			arguments, err := getArgs()
			if err != nil {
				log.Fatal(err)
			}
			err = os.Chdir((*arguments)[0])
			if err != nil {
				log.Fatal(err)
			}
		case "pwd":
			dir, _ := os.Getwd()
			fmt.Println(dir)
		case "echo":
			arguments, err := getArgs()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(arguments)
		case "kill":
			arguments, err := getArgs()
			if err != nil {
				log.Fatal(err)
			}
			*arguments = append([]string{"/C", "taskkill /PID "}, *arguments...)
			cmd := exec.Command("cmd", *arguments...)
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		case "ps":
			cmd := exec.Command("cmd", "/C", "tasklist")
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(out))
		case "fork":
			var procAttr os.ProcAttr
			procAttr.Files = []*os.File{os.Stdin,
				os.Stdout, os.Stderr}
			_, err := os.StartProcess(os.Args[0], os.Args, &procAttr)
			if err != nil {
				log.Fatal(err)
			}
		case "exec":
			arguments, err := getArgs()
			if err != nil {
				log.Fatal(err)
			}
			*arguments = append([]string{"/C", "start"}, *arguments...)
			cmd := exec.Command("cmd.exe", *arguments...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(out))
		}

	}
}

func getArgs() (*[]string, error) {
	sc := bufio.NewScanner(os.Stdin)
	var arguments []string
	for sc.Scan() {
		arguments = strings.Fields(sc.Text())
		return &arguments, nil
	}
	return nil, errors.New("Wrong arguments")
}
