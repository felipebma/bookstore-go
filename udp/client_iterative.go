package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("UDP client exiting...")
			break
		}

		_, err = conn.Write([]byte(text))
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

		fmt.Println(string(received))
	}
}
