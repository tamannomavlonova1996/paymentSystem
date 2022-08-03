package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		productName    string
		price          float64
		creditDiapazon int
		phone          string
	)

	fmt.Println("Здравствуйте! Мы рады, что вы выбрали нашу платежную систему.\n " +
		"Введите название продукта, которую хотите приобрести.Для вашего удобства покажем вам список товаров.")

	products := []string{"Смартфон", "Компьютер", "Телевизор"}
	fmt.Println(products)
	fmt.Scan(&productName)

	for i, p := range products {
		if reflect.DeepEqual(productName, p) {
			break
		} else {
			if i == len(products)-1 {
				fmt.Println("Упс, этого товара нет в магазине.")
				return
			}
		}
	}

	fmt.Println("Введите сумму продукта, которую хотите приобрести.")
	fmt.Scan(&price)

	if price <= 0 {
		fmt.Println("Введите правильную сумму, сумма не может быть меньше или равно 0.")
		return
	}

	fmt.Println("Введите номер телефона,чтобы мы могли вам отправить чек +992*********")
	fmt.Scan(&phone)

	if len(phone) != 9 {
		fmt.Println("Неверный номер телефона")
		return
	}

	fmt.Println("Введите срок кредита. Для вашего удобства вывели список допустимых месяцев для оформления кредита.")
	fmt.Scan(&creditDiapazon)

	cd := creditDiap(creditDiapazon)
	if cd == -1 {
		fmt.Println("Выберите правильный срок")
		return
	}

	total := paymentSystem(productName, price, creditDiapazon)
	fmt.Println("Общяя сумма вашего кредита состовляет ", total)

	paymentDetails := SMSMessage{
		ProductName:    productName,
		Price:          price,
		CreditDiapazon: creditDiapazon,
		Phone:          phone,
		TotalSum:       total,
	}

	err := sendSms(paymentDetails)
	if err != nil {
		fmt.Println("Error with sending sms")
		return
	}

}
