package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"book": "livro"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("book")
		want := "livro"
		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("strange word")
		want := ""
		assertStrings(t, got, want)
		assertError(t, err, ErrUnknownWord.Error())

	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()

	if got == nil {
		t.Error("wanted an error bug didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Error("got an error but didn't want one")
	}
}
