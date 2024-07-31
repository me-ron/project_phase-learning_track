package controllers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	sconv "strconv"
	model "library/models"
	service "library/services"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// Helper function to get string input and validate with regex
func getStringInput(prompt string, regex string) (string, error) {
	fmt.Println(prompt)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		input = input[:len(input)-1]
		matched, _ := regexp.MatchString(regex, input)
		if matched {
			return input, nil
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

// Helper function to get integer input
func getIntInput(prompt string) (int, error) {
	fmt.Println(prompt)
	for {
		sInput, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		input, err := sconv.Atoi(sInput[:len(sInput)-1]) // Remove the newline character
		if err == nil {
			return input, nil
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
}

func ADD(lib *service.Library) {
	title, t_err := getStringInput("Title:", `^[\w\s]+$`)
	if t_err != nil {
		fmt.Println(t_err.Error())
		return
	}

	author, a_err := getStringInput("Author:", `^[\w\s]+$`)
	if a_err != nil {
		fmt.Println(a_err.Error())
		return
	}

	book := model.Book{
		Title:  title,
		Author: author,
		Status: "Available",
		ID:     lib.NextbookID,
	}

	lib.AddBook(book)
	fmt.Println("Book added successfully.")
}

func REMOVE(lib *service.Library) {
	id, err := getIntInput("ID:")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lib.RemoveBook(id)
	fmt.Println("Book removed successfully.")
}

func BORROW(lib *service.Library) {
	bookID, b_err := getIntInput("Book ID:")
	if b_err != nil {
		fmt.Println(b_err.Error())
		return
	}

	memberID, m_err := getIntInput("Member ID:")
	if m_err != nil {
		fmt.Println(m_err.Error())
		return
	}

	err := lib.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func RETURN(lib *service.Library) {
	bookID, b_err := getIntInput("Book ID:")
	if b_err != nil {
		fmt.Println(b_err.Error())
		return
	}

	memberID, m_err := getIntInput("Member ID:")
	if m_err != nil {
		fmt.Println(m_err.Error())
		return
	}

	err := lib.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func LISTAVAILABLE(lib *service.Library) {
	availableBooks := lib.ListAvailableBooks()
	if len(availableBooks) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("Available books:")
		for _, book := range availableBooks {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func LISTBORROWED(lib *service.Library) {
	memberID, m_err := getIntInput("Member ID:")
	if m_err != nil {
		fmt.Println(m_err.Error())
		return
	}

	borrowedBooks := lib.ListBorrowedBooks(memberID)
	if len(borrowedBooks) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed books:")
		for _, book := range borrowedBooks {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func ADDMEMBER(lib *service.Library) {
	name, n_err := getStringInput("Name:", `^[\w\s]+$`)
	if n_err != nil {
		fmt.Println(n_err.Error())
		return
	}

	member := model.Member{
		Name: name,
		ID:   lib.NextmemberID,
	}

	lib.AddMember(member)
	fmt.Println("Member added successfully.")
}
