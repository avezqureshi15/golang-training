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

type Library struct {
	Books   map[string]*Book
	Members map[string]*Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[string]*Book),
		Members: make(map[string]*Member),
	}
}

func (l *Library) AddBook(book *Book) {
	l.Books[book.ISBN] = book
}
func (l *Library) AddMember(member *Member) {
	l.Members[member.ID] = member
}
func (l *Library) Borrow(memberId string, isbn string) error {
	member, ok := l.Members[memberId]
	if !ok {
		return fmt.Errorf("member not found")
	}

	book , ok := l.Books[isbn]
	if !ok {
		return fmt.Errorf("book not found")
	}

	if !book.Available{
		return fmt.Errorf("book not available")
	}

	book.Available = false 
	member.BorrowedBooks = append(member.BorrowedBooks,isbn)
	return nil
}

func (l *Library) ReturnBook(memberId string, isbn string) error {
	member, ok := l.Members[memberId]
	if !ok {
		return fmt.Errorf("member not found")
	}
	book , ok := l.Books[isbn]
	if !ok {
		return fmt.Errorf("book not found")
	}
	
	book.Available = true
	
	for i, b := range member.BorrowedBooks{
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
func (l *Library) ListAvailable(){
	fmt.Println("Available books: ")
	for _,book := range l.Books{
		if book.Available{
		fmt.Printf("%s by %s (ISBN: %s)\n",
				book.Title, book.Author, book.ISBN)
		}
	}
}
func main() {
 lib := NewLibrary()

 lib.AddBook(&Book{"Go basics","Alice","101",true})
 lib.AddBook(&Book{"DSA Deep Dive","Bob","102",true})

 lib.AddMember(&Member{"Avez","M1",[]string{}})

 err := lib.Borrow("M1","101")
 if err != nil{
	fmt.Println(err)
 }

 lib.ListAvailable()

 lib.ReturnBook("M1","101")

 fmt.Println("After Return")
 lib.ListAvailable()
}