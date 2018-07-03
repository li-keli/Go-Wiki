package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var DbChan = make(chan Detail, 10000)

type DbMysql struct {
	db *sql.DB
}

func (d *DbMysql) init() {
	path := strings.Join([]string{"quner", ":", "HXd!@#123", "@tcp(", "172.16.5.182", ":", "3306", ")/", "elong_channel", "?charset=utf8"}, "")
	log.Println(path)

	d.db, _ = sql.Open("mysql", path)

	if err := d.db.Ping(); err != nil {
		log.Fatalf("opon database fail -> %s", err)
	}

	d.db.SetMaxOpenConns(200)
	d.db.SetMaxIdleConns(100)
	log.Println("等待写入数据...")

	for {
		select {
		case detail := <-DbChan:
			stmt, err := d.db.Prepare("replace into elong_hotel (HotelCode, Name, Address, PostalCode, StarRate, Category, Phone, Fax, EstablishmentDate, RenovationDate, GroupId, BrandId, IsEconomic, IsApartment, ArrivalTime, DepartureTime, GoogleLat, GoogleLon, BaiduLat, BaiduLon, CityId, CityId2, District, BusinessZone, BusinessZone2, CreditCards, IntroEditor, Description, AirportPickUpService, GeneralAmenities, RoomAmenities, RecreationAmenities, ConferenceAmenities, DiningAmenities, Traffic, Surroundings, Features, Facilities, ServiceRank, HasCoupon, FacilitiesV2, Themes, RoomTotalAmount, HotelStatus) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
			_, err = stmt.Exec(detail.HotelCode, detail.Name, detail.Address, detail.PostalCode, detail.StarRate, detail.Category, detail.Phone, detail.Fax, detail.EstablishmentDate, detail.RenovationDate, detail.GroupId, detail.BrandId, detail.IsEconomic, detail.IsApartment, detail.ArrivalTime, detail.DepartureTime, detail.GoogleLat, detail.GoogleLon, detail.BaiduLat, detail.BaiduLon, detail.CityId, detail.CityId2, detail.District, detail.BusinessZone, detail.BusinessZone2, detail.CreditCards, detail.IntroEditor, detail.Description, detail.AirportPickUpService, detail.GeneralAmenities, detail.RoomAmenities, detail.RecreationAmenities, detail.ConferenceAmenities, detail.DiningAmenities, detail.Traffic, detail.Surroundings, detail.Features, detail.Facilities, detail.ServiceRank, detail.HasCoupon, detail.FacilitiesV2, detail.Themes, detail.RoomTotalAmount, detail.HotelStatus)
			if err != nil {
				log.Printf("写入异常 -> %s \n", err)
			}
		}

	}
}
