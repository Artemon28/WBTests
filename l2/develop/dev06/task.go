package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
first	second	third
hello wo,rld he,re
egefgedf	fesfe	fewfwefwef
waefewfew	regesrfrfewf	wefewfewfwe	ewfgewfgewf	fgg
*/

type flags struct {
	f int
	d string
	s bool
}

func main() {
	var fl flags
	flag.IntVar(&fl.f, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&fl.d, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&fl.s, "s", false, "только строки с разделителем")
	flag.Parse()
	cut(fl)
}

//read until 0 goes
func cut(fl flags) {
	lines := make([]string, 0)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if sc.Text() == "0" {
			break
		}
		lines = append(lines, sc.Text())
	}

	if fl.f <= 0 && !fl.s {
		log.Fatal("no options")
	}

	doCut(fl, &lines)
}

func doCut(fl flags, lines *[]string) {
	for _, v := range *lines {
		words := strings.Split(v, fl.d)
		if fl.f != 0 {
			if len(words) == 1 {
				if fl.s {
					continue
				}
				fmt.Println(words[0])
				continue
			}
			if len(words) >= fl.f {
				fmt.Println(words[fl.f-1])
			} else {
				fmt.Println("")
			}
		} else {
			if strings.Contains(v, fl.d) {
				fmt.Println(v)
			}
		}
	}
}
