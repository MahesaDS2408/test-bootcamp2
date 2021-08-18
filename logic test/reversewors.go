package main

import (
	"fmt"
	"strings"
)

func balik(masuk string) string {
	nampung1 := ""
	for i := len(masuk) - 1; i >= 0; i-- {
		nampung1 += string(masuk[i])
	}
	return nampung1
}
func input(kata string) string {
	nampung1 := ""
	nampung := strings.Split(kata, " ")

	//fmt.Println(strings.Split(kata, " "))
	for _, el := range nampung {

		nampung1 += (balik(el) + " ")
		//fmt.Print(balik(el), " ")
	}
	return nampung1
}
func main() {
	//var kata string
	fmt.Println("\"i am A Great human\"")
	fmt.Println(input("\"i am A Great human\""))

	//fmt.Println("masukan kata")
	//fmt.Scanf("%s", &kata)
	//strings.Split(kata, " ")

}
