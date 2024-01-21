package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

func main() {
	fmt.Println("URLs in GOLANG")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n", qparams)

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//building a url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=aman",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
