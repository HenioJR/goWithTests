package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"book": "livro"}

	got := dictionary.Search("book")
	want := "livro"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
