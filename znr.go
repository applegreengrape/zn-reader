package znr

// VocabList defines a list of Vocab items.
type VocabList []Vocab

// Vocab defines a single character, word, or phrase in Chinese.
type Vocab struct {
	Lang             string     `json:"lang,omitempty"`
	HeisigDefinition string     `json:"heisigDefinition,omitempty"`
	Ilk              string     `json:"ilk,omitempty"`
	Writing          string     `json:"writing,omitempty"`
	Definitions      Definition `json:"definitions,omitempty"`
	Reading          string     `json:"reading,omitempty"`
	ID               string     `json:"id,omitempty"`
}

// Definition maps a language to a meaning e.g. {"en": "hello"}
type Definition map[string]string

// KnownPhrases takes a text string and a VocabList and returns a slice
// containing the phrases in the VocabList that appear in the text.
func KnownPhrases(t string, vl VocabList) ([]string, error) {
	known := []string{}
	for _, c := range t {
		for _, v := range vl {
			if string(c) == v.Writing {
				known = append(known, v.Writing)
			}
		}
	}
	return known, nil
}
