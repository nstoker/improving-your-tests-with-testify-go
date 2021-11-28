package main

import "fmt"

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

type MessageService interface {
	SendChargeNotification(int) error
}

type SMSService struct{}

type MyService struct {
	messageService MessageService
}

func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending production charge notification")
	return nil
}

func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging customer for %d\n", value)
	return nil
}

func main() {
	fmt.Println("Hello World")

	smsService := SMSService{}
	MyService := MyService{smsService}
	MyService.ChargeCustomer(100)
}
