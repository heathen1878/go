package bookstore

import (
	"errors"
	"fmt"
)

// Book represents info about a book.
type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PricePence      int
	DiscountPercent int
	SpecialOffer    bool
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]

	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}

	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (b Book) GetNetPriceInPence() int {

	discount := b.PricePence * b.DiscountPercent / 100

	netprice := b.PricePence - discount

	return netprice

}
