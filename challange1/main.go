package main

import "fmt"

type Kurir struct {
	id    int
	name  string
	area  string
	paket int
}

type ListKurir []*Kurir

type KurirGroup []*Kurir

type KurirGroupArea = map[string]KurirGroup

func main() {
	fmt.Println("start")
	kurirList := ListKurir{
		&Kurir{id: 1, name: "asep", area: "Area A", paket: 10},
		&Kurir{id: 2, name: "budi", area: "Area B", paket: 15},
		&Kurir{id: 3, name: "charlie", area: "Area C", paket: 25},
		&Kurir{id: 4, name: "dendy", area: "Area B", paket: 5},
		&Kurir{id: 5, name: "erwin", area: "Area A", paket: 30},
	}

	kurirGroupArea := make(KurirGroupArea)

	for _, v := range kurirList {
		if _, ok := kurirGroupArea[v.area]; ok {
			kurirGroupArea[v.area] = append(kurirGroupArea[v.area], v)
		} else {
			kurirGroupArea[v.area] = KurirGroup{v}
		}
	}

	// declare channel
	channel := make(chan int, 2)

	for area, group := range kurirGroupArea {

		go func(area string, data KurirGroup, resultChannel chan int) {
			subTotal := 0
			for _, v := range data {
				subTotal += v.paket
			}
			fmt.Println(area, "subtotal : ", subTotal)
			resultChannel <- subTotal
		}(area, group, channel)
	}

	//extract value from channel
	totalPaket := 0
	for i := 0; i < len(kurirGroupArea); i++ {
		totalPaket += <-channel
	}

	close(channel)
	fmt.Println("Total Paket : ", totalPaket)
}
