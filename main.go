package main

import (
	"fmt"
	"log"
	"net/url"
)

func EncodeURL(text string) string {
	encText := url.QueryEscape(text)
	return encText
}

func DecodeURL(text string) string {
	decText, err := url.QueryUnescape(text)
	if err != nil {
		log.Fatal(err)
	}
	return decText
}

func main() {
	s := "<scritp>alert(0); document.querySelector('var')</scritp>"
	encod := EncodeURL(s)
	decod := DecodeURL(encod)
	fmt.Println(encod)
	fmt.Println(decod)

}
