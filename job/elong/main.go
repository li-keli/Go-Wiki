package main

import (
	"log"
	"sync"
	"time"
)

var (
	wait    sync.WaitGroup
	mysqlDb DbMysql
)

func main() {
	var hotels = readAllHotelId()
	go mysqlDb.init()
	time.Sleep(3 * time.Second)

	go do(hotels[:30000])
	go do(hotels[30000:60000])
	go do(hotels[60000:])

	log.Println("wating...")
	wait.Wait()

	log.Println("job 完成")
}

func do(hotels []Hotel) {
	wait.Add(1)
	for index, hotel := range hotels {
		log.Printf("检索索引 -> %d, 酒店编号 -> %d", index, hotel.HotelId)
		getHotelDetail(hotel)
	}
	wait.Done()
}

func checkError(info string, args ...interface{}) {
	log.Printf(info, args)
	panic(args)
}
