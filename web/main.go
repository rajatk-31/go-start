package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://aaditya.dev:3000/path1?username=rahst&age=56"

func main() {
	fmt.Println("Hello from web request")
	// response, err := http.Get(myUrl)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%T\n", response)
	// defer response.Body.Close()
	// data, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	parsed, _ := url.Parse(myUrl)
	fmt.Println(parsed.Scheme)
	fmt.Println(parsed.Path)
	fmt.Println(parsed.Host)
	fmt.Println(parsed.Port())
	fmt.Println(parsed.RawQuery)

	lquery := parsed.Query()

	fmt.Println(lquery)

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "aaditya.dev",
		Path:    "/rajat",
		RawPath: "user=rajat",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
