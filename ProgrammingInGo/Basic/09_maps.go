package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["joseph"] = 14
    m["jhon"] = 24
    fmt.Printf("joseph: %d. jhon: %d \n", m["joseph"], m ["jhon"])
    for key, value := range m {
        fmt.Printf("%s have year %d old \n", key, value)
    }
    
    key := "david"
    valueKey, isOk := m[key]
    if isOk {
        fmt.Println("The value of ", key, " is: ", valueKey)
    }
    key2 := "jhon"
    valueKey2, isOk2 := m[key2]
    if isOk2 {
        fmt.Println("The value of ", key, " is:", valueKey2)
    }
}
