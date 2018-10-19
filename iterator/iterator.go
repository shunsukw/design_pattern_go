package main

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() *Book
}

type Book struct {
	Name string
}

func (book *Book) getName() string {
	return book.Name
}

//  Aggregate インターフェイスを実装する
type BookShelf struct {
	Books []*Book
	Last  int
}

func NewBookShelf() *BookShelf {
	return &BookShelf{
		Books: []*Book{},
		Last:  0,
	}
}

func (bookShelf *BookShelf) getBookAt(index int) *Book {
	return bookShelf.Books[index]
}

func (bookShelf *BookShelf) appendBook(book *Book) {
	bookShelf.Books = append(bookShelf.Books, book)
	bookShelf.Last++
}

func (bookShelf *BookShelf) getLength() int {
	return bookShelf.Last
}

func (bookShelf *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{
		BookShelf: bookShelf,
		Index:     0,
	}
}

// Iterator インターフェイスを実装する
type BookShelfIterator struct {
	BookShelf *BookShelf
	Index     int
}

func (bookShelfIterator *BookShelfIterator) HasNext() bool {
	if bookShelfIterator.Index < bookShelfIterator.BookShelf.getLength() {
		return true
	} else {
		return false
	}
}

func (bookShelfIterator *BookShelfIterator) Next() *Book {
	book := bookShelfIterator.BookShelf.getBookAt(bookShelfIterator.Index)
	bookShelfIterator.Index++
	return book
}

func main() {
	bookShelf := NewBookShelf()
	bookShelf.appendBook(&Book{
		Name: "本1",
	})
	bookShelf.appendBook(&Book{
		Name: "本2",
	})

	iterator := bookShelf.Iterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next().getName())
	}
}
