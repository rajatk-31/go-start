package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000"
	response, _ := http.Get(myUrl)
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	var resString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := resString.Write(content)
	// fmt.Println("----", string(content))
	fmt.Println(byteCount)
	fmt.Println(resString.String())
}

func PerformPostRequest() {
	requestBody := strings.NewReader(`
	 {
		"country" :"India",
		"Name": "rajat"
	 }
	`)

	res, _ := http.Post("http://localhost:3000/post", "application/json", requestBody)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(data))
}

func PerformFormRequest() {
	const myUrl = "http://localhost:3000/postForm"

	data := url.Values{}

	data.Add("firstName", "rajat")
	data.Add("lastName", "yadav")
	data.Add("greet", "hello")

	res, _ := http.PostForm(myUrl, data)
	defer res.Body.Close()
	datas, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(datas))
}
