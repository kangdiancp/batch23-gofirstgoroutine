package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 1. create struct
type Kurir struct {
	id    int
	paket int
}

// 2. slice for kurir
type ListKurir []Kurir

func main() {
	//3. declare variable for listkurir
	listOfKurir := ListKurir{}

	//4. isi kurir field menggunakan random number
	for i := 1; i <= 5; i++ {
		listOfKurir = append(listOfKurir, Kurir{id: i, paket: rand.Intn(5) + 1})
	}

	//5. display kurir list
	for _, v := range listOfKurir {
		fmt.Println(v)
	}

	//6. declare channel
	ch := make(chan int, 2)

	//7. summary total paket
	totalPaket := 0

	//11. add waitgroup
	var wg sync.WaitGroup
	wg.Add(5)
	//subTotal := 0
	//8. loop & use goroutine
	for _, v := range listOfKurir {

		//9. create goroutine with anonymous function
		go func(x Kurir) {
			subTotal := 0
			subTotal += x.paket
			fmt.Println("Kurir", x.id, " kirim paket : ", x.paket)
			ch <- subTotal
			fmt.Println("Kurir", x.id, " selesai kirim paket : ")

		}(v)
		wg.Done()
	}
	wg.Wait()

	//extract value from channel
	for i := 0; i < len(listOfKurir); i++ {
		totalPaket += <-ch
	}

	//result := <-ch

	//10. close channel
	close(ch)

	fmt.Println("Total Paket : ", totalPaket)

}
