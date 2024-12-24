package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]Book)}
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists {
		if book.IsBorrowed {
			fmt.Println("The book is already borrowed.")
			return
		}
		book.IsBorrowed = true
		l.Books[id] = book
		fmt.Println("Book borrowed successfully!")
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists {
		if !book.IsBorrowed {
			fmt.Println("The book is already returned.")
			return
		}
		book.IsBorrowed = false
		l.Books[id] = book
		fmt.Println("Book returned successfully!")
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	lib := NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Add a book")
		fmt.Println("2. Borrow a book")
		fmt.Println("3. Return a book")
		fmt.Println("4. List available books")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter book author: ")
			scanner.Scan()
			author := scanner.Text()

			lib.AddBook(Book{ID: id, Title: strings.TrimSpace(title), Author: strings.TrimSpace(author)})
			fmt.Println("Book added successfully!")
		case 2:
			var id string
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scan(&id)
			lib.BorrowBook(id)
		case 3:
			var id string
			fmt.Print("Enter book ID to return: ")
			fmt.Scan(&id)
			lib.ReturnBook(id)
		case 4:
			lib.ListBooks()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
