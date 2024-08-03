# Library API Documentation

The Library API provides a set of methods to manage a library, including adding and removing books, adding members, borrowing and returning books, and listing available and borrowed books.

## Data Structures

### `model.Book`
- `ID`: unique identifier for the book
- `Title`: title of the book
- `Author`: author of the book
- `Status`: status of the book (either "Available" or "Borrowed")

### `model.Member`
- `ID`: unique identifier for the member
- `Name`: name of the member
- `BorrowedBooks`: list of books currently borrowed by the member

### `Library`
- `Books`: a map of book IDs to book pointers
- `Members`: a map of member IDs to member pointers
- `NextbookID`: the next available book ID
- `NextmemberID`: the next available member ID

### `LibraryManager` Interface
- `AddBook(book model.Book)`: adds a new book to the library
- `AddMember(member model.Member)`: adds a new member to the library
- `RemoveBook(bookID int)`: removes a book from the library
- `BorrowBook(bookID int, memberID int) error`: borrows a book for a member
- `ReturnBook(bookID int, memberID int) error`: returns a book for a member
- `ListAvailableBooks() []model.Book`: returns a list of all available books
- `ListBorrowedBooks(memberID int) []model.Book`: returns a list of all books borrowed by a member

## API Methods

### `AddBook(book model.Book)`
Adds a new book to the library.

### `AddMember(member model.Member)`
Adds a new member to the library.

### `RemoveBook(bookID int)`
Removes a book from the library.

### `BorrowBook(bookID int, memberID int) error`
Borrows a book for a member. Returns an error if the book is not available or does not exist.

### `ReturnBook(bookID int, memberID int) error`
Returns a book for a member. Returns an error if the book does not exist.

### `ListAvailableBooks() []model.Book`
Returns a list of all available books in the library.

### `ListBorrowedBooks(memberID int) []model.Book`
Returns a list of all books currently borrowed by a member.