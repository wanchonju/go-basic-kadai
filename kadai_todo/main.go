package main

import "fmt"

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *Todo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.ID)
	t.Completed = true
}

func printTodos(todos []Todo) {
	for _, todo := range todos {
		done := "未完了"
		if todo.Completed {
			done = "完　了"
		}
		fmt.Printf("[%s] (ID: %d) %s\n", done, todo.ID, todo.Title)
	}
}

func main() {

	// 1. []ToDo型のtodosスライスを作成
	todos := []Todo{
		{ID: 1, Title: "学習計画", Completed: false},
		{ID: 2, Title: "環境構築", Completed: false},
		{ID: 3, Title: "基礎文法", Completed: false},
	}

	// printTodos(todos)

	todos[0].Complete()
	todos[1].Complete()

	printTodos(todos)
}
