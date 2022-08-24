package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func main() {
	var fl flags
	flag.IntVar(&fl.A, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&fl.B, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&fl.C, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&fl.c, "c", false, "количество строк")
	flag.BoolVar(&fl.i, "i", false, "игнорировать регистр")
	flag.BoolVar(&fl.v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fl.F, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&fl.n, "n", false, "напечатать номер строки")
	flag.Parse()
	templ := flag.Args()[0]
	fileName := flag.Args()[1]
	grep(fl, templ, fileName)
}

func grep(fl flags, templ, fileName string) {
	lines := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("i can't close this file")
		}
	}(file)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if fl.i {
			lines = append(lines, strings.ToLower(sc.Text()))
		} else {
			lines = append(lines, sc.Text())
		}
	}

	if fl.A > 0 && fl.B > 0 || fl.C > 0 {
		lines = grepC(int(math.Max(float64(fl.B), float64(fl.C))), int(math.Max(float64(fl.A), float64(fl.C))), fl.F, templ, &lines)
	} else if fl.A > 0 {
		lines = grepA(fl.A, fl.F, templ, &lines)
	} else if fl.B > 0 {
		lines = grepB(fl.B, fl.F, templ, &lines)
	} else if fl.n {
		lines = grepn(fl.F, templ, &lines)
	} else if fl.c {
		lines = grepc(fl.F, templ, &lines)
	} else if fl.v {
		lines = grepv(fl.F, templ, &lines)
	} else {
		lines = grepOrig(fl.F, templ, &lines)
	}

	writeFile, err := os.Create(fileName[0:len(fileName)-4] + "_sorted.txt")
	defer func(writeFile *os.File) {
		err := writeFile.Close()
		if err != nil {
			log.Println("i can't close this file")
		}
	}(writeFile)
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

func grepA(after int, f bool, templ string, lines *[]string) []string {
	answer := make([]string, 0)
	for i := 0; i < len(*lines); i++ {
		if f {
			if (*lines)[i] == templ {
				for j := i; j < len(*lines) && j < i+after; j++ {
					answer = append(answer, (*lines)[j])
				}
			}
		} else if strings.Contains((*lines)[i], templ) {
			for j := i; j < len(*lines) && j < i+after; j++ {
				answer = append(answer, (*lines)[j])
			}
		}
	}
	return answer
}

func grepB(before int, f bool, templ string, lines *[]string) []string {
	answer := make([]string, 0)
	for i := 0; i < len(*lines); i++ {
		if f {
			if (*lines)[i] == templ {
				for j := int(math.Max(0, float64(i-before))); j <= i; j++ {
					answer = append(answer, (*lines)[j])
				}
			}
		} else if strings.Contains((*lines)[i], templ) {
			for j := int(math.Max(0, float64(i-before))); j <= i; j++ {
				answer = append(answer, (*lines)[j])
			}
		}
	}
	return answer
}

func grepC(before, after int, f bool, templ string, lines *[]string) []string {
	answer := make([]string, 0)
	for i := 0; i < len(*lines); i++ {
		if f {
			if (*lines)[i] == templ {
				for j := int(math.Max(0, float64(i-before))); j < len(*lines) && j < i+after; j++ {
					answer = append(answer, (*lines)[j])
				}
			}
		} else if strings.Contains((*lines)[i], templ) {
			for j := int(math.Max(0, float64(i-before))); j < len(*lines) && j < i+after; j++ {
				answer = append(answer, (*lines)[j])
			}
		}
	}
	return answer
}

//количество строк
func grepc(f bool, templ string, lines *[]string) []string {
	qty := 0
	for _, l := range *lines {
		if f {
			if l == templ {
				qty++
			}
		} else if strings.Contains(l, templ) {
			qty++
		}
	}
	return []string{strconv.Itoa(qty)}
}

//вместо совпадения, исключать
func grepv(f bool, templ string, lines *[]string) []string {
	answer := make([]string, 0)
	for _, l := range *lines {
		if f {
			if l == templ {
				continue
			}
		} else if strings.Contains(l, templ) {
			continue
		}
		answer = append(answer, l)
	}
	return answer
}

//напечатать номер строки
func grepn(f bool, templ string, lines *[]string) []string {
	answer := make([]string, 0)
	for i, l := range *lines {
		if f {
			if l == templ {
				answer = append(answer, strconv.Itoa(i+1))
			}
		} else if strings.Contains(l, templ) {
			answer = append(answer, strconv.Itoa(i+1))
		}
	}
	return answer
}

func grepOrig(f bool, templ string, lines *[]string) []string {
	answer := make([]string, 0)
	for _, l := range *lines {
		if f {
			if l == templ {
				answer = append(answer, l)
			}
		} else if strings.Contains(l, templ) {
			answer = append(answer, l)
		}
	}
	return answer
}
