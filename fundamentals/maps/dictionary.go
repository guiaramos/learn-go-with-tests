package maps

// Dictionary is a map for storing words
type Dictionary map[string]string

// DictionaryErr is a mpa for storing dictionary errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound creates an error for not found words
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists creates an error for add existing words
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExists creates an error for update words that does not exists
	ErrWordDoesNotExists = DictionaryErr("could not update word because it does not exists")
)

// Search look up for values in dict
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add create a new word with definition
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update a word with new definition
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = def
	default:
		return err
	}

	return nil
}

// Delete a word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
