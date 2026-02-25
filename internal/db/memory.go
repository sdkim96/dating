package db

import "sync"

type Entry struct {
	Key   string
	Value string
}
type MemoryStore struct {
	entries []Entry
	mu      sync.RWMutex
}

func (s *MemoryStore) ReadAll() []Entry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return append([]Entry(nil), s.entries...)
}

func (s *MemoryStore) Read(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, entry := range s.entries {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return "", false
}

func (s *MemoryStore) Write(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, entry := range s.entries {
		if entry.Key == key {
			s.entries[i].Value = value
			return
		}
	}
	s.entries = append(s.entries, Entry{Key: key, Value: value})
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		entries: []Entry{},
	}
}

var Store = NewMemoryStore()
