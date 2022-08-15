package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
		commandKeys := make(map[byte]int)
		fileName := strings.Builder{}
		for i := len(utility); i < len(str); i++ {
			if str[i] == '-' {
				i++
				for str[i] != ' ' {
					if str[i] == 'u' {
						commandKeys['u'] = 1
					}
					if str[i] == 'k' {
						i++
						commandKeys['k'], _ = strconv.Atoi(string(str[i]))
					}
					if str[i] == 'n' {
						commandKeys['n'] = 1
					}
					if str[i] == 'r' {
						commandKeys['r'] = 1
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

type linesToSort struct {
	lines []string
	keys  map[byte]int
}

func (s linesToSort) Len() int {
	return len(s.lines)
}

func (s linesToSort) Less(i, j int) bool {
	str1 := s.lines[i]
	str2 := s.lines[j]
	if column, ok := s.keys['k']; ok {
		flag1 := false
		flag2 := false
		if len(strings.Fields(str1)) < column {
			flag1 = true
		} else {
			str1 = strings.Fields(str1)[column-1]
		}
		if len(strings.Fields(str2)) < column {
			flag2 = true
		} else {
			str2 = strings.Fields(str2)[column-1]
		}
		if flag1 && !flag2 {
			return true
		} else if !flag1 && flag2 {
			return false
		}
	}
	if _, ok := s.keys['n']; ok {
		str3, _ := strconv.Atoi(str1)
		str4, _ := strconv.Atoi(str2)
		if _, ok := s.keys['r']; ok {
			return str3 > str4
		}
		return str3 < str4
	}
	if _, ok := s.keys['r']; ok {
		return str1 > str2
	}
	return str1 < str2
}

func (s linesToSort) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func sortUtil(fileName string, keys map[byte]int) {
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
	if _, ok := keys['u']; ok {
		for i := 0; i < len(lines)-1; i++ {
			if lines[i] == lines[i+1] {
				lines = append(lines[0:i], lines[i+1:]...)
				i--
			}
		}
	}
	sort.Sort(linesToSort{lines, keys})

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
