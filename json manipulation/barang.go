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

	fmt.Println("item meeting room")
	for _, el := range barang {
		if el.Placement.Name == "Meeting Room" {
			fmt.Println(el)

		}
	}
	/*for _, el := range barang {
		if el.Placement.Name == "Meeting Room" {
			fmt.Println(el.Name)

		}
	}*/
}
