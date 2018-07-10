package main

import (
	"log"
	"time"
	"net/http"
)

var (
	mysqlDb DbMysql
	maxChan = make(chan int, 100)
)

func main() {
	var hotels = readAllHotelId()
	go mysqlDb.init()
	time.Sleep(3 * time.Second)


	for index, hotel := range hotels {
		log.Printf("检索索引 -> %d, 酒店编号 -> %d", index, hotel.HotelId)
		maxChan <- 1
		go getHotelDetail(hotel)
	}

	req, _ := http.Get("https://sc.ftqq.com/SCU26858T43a625a0d78cdf7ad88844fc6b2047b35b05149cf0fd7.send?text=艺龙全量酒店同步完成")
	defer req.Body.Close()

	log.Println("job 艺龙全量酒店同步完成")
}

func checkError(info string, args ...interface{}) {
	log.Printf(info, args)
	panic(args)
}
