package znr

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
