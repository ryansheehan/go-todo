package todo

type Todo struct {
	Todo     string
	complete bool
}

type Todos []Todo

func NewTodo() *Todo {
	return &Todo{
		Todo:     "",
		complete: false,
	}
}

func (t *Todo) MarkComplete() {
	t.complete = true
}

func (t *Todo) IsComplete() bool {
	return t.complete
}
