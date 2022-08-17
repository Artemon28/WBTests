package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		var command string
		fmt.Scan(&command)
		switch command {
		case "\\q":
			os.Exit(0)
		case "cd":
			var argument string
			_, err2 := fmt.Scan(&argument)
			if err2 != nil {
				log.Fatal(err2)
			}
			err := os.Chdir(argument)
			if err != nil {
				log.Fatal(err)
			}
		case "pwd":
			dir, _ := os.Getwd()
			fmt.Println(dir)
		case "echo":
			var argument string
			_, err2 := fmt.Scan(&argument)
			if err2 != nil {
				log.Fatal(err2)
			}
			fmt.Print(argument)
		case "kill":
			var argument string
			_, err2 := fmt.Scan(&argument)
			if err2 != nil {
				log.Fatal(err2)
			}
			cmd := exec.Command("cmd", "/C", "taskkill /PID ", argument)
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
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				arguments := strings.Fields(sc.Text())
				arguments = append([]string{"/C", "start"}, arguments...)
				cmd := exec.Command("cmd.exe", arguments...)
				out, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Print(string(out))
			}
		}

	}
}
