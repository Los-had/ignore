package main

import (
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
		content, err := GetRawFileContent(NormalizeUserInput(*language))
		if err != nil {
            log.Fatalln(err)
		}

		if err := HandleFile(content); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("Language not specified")
	}
}

func HandleFile(content string) error {
	err := ioutil.WriteFile(".gitignore", []byte(content), 0644)
    if err != nil {
		return err
	}

	return nil
}

func GetRawFileContent(lang string) (string, error) {
	resp, err := http.Get(RawGithubURL + lang + ".gitignore")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        return "", err
    }

	return string(body), nil
}

func NormalizeUserInput(input string) string {
	return strings.Title(strings.ToLower(input))
}