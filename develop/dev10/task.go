package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type ConnArgs struct {
	host    string        // Host address of the server
	port    string        // Port number for the connection
	timeout time.Duration // Timeout duration for the connection
}

var (
	args ConnArgs
	conn net.Conn
	err  error
)

func main() {
	// parsing command-line arguments
	flag.DurationVar(&args.timeout, "timeout", 10, "Timeout for connection")
	flag.Parse()

	args.host = flag.Arg(0)
	args.port = flag.Arg(1)

	// try connection and handle error
	if args.timeout == 10*time.Nanosecond {
		conn, err = net.Dial("tcp", (args.host + ":" + args.port))
	} else {
		conn, err = net.DialTimeout("tcp", (args.host + ":" + args.port), args.timeout)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to %s:%s: %s\n", args.host, args.port, err)
		os.Exit(1)
	}
	defer conn.Close()

	// setting up a signal handler for graceful termination (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		<-sigCh
		fmt.Println("Closing connection...")
		conn.Close()
		os.Exit(0)
	}()

	for {
		var input string

		// reading user input
		fmt.Print("Enter message: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// sending user input to the server
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// reading and displaying the server's response
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("Error with reading answer: %v \n", err)
			return
		}
		result := string(buff[:n])
		fmt.Printf("Got answer: %s \n", result)

		time.Sleep(1 * time.Second)
	}
}
