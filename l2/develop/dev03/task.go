package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	ListenCommand()
}

func ListenCommand() {
	fmt.Printf("Write you command here:\n")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	str := myscanner.Text()
	str = strings.TrimSpace(str)
	utility := "sort"
	if strings.HasPrefix(str, utility) {
		commandKeys := make(map[byte]struct{})
		fileName := strings.Builder{}
		for i := len(utility); i < len(str); i++ {
			if str[i] == '-' {
				i++
				for str[i] != ' ' {
					if str[i] == 'u' {
						commandKeys['u'] = struct{}{}
					}
					if str[i] == 'k' {
						commandKeys['k'] = struct{}{}
					}
					if str[i] == 'n' {
						commandKeys['n'] = struct{}{}
					}
					if str[i] == 'r' {
						commandKeys['r'] = struct{}{}
					}
					i++
				}
			}
			if str[i] != ' ' {
				fileName.WriteRune(rune(str[i]))
			}
		}
		if _, err := os.Stat(fileName.String()); os.IsNotExist(err) {
			log.Fatal("file" + fileName.String() + " does not exist")
		}
		sortUtil(fileName.String(), commandKeys)
	}
}

func sortUtil(fileName string, keys map[byte]struct{}) {
	lines := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	sort.Strings(lines)

	writeFile, err := os.Create(fileName[0:len(fileName)-4] + "_sorted.txt")
	defer writeFile.Close()
	if err != nil {
		log.Fatal("Unable to create file:", err)
	}
	for _, v := range lines {
		_, err := writeFile.WriteString(v + "\n")
		if err != nil {
			log.Fatal("Can't write string to output file")
		}
	}
}
