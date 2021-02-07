package lemma

type item struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Lemmas       []string `json:"lemmas"`
	Exists       bool     `json:"exists"`
}

// Content of a lemmas file
type Content []item
