package combined

type lemmaItem struct {
	Word        string   `json:"word"`
	Definitions []string `json:"definitions"`
}

type contentItem struct {
	PartOfSpeech string      `json:"part_of_speech"`
	Lemmas       []lemmaItem `json:"lemmas"`
}

// Content is the content of a combined file
type Content []contentItem
