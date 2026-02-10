package todo

import "sync"

// Item represents a todo item.
type Item struct {
	ID   int
	Text string
	Done bool
}

// Todo manages todo items and handles HTTP requests.
type Todo struct {
	mu    sync.RWMutex
	items []Item
}

// Add creates a new todo and returns the total count.
func (t *Todo) Add(text string) int {
	t.mu.Lock()
	defer t.mu.Unlock()

	item := Item{
		ID:   len(t.items) + 1,
		Text: text,
		Done: false,
	}
	t.items = append(t.items, item)
	return len(t.items)
}

// All returns a copy of all todos.
func (t *Todo) All() []Item {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return append([]Item(nil), t.items...)
}
