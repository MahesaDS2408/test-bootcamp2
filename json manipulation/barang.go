package main

import (
	"encoding/json"
	"fmt"
)

type Barang []BarangElement

func UnmarshalBarang(data []byte) (Barang, error) {
	var r Barang
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Barang) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BarangElement struct {
	InventoryID int64     `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int64     `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomID int64  `json:"room_id"`
	Name   string `json:"name"`
}

func main() {
	json := `
	[
	  {
		"inventory_id": 9382,
		"name": "Brown Chair",
		"type": "furniture",
		"tags": [
		  "chair",
		  "furniture",
		  "brown"
		],
		"purchased_at": 1579190471,
		"placement": {
		  "room_id": 3,
		  "name": "Meeting Room"
		}
	  },
	  {
		"inventory_id": 9380,
		"name": "Big Desk",
		"type": "furniture",
		"tags": [
		  "desk",
		  "furniture",
		  "brown"
		],
		"purchased_at": 1579190642,
		"placement": {
		  "room_id": 3,
		  "name": "Meeting Room"
		}
	  },
	  {
		"inventory_id": 2932,
		"name": "LG Monitor 50 inch",
		"type": "electronic",
		"tags": [
		  "monitor"
		],
		"purchased_at": 1579017842,
		"placement": {
		  "room_id": 3,
		  "name": "Lavender"
		}
	  },
	  {
		"inventory_id": 232,
		"name": "Sharp Pendingin Ruangan 2PK",
		"type": "electronic",
		"tags": [
		  "ac"
		],
		"purchased_at": 1578931442,
		"placement": {
		  "room_id": 5,
		  "name": "Dhanapala"
		}
	  },
	  {
		"inventory_id": 9382,
		"name": "Alat Makan",
		"type": "tableware",
		"tags": [
		  "spoon",
		  "fork",
		  "tableware"
		],
		"purchased_at": 1578672242,
		"placement": {
		  "room_id": 10,
		  "name": "Rajawali"
		}
	  }
	]
	`
	barang, _ := UnmarshalBarang([]byte(json))

	fmt.Println("1. Find items in the Meeting Room.")
	for _, el := range barang {
		if el.Placement.Name == "Meeting Room" {
			fmt.Println(el.Name)

		}
	}
	fmt.Println("")
	fmt.Println("2. Find all electronic devices.")
	for _, el := range barang {
		if el.Type == "electronic" {
			fmt.Println(el.Name)

		}
	}
	fmt.Println("")
	fmt.Println("3. Find all the furniture.")
	for _, el := range barang {
		if el.Type == "furniture" {
			fmt.Println(el.Name)

		}
	}

	//convert tanggal
	/*i, err := strconv.ParseInt("1579190642", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)*/

	fmt.Println("")
	fmt.Println("4. Find all items were purchased on 16 Januari 2020.")
	for _, el := range barang {
		if el.PurchasedAt == int64(1579190642) {
			fmt.Println(el.Name)

		}
	}

	/*fmt.Println("")
	fmt.Println("5. Find all items with brown color.")
	//var tags[0] = "brown"
	for _, el := range barang {
		//fmt.Println(el.Tags[2])
		/*for el.Tags[2] == "brown" {
			fmt.Println(el.Name)

		}
	}*/
}
