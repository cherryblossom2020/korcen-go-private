package korcen

import "strings"

type Profanity struct {
	CensorWords     map[string]bool
	CharMapping     map[rune][]rune
	AllowedChars    map[rune]bool
	MaxCombinations int
}

func NewProfanity(words []string) *Profanity {
	censorWords := make(map[string]bool)
	for _, word := range words {
		censorWords[strings.ToLower(word)] = true
	}
	charMapping := map[rune][]rune{
		'a': {'a', '@', '*', '4'},
		'i': {'i', '*', 'l', '1'},
		'o': {'o', '*', '0', '@'},
		'u': {'u', '*', 'v'},
	}

	allowedChars := make(map[rune]bool)
	for _, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" {
		allowedChars[r] = true
	}

	return &Profanity{
		CensorWords:     censorWords,
		CharMapping:     charMapping,
		AllowedChars:    allowedChars,
		MaxCombinations: 1,
	}
}

func (p *Profanity) Censor(text string) (bool, string) {
	words := strings.Fields(text)
	for _, word := range words {
		if p.ContainsProfanity(word) {
			return true, word
		}
	}
	return false, ""
}

func (p *Profanity) ContainsProfanity(word string) bool {
	_, exists := p.CensorWords[strings.ToLower(word)]
	return exists
}
