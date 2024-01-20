package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Files in GOLANG")
	content := "This needs to go in a file"
	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		log.Fatal(err)
	}
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Printf("All done with file of %v characters\n", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text file content: ", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
