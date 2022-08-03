package main

func paymentSystem(productName string, price float64, creditDiapazon int) float64 {
	creditComissions := map[string]int{
		"Смартфон":  3,
		"Компьютер": 4,
		"Телевизор": 5,
	}

	var total float64

	switch productName {
	case "Смартфон":
		if creditDiapazon > 9 {
			cd := creditDiap(creditDiapazon) - 2
			total = price + price*float64(cd*creditComissions["Смартфон"])/100
		} else {
			return price
		}

	case "Компьютер":
		if creditDiapazon > 12 {
			cd := creditDiap(creditDiapazon) - 3
			total = price + price*float64(cd*creditComissions["Компьютер"])/100
		} else {
			return price
		}
	case "Телевизор":
		if creditDiapazon > 18 {
			cd := creditDiap(creditDiapazon) - 4
			total = price + price*float64(cd*creditComissions["Телевизор"])/100
		} else {
			return price
		}
	}
	return total
}

func creditDiap(diapazon int) (indexDiapazon int) {
	creditDiapazons := []int{3, 6, 9, 12, 18, 24}
	for i, val := range creditDiapazons {
		if diapazon == val {
			return i
		}
	}
	return -1
}
