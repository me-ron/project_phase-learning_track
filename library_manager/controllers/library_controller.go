package controllers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"errors"
	sconv "strconv"
	model "library/models"
	service "library/services"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func display(Books []model.Book){
	for _, book := range Books {
		fmt.Printf("%-20d %-30s %-30s %-30s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

func getStringInput(prompt string, regex string) (string, error) {
	fmt.Println(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input[:len(input)-1])
	matched, _ := regexp.MatchString(regex, input)
	if matched {
		return input, nil
	}
	return "", errors.New("invalid input")

}


func getIntInput(prompt string) (int, error) {
	fmt.Println(prompt)
	sInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input, err := sconv.Atoi(sInput[:len(sInput)-1])
	if err == nil {
		return input, nil
	}
		
	return 0, errors.New("invalid integer")
}

func ADDBOOK(lib *service.Library) {
	title, t_err := getStringInput("Title:", `^(?=.*\w)[\w\s]+$`)
	for t_err != nil {
		fmt.Println(t_err.Error())
		title, t_err = getStringInput("Title:", `^(?=.*\w)[\w\s]+$`)
		
	}

	author, a_err := getStringInput("Author:", `^(?=.*\w)[\w\s]+$`)
	for a_err != nil {
		fmt.Println(a_err.Error())
		author, a_err = getStringInput("Author:", `^(?=.*\w)[\w\s]+$`)
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
	for err != nil {
		fmt.Println(err.Error())
		id, err = getIntInput("ID:")
	}

	lib.RemoveBook(id)
	fmt.Println("Book removed successfully.")
}

func BORROW(lib *service.Library) {
	bookID, b_err := getIntInput("Book ID:")
	for b_err != nil {
		fmt.Println(b_err.Error())
		bookID, b_err = getIntInput("Book ID:")
	}

	memberID, m_err := getIntInput("Member ID:")
	for m_err != nil {
		fmt.Println(m_err.Error())
		memberID, m_err = getIntInput("Member ID:")
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
	for b_err != nil {
		fmt.Println(b_err.Error())
		bookID, b_err = getIntInput("Book ID:")
	}

	memberID, m_err := getIntInput("Member ID:")
	for m_err != nil {
		fmt.Println(m_err.Error())
		memberID, m_err = getIntInput("Member ID:")
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
		display(availableBooks)

	}
}

func LISTBORROWED(lib *service.Library, member model.Member) {
	memberID := member.ID
	
	borrowedBooks := lib.ListBorrowedBooks(memberID)
	if len(borrowedBooks) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed books:")
		display(borrowedBooks)
	}
}

func SIGNUP(lib *service.Library) model.Member{
	name, n_err := getStringInput("Name:", `^(?=.*\w)[\w\s]+$`)
	for n_err != nil {
		fmt.Println(n_err.Error())
		name, n_err = getStringInput("Name:", `^(?=.*\w)[\w\s]+$`)
	}

	member := model.Member{
		Name: name,
		ID:   lib.NextmemberID,
	}

	lib.AddMember(member)
	return member
}

func getmem()(int, error){
	ID, m_err := getIntInput("ID:")
	for m_err != nil{
		fmt.Println(m_err.Error())
		ID, m_err = getIntInput("ID:")
	
	}
	return ID, nil
}

func LOGIN(lib *service.Library)model.Member{
	ID, _ := getmem()
	member, ok := lib.Members[ID]
	for !ok{
		ID, _ = getmem()
		member, ok = lib.Members[ID]	
	}

	return *member
}
