package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	file, err := os.Open("raw-test-file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	str, err1 := getStringFromFile(file)
	if err1 != nil {
		panic(err1)
	}
	file.Close()
	reg := regexp.MustCompile(`(?m)0;31m([^=.]*)\[0m`)
	matches := reg.FindStringSubmatch(str)
	fmt.Printf("%v", matches[1])

}

func getStringFromFile(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
