package main

import (
	"fmt"
	"math"
)

func main() {
	var awal, akhir int

	//masukin tahun
	fmt.Println("masukan tahun awal")
	fmt.Scan(&awal)
	fmt.Println("masukan tahun akhir")
	fmt.Scan(&akhir)

	//tahun keluar
	//fmt.Println(awal)
	//math.Mod(float64(awal), float64(4))
	fmt.Println("")
	for i := awal; i <= akhir; i++ {

		if math.Mod(float64(i), float64(4)) == 0 {
			fmt.Print(i)
			if i != (akhir - int(math.Mod(float64(akhir), float64(4)))) {
				fmt.Print(",")
			}
		}
	}
}
