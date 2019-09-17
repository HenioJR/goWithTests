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

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		errAdd := dictionary.Add("new word", "nova palavra")

		got, err := dictionary.Search("new word")
		want := "nova palavra"

		assertStrings(t, got, want)
		assertNoError(t, err)
		assertNoError(t, errAdd)
	})
	t.Run("add repeated word", func(t *testing.T) {
		dictionary := Dictionary{"repeated word": "palavra repetida"}
		errAdd := dictionary.Add("repeated word", "nova traducao")

		got, err := dictionary.Search("repeated word")
		want := "palavra repetida"

		assertStrings(t, got, want)
		assertNoError(t, err)
		assertError(t, errAdd, ErrRepeatedWord.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"apple": "banana"}
		dictionary.Update("apple", "maca")

		got, err := dictionary.Search("apple")
		want := "maca"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"apple": "banana"}
		dictionary.Update("car", "carro")

		got, err := dictionary.Search("car")
		want := "carro"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"car": "carro"}
	dictionary.Delete("car")

	got, err := dictionary.Search("car")
	want := ""

	assertStrings(t, got, want)
	assertError(t, err, ErrUnknownWord.Error())
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
