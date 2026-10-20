package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Data JSON dari playground
var jsonString = `[
	{"name": "Samsung Galaxy A52", "vendor": "Samsung"},
	{"name": "iPhone 12", "vendor": "Apple"},
	{"name": "Sony Xperia 1 III", "vendor": "Sony"},
	{"name": "Samsung Galaxy S21", "vendor": "Samsung"},
	{"name": "Sony Xperia 5 III", "vendor": "Sony"}
]`

// struct untuk soal 2
type Device struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
}

// struct untuk soal 3
type Book struct {
	Title, Desc, Author string
	ReleaseYear         int
}

func main() {
	// soal 1
	var dataMap []map[string]string
	err1 := json.Unmarshal([]byte(jsonString), &dataMap)
	if err1 != nil {
		fmt.Println("Error decode json:", err1)
	} else {
		fmt.Println("=== SOAL 1 (Vendor Samsung) ===")
		for _, item := range dataMap {
			if strings.EqualFold(item["vendor"], "samsung") {
				fmt.Printf("Nama: %s, Vendor: %s\n", item["name"], item["vendor"])
			}
		}
	}

	fmt.Println()

	// soal 2
	var dataStruct []Device
	err2 := json.Unmarshal([]byte(jsonString), &dataStruct)
	if err2 != nil {
		fmt.Println("Error decode json:", err2)
	} else {
		fmt.Println("=== SOAL 2 (Vendor Sony) ===")
		for _, item := range dataStruct {
			if strings.EqualFold(item.Vendor, "sony") {
				fmt.Printf("Nama: %s, Vendor: %s\n", item.Name, item.Vendor)
			}
		}
	}

	fmt.Println()

	// soal 3
	books := []Book{
		{Title: "Laskar Pelangi", Desc: "Kisah perjuangan anak-anak Belitung", Author: "Andrea Hirata", ReleaseYear: 2005},
		{Title: "Bumi Manusia", Desc: "Kisah Minke di era kolonial", Author: "Pramoedya Ananta Toer", ReleaseYear: 1980},
		{Title: "Filosofi Kopi", Desc: "Kumpulan cerita pendek tentang kehidupan", Author: "Dee Lestari", ReleaseYear: 2006},
		{Title: "Cantik Itu Luka", Desc: "Kisah sejarah dan kehidupan di Halimunda", Author: "Eka Kurniawan", ReleaseYear: 2002},
		{Title: "Negeri 5 Menara", Desc: "Kisah santri mengejar impian", Author: "A. Fuadi", ReleaseYear: 2009},
	}

	jsonData, err3 := json.MarshalIndent(books, "", "  ")
	if err3 != nil {
		fmt.Println("Error encode json:", err3)
	} else {
		fmt.Println("=== SOAL 3 (Encode JSON) ===")
		fmt.Println(string(jsonData))
	}
}
