package main

import (
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
)

const (
	PORT = ":8082"
	TYPE = "udp"
)

type Book struct {
	keywords int
	bookName string
}

func main() {
	udpServer, err := net.ListenPacket(TYPE, PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		go handleConnection(udpServer, addr, buf[:n])
	}
}

func handleConnection(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	books := [7]string{"Harry Potter e a Pedra Filosofal", "Harry Potter e a Camara Secreta",
		"Harry Potter e o Prisioneiro de Azkaban", "Harry Potter e o Calice de Fogo",
		"Harry Potter e a Ordem da Fenix", "Harry Potter e o Enigma do Principe",
		"Harry Potter e as Reliquias da Morte"}

	clientRequest := strings.Trim(string(buf), "\n")
	keywords := strings.Split(clientRequest, " ")
	rep := booksWithKeyWords(books, keywords)

	ans := ""

	for i := 0; i < len(rep); i++ {
		ans += strconv.Itoa(rep[i].keywords) + ": " + rep[i].bookName
		if i < len(rep)-1 {
			ans += ", "
		}
	}

	udpServer.WriteTo([]byte(ans+"\n"), addr)
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
		if strings.Index(book, keyword) > -1 {
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
