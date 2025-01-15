package korcen

import (
	"strings"
	"unicode"

	"github.com/fluffy-melli/korcen-go/cache"
	"github.com/gyarang/gohangul"
)

// Copyright© All rights reserved.
//  _____                 _
// |_   _|_ _ _ __   __ _| |_
//   | |/ _` | '_ \ / _` | __|
//   | | (_| | | | | (_| | |_
//   |_|\__,_|_| |_|\__,_|\__|

type CheckInfo struct {
	Detect  bool      // 비속어 감지 여부
	NewText string    // 입력된 메시지
	Swear   []IndexOF // 감지된 비속어
}

type IndexOF struct {
	Swear string // 감지된 비속어
	Type  int    // 감지된 비속어의 유형
	Start int    // 시작위치
	End   int    // 종료위치
}

// 비속어의 유형
const (
	DNone int = iota
	DGeneral
	DMinor
	DSexual
	DBelittle
	DRace
	DParent
	DPolitics
	DEnglish
	DJapanese
	DChinese
	DSpecial
)

func ChangeUnicode(unicode string) string {
	unicode = strings.ToLower(unicode)
	unicode = strings.ReplaceAll(unicode, "ㅿ", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "^", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "^", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "ハ", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "人", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "∧", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "／＼", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "/＼", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "㉦", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "甘", "ㅂ")
	unicode = strings.ReplaceAll(unicode, "廿", "ㅂ")
	unicode = strings.ReplaceAll(unicode, "己", "ㄹ")
	unicode = strings.ReplaceAll(unicode, "卜", "ㅏ")
	unicode = strings.ReplaceAll(unicode, "l", "ㅣ")
	unicode = strings.ReplaceAll(unicode, "r", "ㅏ")
	unicode = strings.ReplaceAll(unicode, "ト", "ㅏ")
	unicode = strings.ReplaceAll(unicode, "ᐲ", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "ᗨ", "ㅂ")
	unicode = strings.ReplaceAll(unicode, "𝗌", "s")
	unicode = strings.ReplaceAll(unicode, "𝗌", "s")
	unicode = strings.ReplaceAll(unicode, "𝘴", "s")
	unicode = strings.ReplaceAll(unicode, "𝙨", "s")
	unicode = strings.ReplaceAll(unicode, "𝚜", "s")
	unicode = strings.ReplaceAll(unicode, "𝐬", "s")
	unicode = strings.ReplaceAll(unicode, "𝑠", "s")
	unicode = strings.ReplaceAll(unicode, "𝒔", "s")
	unicode = strings.ReplaceAll(unicode, "𝓈", "s")
	unicode = strings.ReplaceAll(unicode, "𝓼", "s")
	unicode = strings.ReplaceAll(unicode, "𝔰", "s")
	unicode = strings.ReplaceAll(unicode, "𝖘", "s")
	unicode = strings.ReplaceAll(unicode, "𝕤", "s")
	unicode = strings.ReplaceAll(unicode, "ｓ", "s")
	unicode = strings.ReplaceAll(unicode, "ş", "s")
	unicode = strings.ReplaceAll(unicode, "ⓢ", "s")
	unicode = strings.ReplaceAll(unicode, "⒮", "s")
	unicode = strings.ReplaceAll(unicode, "🅢", "s")
	unicode = strings.ReplaceAll(unicode, "🆂", "s")
	unicode = strings.ReplaceAll(unicode, "🅂", "s")
	unicode = strings.ReplaceAll(unicode, "𝑺", "s")
	unicode = strings.ReplaceAll(unicode, "𝖾", "e")
	unicode = strings.ReplaceAll(unicode, "𝘦", "e")
	unicode = strings.ReplaceAll(unicode, "𝙚", "e")
	unicode = strings.ReplaceAll(unicode, "𝚎", "e")
	unicode = strings.ReplaceAll(unicode, "𝐞", "e")
	unicode = strings.ReplaceAll(unicode, "𝑒", "e")
	unicode = strings.ReplaceAll(unicode, "𝒆", "e")
	unicode = strings.ReplaceAll(unicode, "ℯ", "e")
	unicode = strings.ReplaceAll(unicode, "𝓮", "e")
	unicode = strings.ReplaceAll(unicode, "𝔢", "e")
	unicode = strings.ReplaceAll(unicode, "𝖊", "e")
	unicode = strings.ReplaceAll(unicode, "𝕖", "e")
	unicode = strings.ReplaceAll(unicode, "ｅ", "e")
	unicode = strings.ReplaceAll(unicode, "ė", "e")
	unicode = strings.ReplaceAll(unicode, "ⓔ", "e")
	unicode = strings.ReplaceAll(unicode, "⒠", "e")
	unicode = strings.ReplaceAll(unicode, "🅔", "e")
	unicode = strings.ReplaceAll(unicode, "🅴", "e")
	unicode = strings.ReplaceAll(unicode, "🄴", "e")
	unicode = strings.ReplaceAll(unicode, "є", "e")
	unicode = strings.ReplaceAll(unicode, "𝗑", "x")
	unicode = strings.ReplaceAll(unicode, "𝘹", "x")
	unicode = strings.ReplaceAll(unicode, "𝙭", "x")
	unicode = strings.ReplaceAll(unicode, "𝚡", "x")
	unicode = strings.ReplaceAll(unicode, "𝐱", "x")
	unicode = strings.ReplaceAll(unicode, "𝑥", "x")
	unicode = strings.ReplaceAll(unicode, "𝒙", "x")
	unicode = strings.ReplaceAll(unicode, "𝓍", "x")
	unicode = strings.ReplaceAll(unicode, "𝔁", "x")
	unicode = strings.ReplaceAll(unicode, "𝔵", "x")
	unicode = strings.ReplaceAll(unicode, "𝖝", "x")
	unicode = strings.ReplaceAll(unicode, "𝕩", "x")
	unicode = strings.ReplaceAll(unicode, "ｘ", "x")
	unicode = strings.ReplaceAll(unicode, "ⓧ", "x")
	unicode = strings.ReplaceAll(unicode, "⒳", "x")
	unicode = strings.ReplaceAll(unicode, "🅧", "x")
	unicode = strings.ReplaceAll(unicode, "🆇", "x")
	unicode = strings.ReplaceAll(unicode, "🅇", "x")
	unicode = strings.ReplaceAll(unicode, "₨", "rs")
	unicode = strings.ReplaceAll(unicode, "ų", "u")
	unicode = strings.ReplaceAll(unicode, "ç", "c")
	unicode = strings.ReplaceAll(unicode, "Ｆ", "F")
	unicode = strings.ReplaceAll(unicode, "Ｋ", "K")
	unicode = strings.ReplaceAll(unicode, "Ｃ", "C")
	unicode = strings.ReplaceAll(unicode, "Ｕ", "U")
	unicode = strings.ReplaceAll(unicode, "し ", "ㅣ")
	unicode = strings.ReplaceAll(unicode, "i ", "ㅣ")
	unicode = strings.ReplaceAll(unicode, "l ", "ㅣ")
	unicode = strings.ReplaceAll(unicode, "の ", "ㅇ")
	unicode = strings.ReplaceAll(unicode, "ㅁ ", "ㅁ")
	unicode = strings.ReplaceAll(unicode, "口 ", "ㅁ")
	unicode = strings.ReplaceAll(unicode, "曰 ", "ㅁ")
	unicode = strings.ReplaceAll(unicode, "H", "ㅐ")
	unicode = strings.ReplaceAll(unicode, "H", "ㅐ")
	unicode = strings.ReplaceAll(unicode, "ス", "ㅈ")
	unicode = strings.ReplaceAll(unicode, "へ", "ㅅ")
	unicode = strings.ReplaceAll(unicode, "旦", "므")
	unicode = strings.ReplaceAll(unicode, "畀", "뷰")
	unicode = strings.ReplaceAll(unicode, "%", "응")
	unicode = strings.ReplaceAll(unicode, "_", "ㅡ")
	unicode = strings.ReplaceAll(unicode, "-", "ㅡ")
	unicode = strings.ReplaceAll(unicode, "/", "ㅣ")
	unicode = EtoK(unicode)
	unicode = Clean(unicode)
	dism := gohangul.Disassemble(unicode)
	unicode = gohangul.Assemble(dism)
	unicode = RpGeneral(unicode)
	unicode = RpMinor(unicode)
	unicode = RpSexual(unicode)
	unicode = RpBelittle(unicode)
	unicode = RpPolitics(unicode)
	unicode = DEL_J(unicode)
	unicode = After(unicode)
	return unicode
}

