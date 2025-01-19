package cache

import "github.com/fluffy-melli/korcen-go/cache"

func PreWord() map[string]int {
	var cv map[string]int = make(map[string]int)
	for _, s := range *cache.English {
		cv[s] = len(cv)
	}
	for _, s := range *cache.General {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Minor {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Sexual {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Belittle {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Race {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Parent {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Politics {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Japanese {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Chinese {
		cv[s] = len(cv)
	}
	for _, s := range *cache.Emoji {
		cv[s] = len(cv)
	}
	return cv
}
