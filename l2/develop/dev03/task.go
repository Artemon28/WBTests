package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	k int
	n bool
	r bool
	u bool
}

func main() {

	var fl flags
	flag.IntVar(&fl.k, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&fl.n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&fl.r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&fl.u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	fileName := flag.Args()[1]
	sortUtil(fileName, fl)
}

type linesToSort struct {
	lines []string
	keys  flags
}

func (s linesToSort) Len() int {
	return len(s.lines)
}

func (s linesToSort) Less(i, j int) bool {
	str1 := s.lines[i]
	str2 := s.lines[j]
	if s.keys.k > 0 {
		flag1 := false
		flag2 := false
		if len(strings.Fields(str1)) < s.keys.k {
			flag1 = true
		} else {
			str1 = strings.Fields(str1)[s.keys.k-1]
		}
		if len(strings.Fields(str2)) < s.keys.k {
			flag2 = true
		} else {
			str2 = strings.Fields(str2)[s.keys.k-1]
		}
		if flag1 && !flag2 {
			return true
		} else if !flag1 && flag2 {
			return false
		}
	}
	if s.keys.n {
		str3, _ := strconv.Atoi(str1)
		str4, _ := strconv.Atoi(str2)
		if s.keys.r {
			return str3 > str4
		}
		return str3 < str4
	}
	if s.keys.r {
		return str1 > str2
	}
	return str1 < str2
}

func (s linesToSort) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func sortUtil(fileName string, fl flags) {
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
	if fl.u {
		for i := 0; i < len(lines)-1; i++ {
			if lines[i] == lines[i+1] {
				lines = append(lines[0:i], lines[i+1:]...)
				i--
			}
		}
	}
	sort.Sort(linesToSort{lines, fl})

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
