package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5 ,5)

    if total != 10 {
        t.Errorf("Syn wa incorrec, go %d expect %d", total, 10)
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
