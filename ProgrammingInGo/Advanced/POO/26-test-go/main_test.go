package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5 ,5)

    if total != 10 {
        t.Errorf("Sum was incorrec, go %d expect %d", total, 10)
    }
}

func TestSumII(t *testing.T) {
    tables := [] struct {
        a int
        b int
        n int
    }{
        {1,2,3}, {2,2,4}, {25,26,51},
    }

    for _, item := range tables {
        total := Sum(item.a, item.b)
        if total != item.n {
            t.Errorf("Sum was incorrect, go %d, expect %d", total, item.n)
        }
    }
}

func TestGetMax(t * testing.T) {
    tables := [] struct {
        a int
        b int
        n int
    }{
        {3,2,3},{5,4,5},{9,4,9},{23,35,35},
    }

    for _, item := range tables {
        more := GetMax(item.a, item.b)
        if more != item.n {
            t.Errorf("Get max is incorrect, got %d expect %d", more, item.n)
        }
    }
}

func TestFibonacci(t *testing.T) {
    tables := [] struct {
        a int
        n int
    }{
        {1,1},{8,21},{50,12586269025},
    }

    for _, item := range tables {
        result := Fibonacci(item.a)
        if result != item.n {
            t.Errorf("The fibonacci secuence is incorrect, get %d expected %d", result, item.n)
        }
    }
}
