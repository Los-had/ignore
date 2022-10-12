package main

import (
	"os"
	"log"
	"flag"
	"net/http"
	"io/ioutil"
	"strings"
)

var RawGithubURL = "https://raw.githubusercontent.com/github/gitignore/main/"

func main() {
	language := flag.String("language", "", "Language of the gitignore file")
	flag.Parse()

	if *language != "" {
		log.Println("Downloading gitignore from the ", *language, " language")
	} else {
		log.Fatalln("Language not specified")
	}
}

func HandleFile(content string) error {
	return
}

func GetRawFileContent(lang string) (string, error) {
	return
}

func NormalizeUserInput(input string) string {
	return
}