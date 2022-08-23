package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type flags struct {
	t      time.Duration
	ip     string
	domain string
	port   int
}

func main() {
	var fl flags
	flag.StringVar(&fl.ip, "i", "", "IP порта для подключения")
	flag.StringVar(&fl.domain, "d", "", "доменное имя для подключения")
	flag.IntVar(&fl.port, "p", 8080, "порт для подключения")
	flag.DurationVar(&fl.t, "t", 0, "таймаут на подключение к серверу")
	if fl.domain != "" && fl.ip != "" {
		log.Fatal("insert only one way to connect")
	}
	flag.Parse()
	telnet(fl)
}

func telnet(fl flags) {
	var conn net.Conn
	var err error
	if fl.ip != "" {
		if fl.t != 0 {
			conn, err = net.DialTimeout("tcp", fl.ip+":"+strconv.Itoa(fl.port), fl.t)
		}

	} else {
		if fl.t != 0 {
			conn, err = net.DialTimeout("tcp", fl.domain+":"+strconv.Itoa(fl.port), fl.t)
		}
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal)
	go handleConnection(fl, conn, shutdownSignal)
	go waitCtrlD(shutdownSignal)
	<-shutdownSignal

}

func waitCtrlD(shutdownSignal chan os.Signal) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyCtrlD {
			shutdownSignal <- syscall.SIGQUIT
			break
		}
	}
}

func handleConnection(fl flags, conn net.Conn, shutdownSignal chan os.Signal) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		n, err := conn.Write([]byte(scanner.Text()))
		if err != nil {
			fmt.Println(err)
			shutdownSignal <- syscall.SIGQUIT
			continue
		}
		if n == 0 {
			log.Println("got zero string from server")
			continue
		}
		buff := make([]byte, 1024)
		n, err3 := conn.Read(buff)
		if err3 != nil {
			log.Println(err3)
			shutdownSignal <- syscall.SIGQUIT
			return
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}
}
