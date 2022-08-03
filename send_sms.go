package main

import (
	"fmt"
	textmagic "github.com/textmagic/textmagic-rest-go"
)

const (
	username = "tamannahmavlon"
	token    = "vsFQRUWwVCOL8iKevNhaQFAGzFBTkf"
)

type SMSMessage struct {
	ProductName    string
	Price          float64
	CreditDiapazon int
	Phone          string
	TotalSum       float64
}

func sendSms(paymentDetails SMSMessage) error {
	client := textmagic.NewClient(username, token)
	// Слишком длинный чек для беслпатной отправки :(
	//message := "Детали платежа \nНаименование товара: " + paymentDetails.ProductName +
	//	"\nЦена товара: " + fmt.Sprintf("%f", paymentDetails.Price) +
	//	"\nСрок рассрочки: " + fmt.Sprintf("%d", paymentDetails.CreditDiapazon) +
	//	"\nНомер телефона получателя: " + paymentDetails.Phone + "\nОбщяя сумма: " +
	//	fmt.Sprintf("%f", paymentDetails.TotalSum)

	params := map[string]string{
		"phones": "+992" + paymentDetails.Phone,
		"text":   "Короткий текст", //message
	}
	_, err := client.CreateMessage(params)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
