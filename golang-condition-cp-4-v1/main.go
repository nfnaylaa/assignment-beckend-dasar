package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	
	var price float32
	var totalPrice int = (VIP*30)+(regular*20)+(student*10)
	var totalTicket int = VIP + regular + student

	if  totalPrice >= 100{
		if day % 2 != 0{
			if totalTicket < 5{
				price = float32(totalPrice) * 0.85
			}else{
				price = float32(totalPrice) * 0.75
			}
		}else{
			if totalTicket < 5{
				price = float32(totalPrice) * 0.90
			}else{
				price = float32(totalPrice) * 0.80
			}
		}
	}else{
		price = float32(totalPrice)
	}
	
	return price
	


}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
