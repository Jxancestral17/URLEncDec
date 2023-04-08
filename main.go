package main

import (
	"fmt"
	"html"
	"log"
	"net/url"
	"os"
	"strings"
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

func EncodeEntity(text string) string {
	encText := html.EscapeString(text)
	return encText
}

func DecodeEntity(text string) string {
	decText := html.UnescapeString(text)
	return decText
}

func Read(text string) string {
	content, err := os.ReadFile(text)
	if err != nil {
		log.Fatal(err)
	}
	k := string(content)
	g := strings.ReplaceAll(k, "\n", "")
	s := strings.ReplaceAll(g, " ", "")
	return s
}

func help() {
	fmt.Println("[+]********************************************[+]")
	fmt.Println("[+]\t☠️    JXSS By ЙАКН³Щ⁺РЕ³!\t☠️      [+]")
	fmt.Println("[+]\t\t\t©2023\t\t       [+]")
	fmt.Println("[+]\t\t\t\t\t       [+]")
	fmt.Println("[+]\t\t\t\t\t       [+]")
	fmt.Println("[+]\t -h Help\t\t\t       [+]\n[+]\t -i input file ex: xss.js\t       [+]\n[+]\t -ue Encode URL\t\t\t       [+]\n[+]\t -ud Decode URL\t\t\t       [+]")
	fmt.Println("[+]\t -ed Decode Entity\t\t       [+]\n[+]\t -ee Encode Entity\t\t       [+]")
	fmt.Println("[+]\t\t\t\t\t       [+]")
	fmt.Println("[+]\t\t\t\t\t       [+]")
	fmt.Println("[+]\tex: ./urlencdec -e -i xss.js\t       [+]")
	fmt.Println("[+]\t\t\t\t\t       [+]")
	fmt.Println("[+]\t\t\t\t\t       [+]")
	fmt.Println("[+]********************************************[+]")
}

func checkError(length int) {
	if length == 2 {
		log.Fatal("Miss params -i path_file ")
	} else if length == 3 {
		log.Fatal("Miss params the path_file ")
	}
}

func main() {
	args := os.Args
	length := len(args)
	if len(args) == 1 {
		help()
	} else {
		if args[1] == "-h" {
			help()
		} else if args[1] == "-ue" {
			checkError(length)
			if args[2] == "-i" {
				content := Read(args[3])
				fmt.Println(EncodeURL(content))
			}
		} else if args[1] == "-ud" {
			checkError(length)
			if args[2] == "-i" {
				content := Read(args[3])
				fmt.Println(DecodeURL(content))
			}
		} else if args[1] == "-ee" {
			checkError(length)
			if args[2] == "-i" {
				content := Read(args[3])
				fmt.Println(EncodeEntity(content))
			}
		} else if args[1] == "-ed" {
			checkError(length)
			if args[2] == "-i" {
				content := Read(args[3])
				fmt.Println(DecodeEntity(content))
			}
		} else {
			fmt.Println("urlencdec -h for help")
		}
	}

}
