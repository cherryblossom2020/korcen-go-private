package korcen

import (
	"strings"
)

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
		MaxCombinations: 3,
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
	word = p.cleanWord(word)
	for _, processedWord := range p.generateCombinations(word) {
		if p.CensorWords[strings.ToLower(processedWord)] {
			return true
		}
	}
	return false
}

func (p *Profanity) cleanWord(word string) string {
	var cleanedWord strings.Builder
	for _, r := range word {
		if _, exists := p.AllowedChars[r]; exists {
			cleanedWord.WriteRune(r)
		}
	}
	return cleanedWord.String()
}

func (p *Profanity) generateCombinations(word string) []string {
	var combinations []string
	combinations = append(combinations, word)
	for i := 0; i < p.MaxCombinations; i++ {
		var newCombinations []string
		for _, comb := range combinations {
			for i, r := range comb {
				if replacements, exists := p.CharMapping[r]; exists {
					for _, repl := range replacements {
						newComb := comb[:i] + string(repl) + comb[i+1:]
						newCombinations = append(newCombinations, newComb)
					}
				}
			}
		}
		combinations = append(combinations, newCombinations...)
	}
	return combinations
}
