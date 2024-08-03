package services

import (
	"errors"
	"fmt"
	model "library/models"
)

type Library struct{
	Books map[int]*model.Book
	Members map[int]*model.Member
	NextbookID int
	NextmemberID int
}

type LibraryManager interface{
	AddBook(book model.Book)
	AddMember(member model.Member)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []model.Book
	ListBorrowedBooks(memberID int) []model.Book
}

func (lib *Library) AddBook(book model.Book){
	lib.Books[book.ID] = &book
	lib.NextbookID ++
}

func (lib *Library) AddMember(member model.Member){
	lib.Members[member.ID] = &member
	lib.NextmemberID ++
}

func(lib *Library) RemoveBook(bookID int){
	delete(lib.Books, bookID)
}

func(lib *Library) BorrowBook(bookID int, memberID int) error{
	book, ok := lib.Books[bookID]
	if ok{
		if book.Status == "Borrowed"{
			return errors.New("the book isn't available")
		} 
		lib.Books[bookID].Status = "Borrowed"
		lib.Members[memberID].BorrowedBooks = append(lib.Members[memberID].BorrowedBooks, *lib.Books[bookID])

		return nil
		}
	return errors.New("book doesn't exist")
	
}

func(lib *Library) ReturnBook(bookID int, memberID int) error{
	_, ok := lib.Books[bookID];
	if ok{
		lib.Books[bookID].Status = "Available"
		member_books := lib.Members[memberID].BorrowedBooks
		for i := range len(member_books){
			book := member_books[i]
			if book.ID == bookID{
				member_books = append(member_books[:i], member_books[i + 1:]...)
				break
			}
		}

		fmt.Println(lib.Books[bookID],"000")
		lib.Members[memberID].BorrowedBooks = member_books

		return nil                       
	}else{
		return errors.New("book doesn't exist")
	}

}

func (lib *Library) ListAvailableBooks() []model.Book{
	var avail_books []model.Book
	for _, book := range lib.Books{
		avail_books = append(avail_books, *book)
	}

	return avail_books
}

func (lib *Library) ListBorrowedBooks(memberID int) []model.Book{
	return lib.Members[memberID].BorrowedBooks
}