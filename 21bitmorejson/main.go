package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("JSON in GOLANG")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"LCO DevOps", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"LCO Kubernetes", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"Name": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Password": "abc123",
		"Tags": ["web-dev", "js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)

	}
}
