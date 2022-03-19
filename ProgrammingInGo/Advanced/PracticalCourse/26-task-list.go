package main

import (
    "fmt"
)

type taskList struct {
    tasks []*task
}

func (tl *taskList) add_to_list(t *task) {
    tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) delete_to_list(index int) {
    fmt.Printf("\nDelete by index %d\n", index)
    tl.tasks =append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl taskList) show_list() {
    fmt.Printf("Print Tasks\n")
    for _, value := range tl.tasks  {
        fmt.Printf("name: %s, description: %s, complete: %v \n", value.name, value.description, value.complete)
    }
}

func (tl taskList) show_list_complete() {
    fmt.Printf("\nPrint Complete Tasks\n")
    for _, value := range tl.tasks  {
        if value.complete {
            fmt.Printf("name: %s, description: %s, complete: %v \n", value.name, value.description, value.complete)
        }
    }
}

type task struct {
    name string
    description string
    complete bool
}

func (t *task) mark_complete() {
    t.complete = true
}

func (t *task) set_description(description string) {
    t.description = description
}

func add_task(name string, description string) *task {
    return &task {
        name: name,
        description: description,
    }
}

func main() {
    t1 := add_task( "Complete my course of Go", "Complete my course of go in Platzi in this week")
    t2 := add_task("Complete my course of Python", "Complete my course of python in Platzi in this week")
    t3 := add_task("Complete my course of Rust", "Complete my course of rust in Youtube of Cyan nyan in this week")

    list := taskList{
        tasks: []*task { t1, t2 },
    }
    list.add_to_list(t3)

    list.show_list()
    list.delete_to_list(0)
    list.show_list()
    list.show_list_complete()
    fmt.Printf("\nComplete task with index 0\n")
    list.tasks[0].mark_complete()
    list.show_list()
    list.show_list_complete()
}
