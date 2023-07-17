package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Book struct {
	keywords int
	bookName string
}

func sortBooks(books []Book) []Book {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].keywords > books[j].keywords
	})
	return books
}

func findKeyWords(book string, keywords []string) int {
	count := 0
	for _, keyword := range keywords {
		keyword = strings.Trim(keyword, "\n")
		if strings.Contains(book, keyword) {
			count = count + 1
		}
	}
	return count
}

func booksWithKeyWords(books [7]string, keywords []string) []Book {
	var response []Book
	for _, book := range books {
		counter := findKeyWords(book, keywords)
		if counter > 0 {
			response = append(response, Book{counter, book})
		}
	}
	sortBooks(response)
	return response
}

func main() {
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// ouvindo na porta 8081 via protocolo tcp
	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	fmt.Println("Aguardando conexões...")

	for {
		// aceitando conexões
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		fmt.Println("Cliente conectado")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	fmt.Println("Conexão aceita...")

	for {
		req, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		books := [7]string{"Harry Potter e a Pedra Filosofal", "Harry Potter e a Camara Secreta", "Harry Potter e o Prisioneiro de Azkaban", "Harry Potter e o Calice de Fogo", "Harry Potter e a Ordem da Fenix", "Harry Potter e o Enigma do Principe", "Harry Potter e as Reliquias da Morte"}

		keywords := strings.Split(req, " ")
		rep := booksWithKeyWords(books, keywords)
		ans := ""

		for i := 0; i < len(rep); i++ {
			ans += strconv.Itoa(rep[i].keywords) + ": " + rep[i].bookName
			if i < len(rep)-1 {
				ans += ", "
			}
		}

		// envia a mensagem processada de volta ao cliente
		_, err = conn.Write([]byte(ans + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
