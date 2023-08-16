package impl

import (
	"sort"
	"strconv"
	"strings"
)

type Bookstore struct {
	database []string
}

func (b *Bookstore) FindBooks(req string, books *string) error {
	keywords := strings.Split(req, " ")
	rep := booksWithKeyWords(b.database, keywords)
	ans := ""

	for i := 0; i < len(rep); i++ {
		ans += strconv.Itoa(rep[i].keywords) + ": " + rep[i].bookName
		if i < len(rep)-1 {
			ans += ", "
		}
	}

	// envia a mensagem processada de volta ao cliente
	*books = ans
	return nil
}

func NewBookstore() *Bookstore {
	return &Bookstore{
		database: []string{"Harry Potter e a Pedra Filosofal", "Harry Potter e a Camara Secreta", "Harry Potter e o Prisioneiro de Azkaban", "Harry Potter e o Calice de Fogo", "Harry Potter e a Ordem da Fenix", "Harry Potter e o Enigma do Principe", "Harry Potter e as Reliquias da Morte"},
	}
}

type Book struct {
	keywords int
	bookName string
}

func sortBooks(books []Book) []Book {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].keywords > books[j].keywords
	})
	return books
}

func findKeyWords(book string, keywords []string) int {
	count := 0
	for _, keyword := range keywords {
		keyword = strings.Trim(keyword, "\n")
		if strings.Contains(book, keyword) {
			count = count + 1
		}
	}
	return count
}

func booksWithKeyWords(books []string, keywords []string) []Book {
	var response []Book
	for _, book := range books {
		counter := findKeyWords(book, keywords)
		if counter > 0 {
			response = append(response, Book{counter, book})
		}
	}
	sortBooks(response)
	return response
}
