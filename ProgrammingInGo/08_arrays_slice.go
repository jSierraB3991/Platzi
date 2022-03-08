package main

import (
    "fmt"
    "strings"
)

func isPalindrome(text string) {
    var reverse string

    for i := len(text)-1; i >=0; i-- {
        reverse += strings.ToLower(string(text[i]))
    }
    if strings.ToLower(text) == reverse {
        fmt.Printf("\nThe text '%s' is palindrome", reverse)
    } else {
        fmt.Printf("\nThe text '%s' not is palindrome", text)
    }
    fmt.Println("");
}

func main() {

    //Array
    fmt.Printf("Arrays: ")
    var array [4]int
    array[0] = 2
    array[1] = 3
    fmt.Println(array)
    fmt.Println("number of elements: ", len(array), " capacity of array: ", cap(array))

    //Slices
    fmt.Printf("\nSlices: ")
    slice := []int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println(slice)
    fmt.Println("number of elements: ", len(slice), " capacity of slice: ", cap(slice))

    //Methods in the slice
    fmt.Println(slice[0])
    fmt.Println(slice[:3])
    fmt.Println(slice[2:4])
    fmt.Println(slice[4:])
    
    slice = append(slice, 8)
    fmt.Println(slice)

    slice2 := []int {9, 19, 11}
    slice = append(slice, slice2...)
    fmt.Println(slice)

    slice3 := []string { "hola", "que" ,"hace" }
    // for _, value := range slice3 {
    // for i, _ := range slice3 {
    for i, value := range slice3 {
        fmt.Println(i, value)
    }

    isPalindrome("amor a romA")
}
