package store

// Store is not thread-safe.
type Store interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Len() int                                 // get store size
	ForEach(func(string, string) error) error // apply func on each each entry. If error occurs, inne loop breaks
}
