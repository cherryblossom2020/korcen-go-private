package korcen

import (
	"strings"
	"sync"

	"github.com/fluffy-melli/korcen-go/cache"
)

type CheckInfo struct {
	Detect  bool
	NewText string
	Swear   []IndexOF
}

type IndexOF struct {
	Swear string
	Type  int
	Start int
	End   int
}

func InitProfanityData() {
	dataMap := map[*[]string]int{
		cache.General:  cache.DGeneral,
		cache.Minor:    cache.DMinor,
		cache.Sexual:   cache.DSexual,
		cache.Belittle: cache.DBelittle,
		cache.Race:     cache.DRace,
		cache.Parent:   cache.DParent,
		cache.Politics: cache.DPolitics,
		cache.English:  cache.DEnglish,
		cache.Japanese: cache.DJapanese,
		cache.Chinese:  cache.DChinese,
		cache.Emoji:    cache.DSpecial,
	}

	var wg sync.WaitGroup

	for listPtr, wtype := range dataMap {
		words := *listPtr
		wg.Add(1)
		go func(wordSlice []string, t int) {
			defer wg.Done()
			for _, w := range wordSlice {
				word := strings.ToLower(strings.TrimSpace(w))
				if word == "" {
					continue
				}
				cache.Mu.Lock()
				cache.ProfanityTrie.Insert(word)
				cache.BTree.ReplaceOrInsert(cache.WordEntry{Word: word, WType: t})
				cache.Mu.Unlock()
			}
		}(words, wtype)
	}

	wg.Wait()
}

func findCandidatesByTrie(input string) []string {
	words := strings.Fields(input)
	candidates := make([]string, 0)

	for _, w := range words {
		wLower := strings.ToLower(w)
		if cache.ProfanityTrie.Search(wLower) {
			candidates = append(candidates, wLower)
		}
	}
	return candidates
}

func Check(input string) *CheckInfo {
	newInput := strings.ToLower(input)
	candidates := findCandidatesByTrie(newInput)

	if len(candidates) == 0 {
		return &CheckInfo{
			Detect:  false,
			NewText: newInput,
			Swear:   nil,
		}
	}

	detected := false
	indexList := make([]IndexOF, 0, len(candidates))

	var wg sync.WaitGroup
	resultChan := make(chan IndexOF, len(candidates))

	for _, word := range candidates {
		wg.Add(1)
		go func(wd string) {
			defer wg.Done()
			cache.Mu.RLock()
			defer cache.Mu.RUnlock()

			found := cache.BTree.Get(cache.WordEntry{Word: wd})
			if found != nil {
				we := found.(cache.WordEntry)
				pos := strings.Index(newInput, wd)
				if pos != -1 {
					resultChan <- IndexOF{
						Swear: wd,
						Type:  we.WType,
						Start: pos,
						End:   pos + len(wd),
					}
				}
			}
		}(word)
	}

	wg.Wait()
	close(resultChan)

	for r := range resultChan {
		indexList = append(indexList, r)
		detected = true
	}

	return &CheckInfo{
		Detect:  detected,
		NewText: newInput,
		Swear:   indexList,
	}
}
