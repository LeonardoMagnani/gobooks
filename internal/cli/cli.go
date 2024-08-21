package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: books <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}

		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book id> <book id> <book id> ...")
			return
		}

		bookIds := os.Args[2:]

		cli.simulateReading(bookIds)
	}
}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)

	if err != nil {
		fmt.Println("error searching books")
		return
	}

	if len(books) == 0 {
		fmt.Println("no books found")
		return
	}

	fmt.Printf("%d books found\n", len(books))

	for _, book := range books {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Gênero: %s\n",
			book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReading(bookIdsStr []string) {
	var bookIds []int

	for _, idStr := range bookIdsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid book Id: ", idStr)
			continue
		}
		bookIds = append(bookIds, id)
	}

	responses := cli.service.SimulateMultipleReading(bookIds, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}
