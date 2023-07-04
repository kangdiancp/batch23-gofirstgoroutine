package main

import (
	"fmt"
	"sync"
)

const totalKurir = 3

// sharing variable
var notifKirim []string

func sudahKirim(id int, wg *sync.WaitGroup) {
	// function yg akan di execute dan resultnya akan disend
	// jika function sudahKirim selesai
	defer wg.Done()
	g := fmt.Sprintf("Kurir %d sudah kirim paket", id)
	notifKirim = append(notifKirim, g)
}

// race conditions
func main() {
	var wait sync.WaitGroup
	wait.Add(totalKurir)
	for i := 0; i < totalKurir; i++ {
		go sudahKirim(i, &wait)
	}
	wait.Wait()

	for _, v := range notifKirim {
		fmt.Println(v)
	}

	fmt.Println("Pengiriman selesai")
}
