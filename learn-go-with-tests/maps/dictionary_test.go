package maps

import "testing"

/* func TestSearch(t *testing.T) {
  dictionary := map[string]string{"test": "this is just a test"}

  word := "test"
  got := Search(dictionary, word)
  want := "this is just a test"

  if got != want {
    t.Errorf("expected %q want %q given %q", got, want, word)
  }
} */

/* func TestSearch(t *testing.T) {
  dictionary := map[string]string{"test": "this is just a test"}

  word := "test"
  got := Search(dictionary, word)
  want := "this is just a test"

  assertStrings(t, got, want)
}
*/

/*
func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  word := "test"
  got := dictionary.Search(word)
  want := "this is just a test"

  assertStrings(t, got, want)
}
*/

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if definition != got {
		t.Errorf("expected %q want %q", got, definition)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		word := "test"
		got, _ := dictionary.Search(word)
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "unknown"
		_, err := dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		got, err := dictionary.Search(word)

		assertDefinition(t, dictionary, word, definition)
		assertError(t, err, nil)
		assertStrings(t, got, definition)
	})

  // [02]
/* 	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionaty := Dictionary{word: definition}
		err := dictionaty.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	}) */

  t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dictionary := Dictionary{word: definition}
    newDefinition := "new definition"

    err := dictionary.Update(word, newDefinition)

    assertError(t, err, nil)
    assertDefinition(t, dictionary, word, newDefinition)
  })

  t.Run("new word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dictionary := Dictionary{}

    err := dictionary.Update(word, definition)

    assertError(t, err, ErrWordDoesNotExist)
  })

  t.Run("delete", func(t *testing.T) {
    word := "test"
    dictionary := Dictionary{word: "test definition"}

    dictionary.Delete(word)

    _, err := dictionary.Search(word)
    if err != ErrNotFound {
      t.Errorf("expected %q to be delete", word)
    }
  })
}
