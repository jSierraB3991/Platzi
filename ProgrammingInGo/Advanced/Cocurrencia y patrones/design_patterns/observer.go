package main

import (
    "fmt"
)

type Observer interface {
    getId() string
    updateValue(string)
}

type Topic interface {
    register(observer Observer)
    broadcast()
}

type Item struct {
    observers []Observer
    name string
    available bool
}

func NewItem(name string) *Item {
    return &Item{name: name}
}

func (i *Item) UpdateValue(){
    fmt.Printf("Item %s is available\n", i.name)
    i.available = true
    i.broadcast()
}

func (i *Item) broadcast(){
    for _, observer := range i.observers {
        observer.updateValue(i.name)
    }
}

func (i *Item) register(observer Observer) {
    i.observers = append(i.observers, observer)
}

type EmailClient struct {
    id string
}

func (e *EmailClient) getId() string {
    return e.id
}

func (e *EmailClient) updateValue(value string){
    fmt.Printf("Send email - %s available from client %s\n", value, e.getId())
}

func main(){
    item := NewItem("Rtx 3080")
    firstObserver := &EmailClient{id:"12AB"}
    secondObserver := &EmailClient{id:"34Dc"}

    item.register(firstObserver)
    item.register(secondObserver)
    item.UpdateValue()
}