func After(input string) string {
	input = strings.ReplaceAll(input, "자아", "자")
	input = strings.ReplaceAll(input, "가아", "가")
	input = strings.ReplaceAll(input, "나아", "나")
	input = strings.ReplaceAll(input, "세엑", "섹")
	input = strings.ReplaceAll(input, "세에", "섹")
	return input
}

func Clean(input string) string {
	var result strings.Builder
	var prevRune rune
	for i, currentRune := range input {
		if (i == 0 || currentRune != prevRune) && (unicode.IsLetter(currentRune) || unicode.IsDigit(currentRune) || unicode.IsSpace(currentRune)) {
			result.WriteRune(currentRune)
		}
		prevRune = currentRune
	}
	return result.String()
}

func DEL_J(input string) string {
	input = strings.ReplaceAll(input, "ㅏ", "")
	input = strings.ReplaceAll(input, "ㅑ", "")
	input = strings.ReplaceAll(input, "ㅓ", "")
	input = strings.ReplaceAll(input, "ㅕ", "")
	input = strings.ReplaceAll(input, "ㅗ", "")
	input = strings.ReplaceAll(input, "ㅛ", "")
	input = strings.ReplaceAll(input, "ㅜ", "")
	input = strings.ReplaceAll(input, "ㅠ", "")
	input = strings.ReplaceAll(input, "ㅡ", "")
	input = strings.ReplaceAll(input, "ㅣ", "")
	input = strings.ReplaceAll(input, "ㅐ", "")
	input = strings.ReplaceAll(input, "ㅔ", "")
	input = strings.ReplaceAll(input, "ㅘ", "")
	input = strings.ReplaceAll(input, "ㅙ", "")
	input = strings.ReplaceAll(input, "ㅚ", "")
	input = strings.ReplaceAll(input, "ㅝ", "")
	input = strings.ReplaceAll(input, "ㅞ", "")
	input = strings.ReplaceAll(input, "ㅟ", "")
	input = strings.ReplaceAll(input, "ㅢ", "")
	input = strings.ReplaceAll(input, "ㆎ", "")
	return input
}

