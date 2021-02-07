package lemma

// Item is a single part of speech in a lemmas file
type Item struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Lemmas       []string `json:"lemmas"`
	Exists       bool     `json:"exists"`
}

// Content of a lemmas file
type Content []Item
