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

	for i := 0; i < 10000; i++ {
		start := time.Now()
		// envia mensagem para o servidor
		req := "Harry Potter"
		_, err = fmt.Fprintf(conn, req+"\n")
		if err != nil {
			fmt.Println(err)
			break
		}

		// recebe resposta do servidor
		result, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		end := time.Now()
		times = append(times, end.Sub(start).Nanoseconds())
		fmt.Print(result)
	}
	fmt.Fprintf(os.Stderr, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(times)), ","), "[]"))
}
