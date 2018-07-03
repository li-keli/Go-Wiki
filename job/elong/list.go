package main

import (
	"io/ioutil"
	"log"
	"encoding/xml"
)

type HotelIndex struct {
	Hotels []Hotels `xml:"Hotels"`
}

type Hotels struct {
	Hotel []Hotel `xml:"Hotel"`
}

type Hotel struct {
	HotelId int `xml:"HotelId,attr"`
	Status  int `xml:"Status,attr"`
}

var xmlAddress = "hotellist.xml"

func readAllHotelId() (hotel []Hotel) {
	content, err := ioutil.ReadFile(xmlAddress)
	if err != nil {
		log.Fatal(err)
	}

	var hotelIndex HotelIndex
	err = xml.Unmarshal(content, &hotelIndex)
	if err != nil {
		log.Fatal(err)
	}

	hotel = hotelIndex.Hotels[0].Hotel[90000:]

	log.Printf("酒店总数量 -> %d", len(hotel))
	return
}
