package main

import (
	"fmt"
	"strings"
)

func polindrom(str string) string {
	nampung := ""
	for i := len(str) - 1; i >= 0; i-- {
		nampung += string(str[i])
	}
	for i := range str {
		//strings.ToLower(str[i])
		if strings.ToLower(string(str[i])) != strings.ToLower(string(nampung[i])) {
			return "not polindrom"
		}
	}
	return "polindrom"
}

func main() {
	print("Radar			# --> ")
	fmt.Println(polindrom("Radar"))
	print("Malam			# --> ")
	fmt.Println(polindrom("Malam"))
	print("Kasur ini rusak		# --> ")
	fmt.Println(polindrom("Kasur ini rusak"))
	print("Ibu Ratna antar ubi	# --> ")
	fmt.Println(polindrom("Ibu Ratna antar ubi"))
	print("Malas			# --> ")
	fmt.Println(polindrom("Malas"))
	print("Makan nasi goreng	# --> ")
	fmt.Println(polindrom("Makan nasi goreng"))
	print("Balonku ada lima	# --> ")
	fmt.Println(polindrom("Balonku ada lima"))

}
