package main

import (
    "fmt"
)

type Payment interface{
    Pay()
}

type CashPayment struct {}

func (CashPayment) Pay() {
    fmt.Println("Payment using cash")
}

func ProcessPayment(p Payment){
    p.Pay()
}

type BankPayment struct {}

func (BankPayment) Pay(bankAccount int) {
    fmt.Printf("Payment using bank account %d\n", bankAccount)
}

type BankPaymentAdapter struct{
    bankPay *BankPayment
    bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
    bpa.bankPay.Pay(bpa.bankAccount)
}

func main() {
    cash := &CashPayment{}
    ProcessPayment(cash)

    bank := &BankPaymentAdapter{
        bankAccount: 5,
        bankPay: &BankPayment{},
    }
    ProcessPayment(bank)
}
