package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

type Member struct {
	Name          string
	ID            string
	BorrowedBooks []string
}

// ---- Interface ----
type BookStore interface {
	AddBook(book *Book)
	GetBook(isbn string) (*Book, bool)
	GetAllBooks() map[string]*Book
}

// ---- In-memory implementation ----
type InMemoryBookStore struct {
	data map[string]*Book
}

func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		data: make(map[string]*Book),
	}
}

func (s *InMemoryBookStore) AddBook(book *Book) {
	s.data[book.ISBN] = book
}

func (s *InMemoryBookStore) GetBook(isbn string) (*Book, bool) {
	b, ok := s.data[isbn]
	return b, ok
}

func (s *InMemoryBookStore) GetAllBooks() map[string]*Book {
	return s.data
}

// ---- Library ----
type Library struct {
	store   BookStore
	Members map[string]*Member
}

func NewLibrary(store BookStore) *Library {
	return &Library{
		store:   store,
		Members: make(map[string]*Member),
	}
}

func (l *Library) AddBook(book *Book) {
	l.store.AddBook(book)
}

func (l *Library) AddMember(member *Member) {
	l.Members[member.ID] = member
}

func (l *Library) Borrow(memberId string, isbn string) error {
	member, ok := l.Members[memberId]
	if !ok {
		return fmt.Errorf("member not found")
	}

	book, ok := l.store.GetBook(isbn)
	if !ok {
		return fmt.Errorf("book not found")
	}

	if !book.Available {
		return fmt.Errorf("book not available")
	}

	book.Available = false
	member.BorrowedBooks = append(member.BorrowedBooks, isbn)

	return nil
}

func (l *Library) ReturnBook(memberId string, isbn string) error {
	member, ok := l.Members[memberId]
	if !ok {
		return fmt.Errorf("member not found")
	}

	book, ok := l.store.GetBook(isbn)
	if !ok {
		return fmt.Errorf("book not found")
	}

	book.Available = true

	for i, b := range member.BorrowedBooks {
		if b == isbn {
			member.BorrowedBooks = append(
				member.BorrowedBooks[:i],
				member.BorrowedBooks[i+1:]...,
			)
			break
		}
	}

	return nil
}

// ---- Search ----
func (l *Library) SearchByTitle(title string) []*Book {
	var result []*Book
	for _, b := range l.store.GetAllBooks() {
		if b.Title == title {
			result = append(result, b)
		}
	}
	return result
}

// ---- Utility ----
func (l *Library) ListAvailable() {
	fmt.Println("Available books:")
	for _, book := range l.store.GetAllBooks() {
		if book.Available {
			fmt.Printf("%s by %s (ISBN: %s)\n",
				book.Title, book.Author, book.ISBN)
		}
	}
}

// ---- Main ----
func main() {
	store := NewInMemoryBookStore()
	lib := NewLibrary(store)

	lib.AddBook(&Book{"Go basics", "Alice", "101", true})
	lib.AddBook(&Book{"DSA Deep Dive", "Bob", "102", true})

	lib.AddMember(&Member{"Avez", "M1", []string{}})

	_ = lib.Borrow("M1", "101")

	lib.ListAvailable()

	_ = lib.ReturnBook("M1", "101")

	fmt.Println("After Return")
	lib.ListAvailable()
}