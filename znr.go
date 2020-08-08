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

// Trie holds all the known words, phrases and sentences.
type Trie struct {
	root *trieNode
}

type trieNode struct {
	isWordEnd bool
	children  map[rune]*trieNode
}

// NewTrie returns a Trie with a root node.
func NewTrie() Trie {
	tr := Trie{root: &trieNode{}}
	return tr
}

// Insert adds a word or phrase to the Trie.
func (tr *Trie) Insert(word string) {
	current := tr.root

	if current.children == nil {
		current.children = make(map[rune]*trieNode)
	}

	for _, c := range word {
		if _, ok := current.children[c]; ok == false {
			current.children[c] = &trieNode{
				children: make(map[rune]*trieNode),
			}
		}
		current = current.children[c]
	}
	current.isWordEnd = true
}

// Find returns whether a word exists within the Trie.
func (tr *Trie) Find(word string) bool {
	current := tr.root
	for _, c := range word {
		if _, ok := current.children[c]; ok == false {
			return false
		}
		current = current.children[c]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

// KnownPhrases takes a text string and returns a slice containing all phrases
// found in the Trie.
func (tr *Trie) KnownPhrases(t string) ([]string, error) {
	known := []string{}

	ph := []rune{}
	cur := tr.root

	rs := []rune(t)
	for i := 0; i < len(rs); i++ {
		r := rs[i]

		// if the current rune isn't in the trie...
		if _, ok := cur.children[r]; ok == false {

			// Check to see if we matched a phrase at the previous node
			// and add it to the slice of known phrases.
			if cur.isWordEnd {
				known = append(known, string(ph))
				i-- // Ff so, we need to check this rune again.
			}

			// Begin searching at the root of the Trie again.
			cur = tr.root
			ph = []rune{}

			continue
		}

		// Rune was found at the current node in the Trie so go deeper.
		cur = cur.children[r]
		ph = append(ph, r)

		// If this is the last rune in the text then check to see if we
		// matched a phrase.
		if i == len(rs)-1 && cur.isWordEnd {
			known = append(known, string(ph))
			ph = []rune{}
		}
	}

	return known, nil
}
