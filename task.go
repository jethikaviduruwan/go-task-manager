package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	Priority  Priority  `json:"priority"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	DoneAt    *time.Time `json:"done_at,omitempty"`
}

type Store struct {
	Tasks  []*Task `json:"tasks"`
	NextID int     `json:"next_id"`
}

func dataPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".taskman", "tasks.json")
}

func loadStore() (*Store, error) {
	path := dataPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Store{NextID: 1}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading store: %w", err)
	}
	var s Store
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parsing store: %w", err)
	}
	return &s, nil
}

func (s *Store) save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataPath(), data, 0644)
}

func (s *Store) AddTask(title string, priority Priority, tags []string) *Task {
	t := &Task{
		ID:        s.NextID,
		Title:     title,
		Priority:  priority,
		Tags:      tags,
		CreatedAt: time.Now(),
	}
	s.NextID++
	s.Tasks = append(s.Tasks, t)
	return t
}

func (s *Store) FindByID(id int) *Task {
	for _, t := range s.Tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}

func (s *Store) DeleteByID(id int) bool {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Store) ClearDone() int {
	var remaining []*Task
	count := 0
	for _, t := range s.Tasks {
		if t.Done {
			count++
		} else {
			remaining = append(remaining, t)
		}
	}
	s.Tasks = remaining
	return count
}

func (s *Store) Pending() []*Task {
	var out []*Task
	for _, t := range s.Tasks {
		if !t.Done {
			out = append(out, t)
		}
	}
	return out
}
// v0-1
// v3-0
// v6-0
// v9-0
