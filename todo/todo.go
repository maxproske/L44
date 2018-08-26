package todo

import (
	"errors"
	"sync"

	"github.com/rs/xid"
)

// Define multiple globally available variables at once
var (
	list []Todo       // Holds all to-do items
	mtx  sync.RWMutex // Mutex allows safe access/manipulation across different goroutines
	once sync.Once    // Assure a specific operation will only run once
)

// Golang runs the init function whenever the package todo is loaded
func init() {
	// Wrap in once.Do to avoid resetting the array on runtime
	once.Do(initializeList)
}

// Initializes the array of to-do items
func initializeList() {
	list = []Todo{}
}

// Todo data structure for a task with a description of what to do
type Todo struct {
	ID       string `json:"id"` // Map public property into its lowercase JSON equivalent
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

// Get retrieves all elements from the current todo list
func Get() []Todo {
	return list
}

// Add will add a new to-do to the global list, based on a user input message
func Add(message string) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	return t.ID
}

// Delete will remove a Todo from the global list
func Delete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	return nil
}

// Complete will set the complete boolean to true, marking a todo as
// completed
func Complete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	return nil
}

// newTodo takes a message and returns a new instance of Todo
func newTodo(msg string) Todo {
	return Todo{
		ID:       xid.New().String(), // 3rd party package for generating UUIDs
		Message:  msg,
		Complete: false,
	}
}

func findTodoLocation(id string) (int, error) {
	mtx.RLock() // Will only be reading from the list and not writing to it
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find todo based on id")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	// Contains all elements from the previous list up to a given location, appended with all elements after (but not including) the same given location
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func setTodoCompleteByLocation(location int) {
	mtx.Lock()
	list[location].Complete = true
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}
