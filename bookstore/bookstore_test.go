package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickins",
		Copies: 99,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {
			ID:    1,
			Title: "For the Love of Go",
		},
		2: {
			ID:    2,
			Title: "The Power of Go: Tools",
		},
	}

	want := []bookstore.Book{
		{ID: 1,
			Title: "For the Love of Go",
		},
		{ID: 2,
			Title: "The Power of Go: Tools",
		},
	}

	got := bookstore.GetAllBooks(catalog)

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got) {
		t.Errorf("Test failed catalog is different: %v", cmp.Diff(want, got))
	}

}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {
			ID:    1,
			Title: "For the Love of Go",
		},
		2: {
			ID:    2,
			Title: "The Power of Go: Tools",
		},
	}

	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go: Tools",
	}

	got, err := bookstore.GetBook(catalog, 2)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("Want error for non-existant ID, got nil")
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickins",
		Copies: 2,
	}

	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatalf("Error trying to buy book %s: %v", b.Title, err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("Wanted %d copies of %s after buying 1 copy, got %d", want, b.Title, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickins",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Wanted error as zero copies should not happen")
	}
}

func TestNetPricePounds(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:           "Nicholas Chuckleby",
		Author:          "Charles Dickins",
		Copies:          2,
		PricePence:      999,
		DiscountPercent: 10,
		SpecialOffer:    true,
	}

	want := 900

	got := b.GetNetPriceInPence()

	if want != got {
		t.Errorf("Wanted the price to be %dp after the discount of %d%% from the original price of %d, got %d", want, b.DiscountPercent, b.PricePence, got)
	}

}
