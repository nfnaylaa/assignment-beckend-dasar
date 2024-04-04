package main

import "fmt"

func TicketPlayground(height, age int) int {

	// - Jika Umur di bawah 5 tahun maka tidak dapat membeli tiket wahana dengan menampilkan `-1` (repersentasi dari `error`)
	// - Jika anak umur 5-7 tahun atau tinggi lebih dari 120 cm maka akan dikenakan tarif Rp 15.000.
	// - Jika anak umur 8-9 tahun atau tinggi lebih dari 135 cm maka akan dikenakan tarif Rp 25.000.
	// - Jika anak umur 10-11 tahun atau tinggi lebih dari 150 cm maka akan dikenakan tarif Rp 40.000.
	// - Jika anak umur 12 tahun atau tinggi lebih dari 160 cm, akan dikenakan tarif Rp 60.000.
	// - Jika di atas 12 tahun, akan mendapat tiket khusus Remaja yaitu sebesar Rp 100.000

	var priceTicket int

	if age > 12 {
		priceTicket = 100000
	} else if age == 12 || height > 160 {
		priceTicket = 60000
	} else if age == 11 || height > 150 {
		priceTicket = 40000
	} else if age == 10 || height > 150 {
		priceTicket = 40000
	} else if age == 9 || height > 135 {
		priceTicket = 25000
	} else if age == 8 || height > 135 {
		priceTicket = 25000
	} else if age >= 5 && age <= 7 || height > 120 {
		priceTicket = 15000
	} else {
		priceTicket = -1
	}

	return priceTicket // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(165, 10))
}
