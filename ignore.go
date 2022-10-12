package main

import (
	"os"
	"log"
	"flag"
	"strings"
	"net/http"
	"io/ioutil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		log.Println("Finished downloading gitignore file")
	} else {
		log.Fatalln("Language not specified")
	}
}

func HandleFile(content string) error {
	err := os.WriteFile(".gitignore", []byte(content), 0644)
    if err != nil {
		return err
	}

	return nil
}

func GetRawFileContent(lang string) (string, error) {
	resp, err := http.Get(RawGithubURL + lang + ".gitignore")
	if err != nil {
		log.Println("Request failed")
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Fail parsing the response body")
        return "", err
    }

	return string(body), nil
}

func NormalizeUserInput(input string) string {
	caser := cases.Title(language.Und)
	return caser.String(strings.TrimSpace(input))
}