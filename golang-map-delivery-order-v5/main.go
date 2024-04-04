package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	tenPercent := "senin rabu jumat"
	bill := make(map[string]float32)
	for _, datas := range data {
		parts := strings.Split(datas, ":")
		firstName := parts[0]
		lastName := parts[1]
		price, _ := strconv.ParseFloat(parts[2], 32)
		code := parts[3]

		if day == tenPercent {
			price = (price * 0.1) + price
		} else {
			price = (price * 0.05) + price
		}

		if code == "BDG" {
			if !(day == "rabu" || day == "kamis" || day == "sabtu") {
				continue
			}
		}
		if code == "JKT" {
			if day == "minggu" {
				continue
			}
		}
		if code == "BKS" {
			if !(day == "selasa" || day == "kamis" || day == "jumat") {
				continue
			}
		}
		if code == "DPK" {
			if !(day == "senin" || day == "selasa") {
				continue
			}
		}

		bill[firstName+"-"+lastName] = float32(price)

	}
	return bill // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
