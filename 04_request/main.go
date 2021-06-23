package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://shibe.online/api/shibes?count=1"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var j []string
	err = json.Unmarshal(byteArray, &j)
	if err != nil {
		panic(err)
	}

	println(j[0])
}
