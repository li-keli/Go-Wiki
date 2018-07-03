package main

import (
	"log"
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"strconv"
)

type HotelDetail struct {
	Detail []Detail `xml:"Detail"`
}
type Detail struct {
	HotelCode            int
	Name                 string  `xml:"Name"`
	Address              string  `xml:"Address"`
	PostalCode           string  `xml:"PostalCode"`
	StarRate             int     `xml:"StarRate"`
	Category             int     `xml:"Category"`
	Phone                string  `xml:"Phone"`
	Fax                  string  `xml:"Fax"`
	EstablishmentDate    string  `xml:"EstablishmentDate"`
	RenovationDate       string  `xml:"RenovationDate"`
	GroupId              int     `xml:"GroupId"`
	BrandId              int     `xml:"BrandId"`
	IsEconomic           int     `xml:"IsEconomic"`
	IsApartment          int     `xml:"IsApartment"`
	ArrivalTime          int     `xml:"ArrivalTime"`
	DepartureTime        int     `xml:"DepartureTime"`
	GoogleLat            float32 `xml:"GoogleLat"`
	GoogleLon            float32 `xml:"GoogleLon"`
	BaiduLat             float32 `xml:"BaiduLat"`
	BaiduLon             float32 `xml:"BaiduLon"`
	CityId               string  `xml:"CityId"`
	CityId2              string  `xml:"CityId2"`
	District             string  `xml:"District"`
	BusinessZone         string  `xml:"BusinessZone"`
	BusinessZone2        string  `xml:"BusinessZone2"`
	CreditCards          string  `xml:"CreditCards"`
	IntroEditor          string  `xml:"IntroEditor"`
	Description          string  `xml:"Description"`
	AirportPickUpService string  `xml:"AirportPickUpService"`
	GeneralAmenities     string  `xml:"GeneralAmenities"`
	RoomAmenities        string  `xml:"RoomAmenities"`
	RecreationAmenities  string  `xml:"RecreationAmenities"`
	ConferenceAmenities  string  `xml:"ConferenceAmenities"`
	DiningAmenities      string  `xml:"DiningAmenities"`
	Traffic              string  `xml:"Traffic"`
	Surroundings         string  `xml:"Surroundings"`
	Features             string  `xml:"Features"`
	Facilities           string  `xml:"Facilities"`
	ServiceRank          string  `xml:"ServiceRank"`
	HasCoupon            bool    `xml:"HasCoupon"`
	FacilitiesV2         string  `xml:"FacilitiesV2"`
	Themes               string  `xml:"Themes"`
	RoomTotalAmount      int     `xml:"RoomTotalAmount"`
	HotelStatus          string  `xml:"HotelStatus"`
}

func getHotelDetail(hotel Hotel) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("getHotelDetail panic 检索异常 -> %s", err)
		}
	}()
	var (
		err            error
		hotelCode      = string([]byte(strconv.Itoa(hotel.HotelId)))
		hotelDetailUrl = fmt.Sprintf("http://api.elongstatic.com/xml/v2.0/hotel/cn/%s/%d.xml", hotelCode[len(hotelCode)-2:], hotel.HotelId)
		details        = HotelDetail{}
	)

	log.Printf("查询URL -> %s", hotelDetailUrl)
	req, err := http.Get(hotelDetailUrl)
	defer req.Body.Close()
	if err != nil {
		checkError("获取响应异常 -> %s", err.Error())
	}

	if req.StatusCode != 200 {
		DbChan <- Detail{HotelCode: hotel.HotelId}
		checkError("响应码:%d, 酒店%s详情检索失败", req.StatusCode, hotelCode)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		checkError("字节转文本异常 -> %s", err.Error())
	}
	err = xml.Unmarshal(body, &details)
	if err != nil {
		checkError("xml反序列化异常 -> %s", err.Error())
	}

	for _, det := range details.Detail {
		det.HotelCode = hotel.HotelId
		DbChan <- det
	}

	return
}
