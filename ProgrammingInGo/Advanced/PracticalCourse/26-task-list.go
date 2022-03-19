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
    tl.tasks =append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl taskList) show_list() {
    for _, value := range tl.tasks  {
        fmt.Printf("name: %s, description: %s, complete: %v \n", value.name, value.description, value.complete)
    }
}

func (tl taskList) show_list_complete() {
    fmt.Printf("Print Complete Tasks\n")
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

func add_task(name string, description string) *task {
    return &task {
        name: name,
        description: description,
    }
}

func main() {
    t1 := add_task("Complete my course of Go", "Complete my course of go in Platzi in this week")
    t2 := add_task("Complete my course of Python", "Complete my course of python in Platzi in this week")
    t3 := add_task("Complete my course of Rust", "Complete my course of rust in Youtube of Cyan nyan in this week")

    t4 := add_task("Complete my course of Java", "Complete my course of go in Platzi in this week")
    t5 := add_task("Complete my course of C#", "Complete my course of python in Platzi in this week")
    
    list := &taskList{
        tasks: []*task { t1, t2 },
    }
    list.add_to_list(t3)

    list.delete_to_list(0)
    list.tasks[0].mark_complete()

    map_task := make(map[string]*taskList)
    map_task["Jhon"] = list

    list2 := &taskList {
        tasks: []*task { t4, t5 },
    }
    map_task["Nestor"] = list2

    fmt.Printf("Task of Jhon\n")
    map_task["Jhon"].show_list()
    fmt.Printf("\nTask of Nestor\n")
    map_task["Nestor"].show_list()
}
