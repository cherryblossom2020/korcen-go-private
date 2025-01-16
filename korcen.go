package korcen

import (
	"strings"
	"sync"

	"github.com/fluffy-melli/korcen-go/cache"
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

func CList(list *[]string, types int, input string, continues bool) (bool, []IndexOF) {
	indexs := make([]IndexOF, 0)
	loweredInput := strings.ToLower(input)
	resultChannel := make(chan []IndexOF, len(*list))
	var wg sync.WaitGroup
	chunkSize := len(*list) / 4
	if chunkSize == 0 {
		chunkSize = 1
	}
	for i := 0; i < len(*list); i += chunkSize {
		wg.Add(1)
		go func(startIdx int) {
			defer wg.Done()
			localIndexs := make([]IndexOF, 0)
			for j := startIdx; j < startIdx+chunkSize && j < len(*list); j++ {
				item := (*list)[j]
				in := strings.Index(input, item)
				if in != -1 {
					localIndexs = append(localIndexs, IndexOF{
						Swear: item,
						Type:  types,
						Start: in,
						End:   in + len(item),
					})
					if !continues {
						resultChannel <- localIndexs
						return
					}
				}
				itemEnglish := KoToEnglish(item)
				if len(itemEnglish) > 1 {
					in = strings.Index(loweredInput, strings.ToLower(itemEnglish))
					if in != -1 {
						localIndexs = append(localIndexs, IndexOF{
							Swear: EnglishToKo(itemEnglish),
							Type:  types,
							Start: in,
							End:   in + len(itemEnglish),
						})
						if !continues {
							resultChannel <- localIndexs
							return
						}
					}
				}
			}
			resultChannel <- localIndexs
		}(i)
	}
	wg.Wait()
	close(resultChannel)
	for localIndexs := range resultChannel {
		indexs = append(indexs, localIndexs...)
	}
	if len(indexs) > 0 {
		return true, indexs
	}
	return false, indexs
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
	return CList(cache.General, DGeneral, input, continues)
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
	return CList(cache.Minor, DMinor, input, continues)
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
	return CList(cache.Sexual, DSexual, input, continues)
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
	return CList(cache.Belittle, DBelittle, input, continues)
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
	return CList(cache.Race, DRace, input, continues)
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
	return CList(cache.Parent, DParent, input, continues)
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
	return CList(cache.Politics, DPolitics, input, continues)
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
	prof := NewProfanity(cache.English)
	df, pr := prof.Censor(input)
	if !df {
		return false, nil
	}
	in := strings.Index(input, pr)
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
	return CList(cache.Japanese, DJapanese, input, continues)
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
	return CList(cache.Chinese, DChinese, input, continues)
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
	return CList(cache.Emoji, DSpecial, input, continues)
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// 비속어 감지 및 결과 반환 함수
// 입력:
//
//	input: 비속어가 포함될 수 있는 문자열.
//
// 출력:
//
//	 *CheckInfo: struct {
//		    Detect bool   // 비속어 감지 여부
//			Swear  string // 감지된 비속어
//			Type   int    // 비속어의 유형
//	 }
func Check(input string) *CheckInfo {
	var detect bool
	var swear []IndexOF
	detect, swear = English(strings.ToLower(input), false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	input = ChangeUnicode(input)
	detect, swear = General(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Minor(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Sexual(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Belittle(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Race(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Parent(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Politics(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Japanese(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Chinese(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	detect, swear = Special(input, false)
	if detect {
		return &CheckInfo{
			Detect:  true,
			NewText: input,
			Swear:   swear,
		}
	}

	return &CheckInfo{
		Detect:  false,
		NewText: input,
		Swear:   nil,
	}
}
