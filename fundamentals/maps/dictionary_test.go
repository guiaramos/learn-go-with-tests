package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "this is just a test"

		dict.Add(word, def)

		assetDefinition(t, dict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assetDefinition(t, dict, word, def)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		newDef := "new definition"

		err := dict.Update(word, newDef)

		assertError(t, err, nil)
		assetDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, def)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	def := "this is just a test"
	dict := Dictionary{word: def}

	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}

}

func assetDefinition(t *testing.T, dict Dictionary, word, def string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, def)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}
