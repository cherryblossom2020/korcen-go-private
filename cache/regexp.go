package cache

import "regexp"

var KRinitial = map[string]string{
	"r": "ㄱ", "R": "ㄲ", "s": "ㄴ", "e": "ㄷ", "E": "ㄸ", "f": "ㄹ", "a": "ㅁ", "q": "ㅂ", "Q": "ㅃ",
	"t": "ㅅ", "T": "ㅆ", "d": "ㅇ", "w": "ㅈ", "W": "ㅉ", "c": "ㅊ", "z": "ㅋ", "x": "ㅌ", "v": "ㅍ", "g": "ㅎ",
}

var KRmedial = map[string]string{
	"k": "ㅏ", "o": "ㅐ", "i": "ㅑ", "I": "ㅒ", "j": "ㅓ", "p": "ㅔ", "u": "ㅕ", "P": "ㅖ", "h": "ㅗ", "y": "ㅛ",
	"n": "ㅜ", "b": "ㅠ", "m": "ㅡ", "l": "ㅣ",
}

var KRfinal = map[string]string{
	"r": "ㄱ", "s": "ㄴ", "e": "ㄷ", "f": "ㄹ", "a": "ㅁ", "q": "ㅂ", "t": "ㅅ", "d": "ㅇ", "w": "ㅈ", "c": "ㅊ",
	"z": "ㅋ", "x": "ㅌ", "v": "ㅍ", "g": "ㅎ",
}

var Regexinitial = regexp.MustCompile(`[rRsseEfFaqQtTdDwWcxyzv]`)
var Regexmedial = regexp.MustCompile(`[koijphyunbl]`)
var Regexfinal = regexp.MustCompile(`[rswacxzv]`)
