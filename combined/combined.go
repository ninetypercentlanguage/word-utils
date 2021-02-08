package combined

// LemmaItem is single item in ContentItem Lemmas
type LemmaItem struct {
	Word        string   `json:"word"`
	Definitions []string `json:"definitions"`
}

// ContentItem is single item in the array contained in a Content file
type ContentItem struct {
	PartOfSpeech string      `json:"part_of_speech"`
	Lemmas       []LemmaItem `json:"lemmas"`
}

// Content is the content of a combined file
type Content []ContentItem

// File is the output of the wiktionary-combined program
type File struct {
	Word       string
	Content    Content
	HasContent bool
}
