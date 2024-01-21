package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Request Verbs in GOLANG")
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response Status:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Response Body:", responseString.String())

	// fmt.Println(content)
	// fmt.Println("Response Body:", string(content))

	defer response.Body.Close()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake Json Payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":"0",
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
