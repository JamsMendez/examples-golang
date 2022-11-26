package maps

/* func Search(dictionary map[string]string, word string) string {
  return dictionary[word]
} */

type Dictionary map[string]string

/* func (d Dictionary) Search(word string) string {
  // [Alternative]
	// d `Dictionary
	// return (*d)[word]
	return d[word]
} */

// [02]
/* var ErrNotFound = errors.New("could not find the word you where looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")
var ErrWordDoesNotExist = errors.New("cannot update word because it does not exist") */


var ErrNotFound = DictionaryErr("could not find the word you where looking for")
var ErrWordExists = DictionaryErr("cannot add word because it already exists")
var ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (value string, err error) {
	value, ok := d[word]
	if !ok {
		err = ErrNotFound
	}

	return value, err
}

// [01]
/* func (d Dictionary) Add(word, definition string) (err error) {
	// So when you pass a map to a function/method,
	// you are indeed copying it, but just the pointer part,
	// not the underlying data structure that contains the data.
	//d[word] = definition

	if d == nil {
	  err = errors.New("")
	  return err
	}

  d[word] = definition

  return
} */

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
  delete(d, word)
}
