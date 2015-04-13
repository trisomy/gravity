package gravity

// Config .
type Config map[string]interface{}

// Task the project task object
type Task struct {
	next *Task  // The next node
	prev **Task // The prev node
	f    func() // The task execute routine
}

// NewTask create new task
func NewTask(prev **Task, f func()) *Task {
	return &Task{
		next: nil,
		prev: prev,
		f:    f,
	}
}
