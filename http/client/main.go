package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
)

var (
	imgAdd    = "./test.jpeg"
	serverURL = "http://localhost:5200/imgcode"
)

func main() {
	f, _ := filepath.Abs(imgAdd)
	bytesOfFile, _ := ioutil.ReadFile(f)

	_, info, _ := gorequest.New().Post(serverURL).
		Type("multipart").
		SendFile(bytesOfFile, "img", "img").
		End()

	fmt.Println(info)
}