func EtoK(input string) string {
	input = strings.ReplaceAll(input, "a", "ㅁ")
	input = strings.ReplaceAll(input, "b", "ㅠ")
	input = strings.ReplaceAll(input, "c", "ㅊ")
	input = strings.ReplaceAll(input, "d", "ㅇ")
	input = strings.ReplaceAll(input, "e", "ㄷ")
	input = strings.ReplaceAll(input, "f", "ㄹ")
	input = strings.ReplaceAll(input, "g", "ㅎ")
	input = strings.ReplaceAll(input, "h", "ㅗ")
	input = strings.ReplaceAll(input, "i", "ㅑ")
	input = strings.ReplaceAll(input, "j", "ㅓ")
	input = strings.ReplaceAll(input, "k", "ㅏ")
	input = strings.ReplaceAll(input, "l", "ㅣ")
	input = strings.ReplaceAll(input, "m", "ㅡ")
	input = strings.ReplaceAll(input, "n", "ㅜ")
	input = strings.ReplaceAll(input, "o", "ㅐ")
	input = strings.ReplaceAll(input, "p", "ㅔ")
	input = strings.ReplaceAll(input, "q", "ㅂ")
	input = strings.ReplaceAll(input, "r", "ㄱ")
	input = strings.ReplaceAll(input, "s", "ㄴ")
	input = strings.ReplaceAll(input, "t", "ㅅ")
	input = strings.ReplaceAll(input, "u", "ㅕ")
	input = strings.ReplaceAll(input, "v", "ㅍ")
	input = strings.ReplaceAll(input, "w", "ㅈ")
	input = strings.ReplaceAll(input, "x", "ㅌ")
	input = strings.ReplaceAll(input, "y", "ㅛ")
	input = strings.ReplaceAll(input, "z", "ㅋ")
	return input
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 일반적인 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func General(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DGeneral,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 사소한 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Minor(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DMinor,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 성적관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Sexual(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DSexual,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 비하발언관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Belittle(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DBelittle,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 인종차별관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Race(input string, continues bool) (bool, []IndexOF) {
	input = strings.ReplaceAll(input, "흑형님", "")
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DRace,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 패드립관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Parent(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DParent,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 정치관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Politics(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DPolitics,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 영어관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func English(input string, continues bool) (bool, []IndexOF) {
	newtext := ChangeUnicode(input)
	prof := NewProfanity(cache.English)
	df, pr := prof.Censor(newtext)
	if !df {
		return false, make([]IndexOF, 0)
	}
	in := strings.Index(newtext, pr)
	return true, []IndexOF{
		{
			Swear: pr,
			Type:  DEnglish,
			Start: in,
			End:   in + len(pr),
		},
	}
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 일본어관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Japanese(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DJapanese,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 중국어관련 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Chinese(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DChinese,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	bool: 비속어가 포함된 경우 true, 그렇지 않으면 false.
//	string: 감지된 비속어가 있으면 해당 비속어를, 없으면 빈 문자열("")을 반환.
func Special(input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	for _, item := range cache.General {
		in := strings.Index(input, item)
		if in != -1 {
			indexs = append(indexs, IndexOF{
				Swear: item,
				Type:  DSpecial,
				Start: in,
				End:   in + len(item),
			})
			if !continues {
				return true, indexs
			}
		}
	}

	return false, indexs
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	 CheckInfo: struct {
//		    Detect bool   // 비속어 감지 여부
//			Swear  string // 감지된 비속어
//			Type   int    // 비속어의 유형
//	 }
func Check(input string) CheckInfo {
	var detect bool
	var swear []IndexOF
	input = ChangeUnicode(input)
	detect, swear = General(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Minor(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Sexual(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Belittle(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Race(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Parent(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Politics(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = English(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Japanese(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Chinese(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Special(input, false)
	if detect {
		return CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	return CheckInfo{
		Detect: false,
		Swear:  make([]IndexOF, 0),
	}
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	 CheckInfo: struct {
//		    Detect bool   // 비속어 감지 여부
//			Swear  string // 감지된 비속어
//			Type   int    // 비속어의 유형
//	 }
func AllCheck(input string) CheckInfo {
	var swear []IndexOF
	var respond []IndexOF
	input = ChangeUnicode(input)
	_, swear = General(input, true)
	respond = append(respond, swear...)

	_, swear = Minor(input, true)
	respond = append(respond, swear...)

	_, swear = Sexual(input, true)
	respond = append(respond, swear...)

	_, swear = Belittle(input, true)
	respond = append(respond, swear...)

	_, swear = Race(input, true)
	respond = append(respond, swear...)

	_, swear = Parent(input, true)
	respond = append(respond, swear...)

	_, swear = Politics(input, true)
	respond = append(respond, swear...)

	_, swear = English(input, true)
	respond = append(respond, swear...)

	_, swear = Japanese(input, true)
	respond = append(respond, swear...)

	_, swear = Chinese(input, true)
	respond = append(respond, swear...)

	_, swear = Special(input, true)
	respond = append(respond, swear...)

	return CheckInfo{
		Detect:  len(respond) != 0,
		NewText: input,
		Swear:   respond,
	}
}
