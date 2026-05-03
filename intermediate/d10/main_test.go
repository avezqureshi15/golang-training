package main

import (
	"strconv"
	"testing"
)

// ---- Mock Store ----
type MockBookStore struct {
	data map[string]*Book
}

func NewMockBookStore() *MockBookStore {
	return &MockBookStore{
		data: make(map[string]*Book),
	}
}

func (m *MockBookStore) AddBook(book *Book) {
	m.data[book.ISBN] = book
}

func (m *MockBookStore) GetBook(isbn string) (*Book, bool) {
	b, ok := m.data[isbn]
	return b, ok
}

func (m *MockBookStore) GetAllBooks() map[string]*Book {
	return m.data
}

// ---- Setup ----
func setupLibrary() *Library {
	store := NewMockBookStore()
	lib := NewLibrary(store)

	lib.AddBook(&Book{"Go", "Alice", "1", true})
	lib.AddBook(&Book{"DSA", "Bob", "2", true})

	lib.AddMember(&Member{"Avez", "M1", []string{}})

	return lib
}

// ---- Tests ----
func TestBorrow(t *testing.T) {
	tests := []struct {
		name      string
		memberID  string
		isbn      string
		setup     func(l *Library)
		wantError bool
	}{
		{"success", "M1", "1", func(l *Library) {}, false},
		{"member not found", "M2", "1", func(l *Library) {}, true},
		{"book not found", "M1", "999", func(l *Library) {}, true},
		{
			"book not available",
			"M1",
			"1",
			func(l *Library) {
				book, _ := l.store.GetBook("1")
				book.Available = false
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lib := setupLibrary()
			tt.setup(lib)

			err := lib.Borrow(tt.memberID, tt.isbn)

			if (err != nil) != tt.wantError {
				t.Errorf("expected error %v, got %v", tt.wantError, err)
			}
		})
	}
}

func TestReturnBook(t *testing.T) {
	tests := []struct {
		name      string
		memberID  string
		isbn      string
		setup     func(l *Library)
		wantError bool
	}{
		{
			"success",
			"M1",
			"1",
			func(l *Library) {
				_ = l.Borrow("M1", "1")
			},
			false,
		},
		{"member not found", "M2", "1", func(l *Library) {}, true},
		{"book not found", "M1", "999", func(l *Library) {}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lib := setupLibrary()
			tt.setup(lib)

			err := lib.ReturnBook(tt.memberID, tt.isbn)

			if (err != nil) != tt.wantError {
				t.Errorf("expected error %v, got %v", tt.wantError, err)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	lib := setupLibrary()

	res := lib.SearchByTitle("Go")

	if len(res) != 1 {
		t.Errorf("expected 1 result, got %d", len(res))
	}
}

func TestAddBookAndMember(t *testing.T) {
	store := NewMockBookStore()
	lib := NewLibrary(store)

	lib.AddBook(&Book{"Test", "A", "123", true})
	lib.AddMember(&Member{"Avez", "M1", nil})

	if _, ok := store.GetBook("123"); !ok {
		t.Errorf("book not added")
	}

	if _, ok := lib.Members["M1"]; !ok {
		t.Errorf("member not added")
	}
}

// ---- Benchmark ----
func BenchmarkSearch(b *testing.B) {
	store := NewMockBookStore()
	lib := NewLibrary(store)

	for i := 0; i < 10000; i++ {
		lib.AddBook(&Book{
			Title:     "Book",
			Author:    "Author",
			ISBN:      strconv.Itoa(i),
			Available: true,
		})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lib.SearchByTitle("Book")
	}
}