package store

type inmemStore struct {
	m map[string]string // The key-value store for the system.
}

const defaultCap = 1024

func NewInmemStore() Store {

	return &inmemStore{
		m: make(map[string]string, defaultCap),
	}
}

func (s *inmemStore) Set(key, value string) error {
	s.m[key] = value

	return nil
}

func (s *inmemStore) Get(key string) (string, error) {
	value := s.m[key]

	return value, nil
}

func (s *inmemStore) Delete(key string) error {
	delete(s.m, key)

	return nil
}

func (s *inmemStore) Len() int {
	return len(s.m)
}

func (s *inmemStore) ForEach(fn func(string, string) error) error {
	for k, v := range s.m {
		if err := fn(k, v); err != nil {
			return err
		}
	}

	return nil
}
