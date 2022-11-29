package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"Node js", 499, "Online", "test", []string{"backend", "web-dev"}},
		{"Angular", 399, "Online", "test1", []string{"frontend", "web-dev"}},
		{"Monog DB", 299, "Online", "test2", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	
        {
                "coursename": "Node js",
                "Price": 499,
                "Platform": "Online",
                "tags": [
                        "backend",
                        "web-dev"
                ]
        }
	`)

	var courseData course
	checkValid := json.Valid(jsonData)
	if checkValid {
		json.Unmarshal(jsonData, &courseData)
		fmt.Printf("%#v\n", courseData)
	}

	var mydata map[string]interface{}
	json.Unmarshal(jsonData, &mydata)
	fmt.Printf("%#v\n", mydata)

}
