package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	udpServer, err := net.ResolveUDPAddr("udp", ":8082")

	if err != nil {
		fmt.Println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		fmt.Println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	var times []int64

	for i := 0; i < 10000; i++ {

		start := time.Now()
		_, err = conn.Write([]byte("Harry Potter"))
		if err != nil {
			fmt.Println("Write data failed:", err.Error())
			os.Exit(1)
		}

		// buffer to get data
		received := make([]byte, 1024)
		_, err = conn.Read(received)
		if err != nil {
			fmt.Println("Read data failed:", err.Error())
			os.Exit(1)
		}
		end := time.Now()

		times = append(times, end.Sub(start).Nanoseconds())

		fmt.Println(string(received))
	}
	fmt.Fprintf(os.Stderr, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(times)), ","), "[]"))
}
