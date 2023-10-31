package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("You must provide a word")
		return
	}

	word := os.Args[1]
	fmt.Printf("Define: %s\n", word)

	res := fetch(word)
	render(res)
	SaveInCache(res.Word, res)
}

func fetch(word string) *DictRes {
	if len(word) == 0 {
		panic("You must provide a word")
	}

	cachedDict := GetFromCache(word)
	if cachedDict != nil {
		return cachedDict
	}

	res, err := http.Get(fmt.Sprintf(
		"https://api.dictionaryapi.dev/api/v2/entries/en/%s", word,
	))
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Could not define", word)
		os.Exit(1)
	}

	var body []DictRes
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		panic(err)
	}

	return &body[0]
}

func render(dict *DictRes) {
	PrintIfNotEmpty(dict.Phonetic)
	PrintIfNotEmpty(dict.Origin)

	for _, meaning := range dict.Meanings {
		fmt.Printf("%s:\n", meaning.PartOfSpeech)
		for _, definition := range meaning.Definitions {
			fmt.Printf("\t%s\n", definition.Definition)
			PrintIfNotEmpty(definition.Example, "\t\tExample: ")
			syn := strings.Join(definition.Synonyms, ", ")
			ant := strings.Join(definition.Antonyms, ", ")
			PrintIfNotEmpty(syn, "\t\tSynonyms: ")
			PrintIfNotEmpty(ant, "\t\tAntonyms: ")
		}
	}
}

func PrintIfNotEmpty(s string, prefix ...string) {
	if s != "" {
		if len(prefix) > 0 {
			for _, p := range prefix {
				s = p + s
			}
		}
		fmt.Println(s)
	}
}
