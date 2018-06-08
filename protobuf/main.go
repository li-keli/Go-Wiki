package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {
	var (
		url = "http://localhost:8080/getCites"
	)
	client := http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("content-type", "application/x-protobuf")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client do err %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var city *City
	err = proto.Unmarshal(body, city)
	if err != nil {
		fmt.Printf("marshal err %s", err.Error())
	}
	fmt.Println(city.GetCityName())
	fmt.Println("end ... ")
}
