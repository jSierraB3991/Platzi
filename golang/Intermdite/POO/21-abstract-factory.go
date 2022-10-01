package main

import (
    "fmt"
)

type ISender interface {
    GetSenderMethod() string
    GetSenderChannel() string
}

type INotificationFactory interface {
    SendNofitication();
    GetSender() ISender
}

type SMSNofitication struct {

}

func (SMSNofitication) SendNofitication() {
    fmt.Println("Sending notification")
}

type SMSNofiticationSender struct {

}

func (SMSNofiticationSender) GetSenderMethod() string {
    return "SMS"
}

func (SMSNofiticationSender) GetSenderChannel() string {
    return "Twilio"
}

func (SMSNofitication) GetSender() ISender {
    return SMSNofiticationSender{}
}

type EmailNofitication struct {

}

func (EmailNofitication) SendNofitication() {
    fmt.Println("Sending eamil")
}

type EmailNofiticationSender struct {

}

func (EmailNofiticationSender) GetSenderMethod() string {
    return "EMAIL"
}

func (EmailNofiticationSender) GetSenderChannel() string {
    return "SES"
}

func (EmailNofitication) GetSender() ISender {
    return EmailNofiticationSender{}
}


func GetNotificarionFactory(notificationType string) (INotificationFactory, error){
    if notificationType == "SMS" {
        return SMSNofitication{}, nil
    }
    if notificationType == "EMAIL" {
        return EmailNofitication{}, nil
    }
    return nil, fmt.Errorf("No notification type")
}

func sendNofitication(f INotificationFactory) {
    f.SendNofitication()
}

func GetMethod(f INotificationFactory){
    fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
    smsFactory, _ := GetNotificarionFactory("SMS")
    emailFactory, _ := GetNotificarionFactory("EMAIL")

    sendNofitication(smsFactory)
    sendNofitication(emailFactory)
    GetMethod(smsFactory)
    GetMethod(emailFactory)
}









