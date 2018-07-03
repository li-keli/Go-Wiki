package main

import (
	"log"
)

var (
	mysqlDb DbMysql
)

func main() {
	var hotels = readAllHotelId()
	go mysqlDb.init()

	go do(hotels[:30000])
	go do(hotels[30000:60000])
	go do(hotels[60000:])

	log.Println("wating...")
	select {}
}

func do(hotels []Hotel) {
	for index, hotel := range hotels {
		log.Printf("检索索引 -> %d", index)
		getHotelDetail(hotel)
	}
}
