package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Int("timeout", 10, "Таймаут на подключение к серверу")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)

	if host == "" || port == "" {
		fmt.Println("Укажте хост и порт")
		return
	}
	path := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", path, time.Duration(*timeout)*time.Second)
	defer conn.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)

			for scanner.Scan() {
				msg := scanner.Text()
				if _, err = conn.Write([]byte(msg)); err != nil {
					fmt.Println(err)
				}

			}
		}
	}()

	go func() {
		tmp := make([]byte, 256)
		for {
			n, err := conn.Read(tmp)
			if err != nil {
				if err != io.EOF {
					fmt.Println("read error:", err)
				}
				break
			}

			io.WriteString(os.Stdout, fmt.Sprintf("%v \n", n))
		}
	}()

	<-quit
}
