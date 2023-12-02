package main

type DictRes struct {
	Word     string        `json:"word"`
	Phonetic string        `json:"phonetic"`
	Origin   string        `json:"origin"`
	Meanings []DictMeaning `json:"meanings"`
	Error    string        `json:"error"`
}

type DictMeaning struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definitions `json:"definitions"`
}

type Definitions struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}
