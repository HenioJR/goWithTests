package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"book": "livro"}

	got := Search(dictionary, "book")
	want := "livro"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
