package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var times []int64

	// conectando na porta 8081 via protocolo tcp/ip na máquina local
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			break
		}

		start := time.Now()
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		end := time.Now()
		times = append(times, end.Sub(start).Nanoseconds())

	}
	fmt.Fprintf(os.Stderr, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(times)), ","), "[]"))
}
