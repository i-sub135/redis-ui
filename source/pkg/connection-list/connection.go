package connectionlist

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("connection not found")

type Connection struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Password string `json:"password,omitempty"`
	DB       int    `json:"db"`
}

// Store manages the connection list persisted as a JSON file.
// The file is created on first write if it doesn't exist.
type Store struct {
	mu   sync.RWMutex
	path string
}

func NewStore(path string) *Store {
	return &Store{path: path}
}

func (s *Store) Load() ([]Connection, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.readFile()
}

func (s *Store) Add(conn Connection) (Connection, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	conns, err := s.readFile()
	if err != nil {
		return Connection{}, err
	}
	conn.ID = uuid.New().String()
	conns = append(conns, conn)
	return conn, s.writeFile(conns)
}

func (s *Store) Update(id string, updated Connection) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	conns, err := s.readFile()
	if err != nil {
		return err
	}
	for i, c := range conns {
		if c.ID == id {
			updated.ID = id
			conns[i] = updated
			return s.writeFile(conns)
		}
	}
	return ErrNotFound
}

func (s *Store) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	conns, err := s.readFile()
	if err != nil {
		return err
	}
	out := conns[:0]
	found := false
	for _, c := range conns {
		if c.ID == id {
			found = true
			continue
		}
		out = append(out, c)
	}
	if !found {
		return ErrNotFound
	}
	return s.writeFile(out)
}

func (s *Store) GetByID(id string) (Connection, error) {
	conns, err := s.Load()
	if err != nil {
		return Connection{}, err
	}
	for _, c := range conns {
		if c.ID == id {
			return c, nil
		}
	}
	return Connection{}, ErrNotFound
}

func (s *Store) readFile() ([]Connection, error) {
	data, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return []Connection{}, nil
	}
	if err != nil {
		return nil, err
	}
	var conns []Connection
	if err := json.Unmarshal(data, &conns); err != nil {
		return nil, err
	}
	return conns, nil
}

func (s *Store) writeFile(conns []Connection) error {
	data, err := json.MarshalIndent(conns, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o644)
}
