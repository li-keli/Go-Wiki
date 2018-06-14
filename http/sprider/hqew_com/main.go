package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"github.com/satori/go.uuid"
)

//func doGet(url string) (resp *http.Response, err error) {
//	resp, err = http.Get(url)
//	if err != nil {
//		fmt.Println(resp.StatusCode)
//		fmt.Println(err)
//		log.Fatal(err)
//	}
//
//	return resp, err
//
//}
//
//func doPost(url string, bodyType string) (*http.Response, error) {
//	resp, err := http.Post(url, bodyType, nil)
//
//	if err != nil {
//		fmt.Println(resp.StatusCode)
//		fmt.Println(err)
//		log.Fatal(err)
//	}
//
//	return resp, err
//}

func main() {

	c := make(chan int, 1)

	for i := 0; i < 10; i++ {
		go func() {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://s.hqew.com/1N4148_____0_00_0_0_0_2.html", nil)
			u1, _ := uuid.NewV4()

			req.Header.Add("Cookie", fmt.Sprintf("HQEWVisitor=%s; _qddaz=QD.wvl0mt.1oqtha.jidue3p7;", u1))
			if err != nil {
				fmt.Println(err)
			}
			resp, err := client.Do(req)
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}

				fmt.Println(strings.Count(string(body), ""))
			}
		}()
	}

	fmt.Println("ok")
	<-c

}
