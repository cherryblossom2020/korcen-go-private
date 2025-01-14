package korcen

import (
	"strings"
	"unicode"

	"github.com/fluffy-melli/korcen-go/cache"
)

// Copyright© All rights reserved.
//  _____                 _
// |_   _|_ _ _ __   __ _| |_
//   | |/ _` | '_ \ / _` | __|
//   | | (_| | | | | (_| | |_
//   |_|\__,_|_| |_|\__,_|\__|

type CheckInfo struct {
	Detect bool   // 비속어 감지 여부
	Swear  string // 감지된 비속어
	Type   int    // 비속어의 유형
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
	unicode = strings.ReplaceAll(unicode, ".", "")
	unicode = strings.ReplaceAll(unicode, "^", "ㅅ")
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
	unicode = strings.ToLower(unicode)
	unicode = Duplicate(unicode)
	unicode = SpecialChar(unicode)
	return unicode
}

func Duplicate(input string) string {
	var result strings.Builder
	var prevRune rune
	for i, currentRune := range input {
		if i == 0 || currentRune != prevRune {
			result.WriteRune(currentRune)
		}
		prevRune = currentRune
	}
	return result.String()
}

func SpecialChar(input string) string {
	var result []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return string(result)
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
func General(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "ㅗ먹어", "ㅗ")
	input = strings.ReplaceAll(input, "오ㅗㅗ", "")
	input = strings.ReplaceAll(input, "오ㅗ", "")
	input = strings.ReplaceAll(input, "해ㅗㅗ", "")
	input = strings.ReplaceAll(input, "해ㅗ", "")
	input = strings.ReplaceAll(input, "호ㅗㅗ", "")
	input = strings.ReplaceAll(input, "호ㅗ", "")
	input = strings.ReplaceAll(input, "로ㅗㅗ", "")
	input = strings.ReplaceAll(input, "로ㅗ", "")
	input = strings.ReplaceAll(input, "옹ㅗㅗ", "")
	input = strings.ReplaceAll(input, "옹ㅗ", "")
	input = strings.ReplaceAll(input, "롤ㅗㅗ", "")
	input = strings.ReplaceAll(input, "롤ㅗ", "")
	input = strings.ReplaceAll(input, "요ㅗ", "")
	input = strings.ReplaceAll(input, "우ㅗ", "")
	input = strings.ReplaceAll(input, "하ㅗ", "")
	input = strings.ReplaceAll(input, "ㅗㅗ오", "")
	input = strings.ReplaceAll(input, "ㅗ오", "")
	input = strings.ReplaceAll(input, "ㅗㅗ호", "")
	input = strings.ReplaceAll(input, "ㅗ호", "")
	input = strings.ReplaceAll(input, "ㅗㅗ로", "")
	input = strings.ReplaceAll(input, "ㅗ로", "")
	input = strings.ReplaceAll(input, "ㅗㅗ옹", "")
	input = strings.ReplaceAll(input, "ㅗ옹", "")
	input = strings.ReplaceAll(input, "ㅗㅗ롤", "")
	input = strings.ReplaceAll(input, "ㅗ롤", "")
	input = strings.ReplaceAll(input, "ㅗ요", "")
	input = strings.ReplaceAll(input, "ㅗ우", "")
	input = strings.ReplaceAll(input, "ㅗ하", "")
	input = strings.ReplaceAll(input, "오ㅗㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "오ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "호ㅗㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "호ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "로ㅗㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "로ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "옹ㅗㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "옹ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "롤ㅗㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "롤ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "요ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "우ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "하ㅗㅗㅗ", "")
	input = strings.ReplaceAll(input, "ㅇㅗ", "")
	input = strings.ReplaceAll(input, "ㅗㄷ", "")
	input = strings.ReplaceAll(input, "ㅗㅜ", "")
	input = strings.ReplaceAll(input, "rㅗ", "")
	input = strings.ReplaceAll(input, "ㅗr", "")
	input = strings.ReplaceAll(input, "sㅗ", "")
	input = strings.ReplaceAll(input, "ㅗs", "")
	input = strings.ReplaceAll(input, "eㅗ", "")
	input = strings.ReplaceAll(input, "ㅗe", "")
	input = strings.ReplaceAll(input, "fㅗ", "")
	input = strings.ReplaceAll(input, "ㅗf", "")
	input = strings.ReplaceAll(input, "aㅗ", "")
	input = strings.ReplaceAll(input, "ㅗa", "")
	input = strings.ReplaceAll(input, "qㅗ", "")
	input = strings.ReplaceAll(input, "ㅗq", "")
	input = strings.ReplaceAll(input, "ㅗw", "")
	input = strings.ReplaceAll(input, "wㅗ", "")
	input = strings.ReplaceAll(input, "ㅗd", "")
	input = strings.ReplaceAll(input, "dㅗ", "")
	input = strings.ReplaceAll(input, "ㅗg", "")
	input = strings.ReplaceAll(input, "gㅗ", "")
	input = strings.ReplaceAll(input, "dㅗ", "")
	input = strings.ReplaceAll(input, "118", "")
	input = strings.ReplaceAll(input, "218", "")
	input = strings.ReplaceAll(input, "318", "")
	input = strings.ReplaceAll(input, "418", "")
	input = strings.ReplaceAll(input, "518", "")
	input = strings.ReplaceAll(input, "618", "")
	input = strings.ReplaceAll(input, "718", "")
	input = strings.ReplaceAll(input, "818", "")
	input = strings.ReplaceAll(input, "918", "")
	input = strings.ReplaceAll(input, "018", "")
	input = strings.ReplaceAll(input, "련", "놈")
	input = strings.ReplaceAll(input, "뇬", "놈")
	input = strings.ReplaceAll(input, "놈", "놈")
	input = strings.ReplaceAll(input, "넘", "놈")
	input = strings.ReplaceAll(input, "8분", "")
	input = strings.ReplaceAll(input, "^", "ㅅ")
	input = strings.ReplaceAll(input, "人", "ㅅ")
	input = strings.ReplaceAll(input, "∧", "ㅅ")
	input = strings.ReplaceAll(input, "／＼", "ㅅ")
	input = strings.ReplaceAll(input, "/＼", "ㅅ")
	input = strings.ReplaceAll(input, "㉦", "ㅅ")
	input = strings.ReplaceAll(input, "丨발", "시발")
	input = strings.ReplaceAll(input, "丨벌", "시발")
	input = strings.ReplaceAll(input, "丨바", "시발")
	input = strings.ReplaceAll(input, "甘", "ㅂ")
	input = strings.ReplaceAll(input, "廿", "ㅂ")
	input = strings.ReplaceAll(input, "己", "ㄹ")
	input = strings.ReplaceAll(input, "卜", "ㅏ")
	input = strings.ReplaceAll(input, "l", "ㅣ")
	input = strings.ReplaceAll(input, "r", "ㅏ")
	input = strings.ReplaceAll(input, "ᐲ", "ㅅ")
	input = strings.ReplaceAll(input, "ᗨ", "ㅂ")
	input = strings.ReplaceAll(input, "시ㅣ", "시")
	input = strings.ReplaceAll(input, "씨ㅣ", "씨")
	input = strings.ReplaceAll(input, "ㅅ1", "시")
	input = strings.ReplaceAll(input, "ㅍㅅㅍ", "")
	input = strings.ReplaceAll(input, "발 닦", "")
	input = strings.ReplaceAll(input, "동시 8", "")
	input = strings.ReplaceAll(input, "8시발", "시발")
	input = strings.ReplaceAll(input, "8시 ", "")
	input = strings.ReplaceAll(input, "!wt ㅂㅇ", "")
	input = strings.ReplaceAll(input, "!ㅈㅅ ㅂㅇ", "")
	input = strings.ReplaceAll(input, "다시 방", "")
	input = strings.ReplaceAll(input, "시 불이익", "")
	input = strings.ReplaceAll(input, "다시 바꿀", "")
	input = strings.ReplaceAll(input, "다시 바꿔", "")
	input = strings.ReplaceAll(input, "다시 불러", "")
	input = strings.ReplaceAll(input, "다시 불안", "")
	input = strings.ReplaceAll(input, "하시바라 이노스케", "")
	input = strings.ReplaceAll(input, "할 시", "")
	input = strings.ReplaceAll(input, "불러드", "")
	input = strings.ReplaceAll(input, "시발음", "")
	input = strings.ReplaceAll(input, "시발택시", "")
	input = strings.ReplaceAll(input, "시발자동차", "")
	input = strings.ReplaceAll(input, "정치발", "")
	input = strings.ReplaceAll(input, "시발점", "")
	input = strings.ReplaceAll(input, "시발유", "")
	input = strings.ReplaceAll(input, "시발역", "")
	input = strings.ReplaceAll(input, "시발수뢰", "")
	input = strings.ReplaceAll(input, "아저씨바", "")
	input = strings.ReplaceAll(input, "아저씨발", "")
	input = strings.ReplaceAll(input, "시바견", "")
	input = strings.ReplaceAll(input, "벌어", "")
	input = strings.ReplaceAll(input, "시바이누", "")
	input = strings.ReplaceAll(input, "시바스리갈", "")
	input = strings.ReplaceAll(input, "시바산", "")
	input = strings.ReplaceAll(input, "시바신", "")
	input = strings.ReplaceAll(input, "오리발", "")
	input = strings.ReplaceAll(input, "발끝", "")
	input = strings.ReplaceAll(input, "다시바", "")
	input = strings.ReplaceAll(input, "다시팔", "")
	input = strings.ReplaceAll(input, "비슈누시바", "")
	input = strings.ReplaceAll(input, "시바핫카이", "")
	input = strings.ReplaceAll(input, "시바타이쥬", "")
	input = strings.ReplaceAll(input, "데스티니시바", "")
	input = strings.ReplaceAll(input, "시바루", "")
	input = strings.ReplaceAll(input, "시바료타로", "")
	input = strings.ReplaceAll(input, "시바라스시", "")
	input = strings.ReplaceAll(input, "임시방편", "")
	input = strings.ReplaceAll(input, "젤리", "")
	input = strings.ReplaceAll(input, "발사", "")
	input = strings.ReplaceAll(input, "크시야", "")
	input = strings.ReplaceAll(input, "크시", "")
	input = strings.ReplaceAll(input, "어찌", "")
	input = strings.ReplaceAll(input, "가시방석", "")
	input = strings.ReplaceAll(input, "발로란트방", "")
	input = strings.ReplaceAll(input, "발로란트", "")
	input = strings.ReplaceAll(input, "발로", "")
	input = strings.ReplaceAll(input, "씨발라", "")
	input = strings.ReplaceAll(input, "무시발언", "")
	input = strings.ReplaceAll(input, "일시불", "")
	input = strings.ReplaceAll(input, "우리", "")
	input = strings.ReplaceAll(input, "혹시", "")
	input = strings.ReplaceAll(input, "아조씨", "")
	input = strings.ReplaceAll(input, "아저씨", "")
	input = strings.ReplaceAll(input, "바로", "")
	input = strings.ReplaceAll(input, "저거시", "")
	input = strings.ReplaceAll(input, "우리발", "")
	input = strings.ReplaceAll(input, "피시방", "")
	input = strings.ReplaceAll(input, "피씨방", "")
	input = strings.ReplaceAll(input, "방장", "")
	input = strings.ReplaceAll(input, "시바사키", "")
	input = strings.ReplaceAll(input, "시발차", "")
	input = strings.ReplaceAll(input, "구로역 시발", "")
	input = strings.ReplaceAll(input, "로벅스", "")
	input = strings.ReplaceAll(input, "쉬바나", "")
	input = strings.ReplaceAll(input, "벌었는데", "")
	input = strings.ReplaceAll(input, "엠씨방", "")
	input = strings.ReplaceAll(input, "빨리", "")
	input = strings.ReplaceAll(input, "파엠", "")
	input = strings.ReplaceAll(input, "벌금", "")
	input = strings.ReplaceAll(input, "시방향", "")
	input = strings.ReplaceAll(input, "불법", "")
	input = strings.ReplaceAll(input, "발릴", "")
	input = strings.ReplaceAll(input, "발표", "")
	input = strings.ReplaceAll(input, "방송", "")
	input = strings.ReplaceAll(input, "역시", "")
	input = strings.ReplaceAll(input, "바보", "")
	input = strings.ReplaceAll(input, "쿨리발리", "")
	input = strings.ReplaceAll(input, "슈발리에", "")
	input = strings.ReplaceAll(input, "방탄", "")
	input = strings.ReplaceAll(input, "방어", "")
	input = strings.ReplaceAll(input, "발표", "")
	input = strings.ReplaceAll(input, "상시", "")
	input = strings.ReplaceAll(input, "슈발리에", "")
	input = strings.ReplaceAll(input, "아", "")
	input = strings.ReplaceAll(input, "이", "")
	input = strings.ReplaceAll(input, "일", "")
	input = strings.ReplaceAll(input, "의", "")
	input = strings.ReplaceAll(input, "즉시발급", "")
	input = strings.ReplaceAll(input, "련", "놈")
	input = strings.ReplaceAll(input, "뇬", "놈")
	input = strings.ReplaceAll(input, "놈", "놈")
	input = strings.ReplaceAll(input, "넘", "놈")
	input = strings.ReplaceAll(input, "련", "놈")
	input = strings.ReplaceAll(input, "뇬", "놈")
	input = strings.ReplaceAll(input, "놈", "놈")
	input = strings.ReplaceAll(input, "넘", "놈")
	input = strings.ReplaceAll(input, "opgg", "")
	input = strings.ReplaceAll(input, "op.gg", "")
	input = strings.ReplaceAll(input, "g랄", "지랄")
	input = strings.ReplaceAll(input, "ji랄", "지랄")
	input = strings.ReplaceAll(input, "己", "ㄹ")
	input = strings.ReplaceAll(input, "있지", "")
	input = strings.ReplaceAll(input, "없지", "")
	input = strings.ReplaceAll(input, "하지", "")
	input = strings.ReplaceAll(input, "알았지", "")
	input = strings.ReplaceAll(input, "몰랐지", "")
	input = strings.ReplaceAll(input, "근데", "")
	input = strings.ReplaceAll(input, "지근거", "")
	input = strings.ReplaceAll(input, "지근하", "")
	input = strings.ReplaceAll(input, "지근지근", "")
	input = strings.ReplaceAll(input, "지근속근", "")
	input = strings.ReplaceAll(input, "속든지근", "")
	input = strings.ReplaceAll(input, "미지근", "")
	input = strings.ReplaceAll(input, "근", "ㄹ")
	input = strings.ReplaceAll(input, "ㄹㅇ", "")
	input = strings.ReplaceAll(input, "있지", "")
	input = strings.ReplaceAll(input, "없지", "")
	input = strings.ReplaceAll(input, "하지", "")
	input = strings.ReplaceAll(input, "지랄탄", "")
	input = strings.ReplaceAll(input, "지랄버릇", "")
	input = strings.ReplaceAll(input, "이", "")
	input = strings.ReplaceAll(input, "알았지", "")
	input = strings.ReplaceAll(input, "몰랐지", "")
	input = strings.ReplaceAll(input, "근데", "")
	input = strings.ReplaceAll(input, "0등신", "")
	input = strings.ReplaceAll(input, "1등신", "")
	input = strings.ReplaceAll(input, "2등신", "")
	input = strings.ReplaceAll(input, "3등신", "")
	input = strings.ReplaceAll(input, "4등신", "")
	input = strings.ReplaceAll(input, "5등신", "")
	input = strings.ReplaceAll(input, "6등신", "")
	input = strings.ReplaceAll(input, "7등신", "")
	input = strings.ReplaceAll(input, "8등신", "")
	input = strings.ReplaceAll(input, "9등신", "")
	input = strings.ReplaceAll(input, "빙", "병")
	input = strings.ReplaceAll(input, "븅", "병")
	input = strings.ReplaceAll(input, "등", "병")
	input = strings.ReplaceAll(input, "붱", "병")
	input = strings.ReplaceAll(input, "뵝", "병")
	input = strings.ReplaceAll(input, "뼝", "병")
	input = strings.ReplaceAll(input, "싄", "신")
	input = strings.ReplaceAll(input, "씬", "신")
	input = strings.ReplaceAll(input, "우", "")
	input = strings.ReplaceAll(input, "웅", "")
	input = strings.ReplaceAll(input, "융", "")
	input = strings.ReplaceAll(input, "유", "")
	input = strings.ReplaceAll(input, "영", "")
	input = strings.ReplaceAll(input, "엉", "")
	input = strings.ReplaceAll(input, "전염병", "")
	input = strings.ReplaceAll(input, "감염병", "")
	input = strings.ReplaceAll(input, "화염병", "")
	input = strings.ReplaceAll(input, "왜꺼져", "")
	input = strings.ReplaceAll(input, "꺼져요", "")
	input = strings.ReplaceAll(input, "이꺼져", "")
	input = strings.ReplaceAll(input, "꺼져서", "")
	input = strings.ReplaceAll(input, "내꺼져", "")
	input = strings.ReplaceAll(input, "제꺼져", "")
	input = strings.ReplaceAll(input, "꺼져있", "")
	input = strings.ReplaceAll(input, "꺼져잇", "")
	input = strings.ReplaceAll(input, "꺼져도", "")
	input = strings.ReplaceAll(input, "계속꺼져", "")
	input = strings.ReplaceAll(input, "꺼져가", "")
	input = strings.ReplaceAll(input, "^", "ㅅ")
	input = strings.ReplaceAll(input, "H", "ㅐ")
	input = strings.ReplaceAll(input, "새로", "")
	input = strings.ReplaceAll(input, "77", "ㄲ")
	input = strings.ReplaceAll(input, "l", "ㅣ")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "10새", "새끼")
	input = strings.ReplaceAll(input, "10섹", "새끼")
	input = strings.ReplaceAll(input, "10쌔", "새끼")
	input = strings.ReplaceAll(input, "10쎄", "새끼")
	input = strings.ReplaceAll(input, "10새", "새끼")
	input = strings.ReplaceAll(input, "10쉑", "새끼")
	input = strings.ReplaceAll(input, "의새끼", "")
	input = strings.ReplaceAll(input, "애", "")
	input = strings.ReplaceAll(input, "에", "")
	input = strings.ReplaceAll(input, "루세끼", "")
	input = strings.ReplaceAll(input, "시세끼", "")
	input = strings.ReplaceAll(input, "세끼먹", "")
	input = strings.ReplaceAll(input, "고양이새끼", "")
	input = strings.ReplaceAll(input, "호랑이새끼", "")
	input = strings.ReplaceAll(input, "용새끼", "")
	input = strings.ReplaceAll(input, "말새끼", "")
	input = strings.ReplaceAll(input, "사자새끼", "")
	input = strings.ReplaceAll(input, "범새끼", "")
	input = strings.ReplaceAll(input, "삵새끼", "")
	input = strings.ReplaceAll(input, "키보드", "")
	input = strings.ReplaceAll(input, "새끼손", "")
	input = strings.ReplaceAll(input, "셰리프", "")
	input = strings.ReplaceAll(input, "로쉐리", "")
	input = strings.ReplaceAll(input, "丕", "조")
	input = strings.ReplaceAll(input, "刀卜", "까")
	input = strings.ReplaceAll(input, "줫습니다", "")
	input = strings.ReplaceAll(input, "줫음", "")
	input = strings.ReplaceAll(input, "줫잖아", "")
	input = strings.ReplaceAll(input, "줫겠지", "")
	input = strings.ReplaceAll(input, "쫒아", "")
	input = strings.ReplaceAll(input, "쫒는", "")
	input = strings.ReplaceAll(input, "쫒기다", "")
	input = strings.ReplaceAll(input, "쫒기라", "")
	input = strings.ReplaceAll(input, "쫒기로", "")
	input = strings.ReplaceAll(input, "쫒기를", "")
	input = strings.ReplaceAll(input, "쫒기며", "")
	input = strings.ReplaceAll(input, "쫒기는", "")
	input = strings.ReplaceAll(input, "쫒기나", "")
	input = strings.ReplaceAll(input, "쫒겨", "")
	input = strings.ReplaceAll(input, "쫒겻", "")
	input = strings.ReplaceAll(input, "쫒겼", "")
	input = strings.ReplaceAll(input, "쫒았", "")
	input = strings.ReplaceAll(input, "쫒다", "")
	input = strings.ReplaceAll(input, "쫒고", "")
	input = strings.ReplaceAll(input, "줫는", "")
	input = strings.ReplaceAll(input, "줫어", "")
	input = strings.ReplaceAll(input, "줬는", "")
	input = strings.ReplaceAll(input, "줫군", "")
	input = strings.ReplaceAll(input, "줬다", "")
	input = strings.ReplaceAll(input, "줬어", "")
	input = strings.ReplaceAll(input, "천조", "")
	input = strings.ReplaceAll(input, "쫒기", "")
	input = strings.ReplaceAll(input, "해줫더니", "")
	input = strings.ReplaceAll(input, "줫다", "")
	input = strings.ReplaceAll(input, "내쫒은", "")
	input = strings.ReplaceAll(input, "내쫒다", "")
	input = strings.ReplaceAll(input, "좇아", "")
	input = strings.ReplaceAll(input, "0개", "")
	input = strings.ReplaceAll(input, "1개", "")
	input = strings.ReplaceAll(input, "2개", "")
	input = strings.ReplaceAll(input, "3개", "")
	input = strings.ReplaceAll(input, "4개", "")
	input = strings.ReplaceAll(input, "5개", "")
	input = strings.ReplaceAll(input, "6개", "")
	input = strings.ReplaceAll(input, "7개", "")
	input = strings.ReplaceAll(input, "8개", "")
	input = strings.ReplaceAll(input, "9개", "")
	input = strings.ReplaceAll(input, "0개", "")
	input = strings.ReplaceAll(input, "1년", "")
	input = strings.ReplaceAll(input, "2년", "")
	input = strings.ReplaceAll(input, "3년", "")
	input = strings.ReplaceAll(input, "4년", "")
	input = strings.ReplaceAll(input, "5년", "")
	input = strings.ReplaceAll(input, "6년", "")
	input = strings.ReplaceAll(input, "7년", "")
	input = strings.ReplaceAll(input, "8년", "")
	input = strings.ReplaceAll(input, "9년", "")
	input = strings.ReplaceAll(input, "🐕", "개")
	input = strings.ReplaceAll(input, "🐦", "새")
	input = strings.ReplaceAll(input, "재밌게 놈", "")
	input = strings.ReplaceAll(input, "있게", "")
	input = strings.ReplaceAll(input, "년생", "")
	input = strings.ReplaceAll(input, "무지개색", "")
	input = strings.ReplaceAll(input, "떠돌이개", "")
	input = strings.ReplaceAll(input, "에게", "")
	input = strings.ReplaceAll(input, "넘는", "")
	input = strings.ReplaceAll(input, "소개", "")
	input = strings.ReplaceAll(input, "생긴게", "")
	input = strings.ReplaceAll(input, "날개같다", "")
	for _, item := range cache.General {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Minor(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "년", "놈")
	input = strings.ReplaceAll(input, "련", "놈")
	input = strings.ReplaceAll(input, "거미", "")
	input = strings.ReplaceAll(input, "친구", "")
	input = strings.ReplaceAll(input, "개미", "")
	input = strings.ReplaceAll(input, "이미친", "")
	input = strings.ReplaceAll(input, "미친증", "")
	input = strings.ReplaceAll(input, "동그라미", "")
	input = strings.ReplaceAll(input, "미틴", "미친")
	input = strings.ReplaceAll(input, "년", "놈")
	input = strings.ReplaceAll(input, "련", "놈")
	input = strings.ReplaceAll(input, "뒤져봐야", "")
	input = strings.ReplaceAll(input, "뒤질뻔", "")
	input = strings.ReplaceAll(input, "뒤져보다", "")
	input = strings.ReplaceAll(input, "뒤져보는", "")
	input = strings.ReplaceAll(input, "뒤져보고", "")
	input = strings.ReplaceAll(input, "뒤져간다", "")
	input = strings.ReplaceAll(input, "뒤져서", "")
	input = strings.ReplaceAll(input, "뒤져본", "")
	input = strings.ReplaceAll(input, "뒤져봄", "")
	input = strings.ReplaceAll(input, "뒤져볼", "")
	for _, item := range cache.Minor {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Sexual(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "보지도 못", "")
	input = strings.ReplaceAll(input, "보지도 않", "")
	input = strings.ReplaceAll(input, "인가 보지", "")
	input = strings.ReplaceAll(input, "면접 보지", "")
	input = strings.ReplaceAll(input, "영화 보지", "")
	input = strings.ReplaceAll(input, "애니 보지", "")
	input = strings.ReplaceAll(input, "만화 보지", "")
	input = strings.ReplaceAll(input, "사진 보지", "")
	input = strings.ReplaceAll(input, "그림 보지", "")
	input = strings.ReplaceAll(input, "면접보지", "")
	input = strings.ReplaceAll(input, "영화보지", "")
	input = strings.ReplaceAll(input, "애니보지", "")
	input = strings.ReplaceAll(input, "만화보지", "")
	input = strings.ReplaceAll(input, "사진보지", "")
	input = strings.ReplaceAll(input, "그림보지", "")
	input = strings.ReplaceAll(input, "을 보지", "")
	input = strings.ReplaceAll(input, "나 보지", "")
	input = strings.ReplaceAll(input, "못 보지", "")
	input = strings.ReplaceAll(input, "안 보지", "")
	input = strings.ReplaceAll(input, "왜 보지", "")
	input = strings.ReplaceAll(input, "뭐 보지", "")
	input = strings.ReplaceAll(input, "다 보지", "")
	input = strings.ReplaceAll(input, "빨리 보지", "")
	input = strings.ReplaceAll(input, "보지도 마", "")
	input = strings.ReplaceAll(input, "보지는 않", "")
	input = strings.ReplaceAll(input, "보지안으", "")
	input = strings.ReplaceAll(input, "보지안아", "")
	input = strings.ReplaceAll(input, "게 보지", "")
	input = strings.ReplaceAll(input, "어케 보지", "")
	input = strings.ReplaceAll(input, "보g", "보지")
	input = strings.ReplaceAll(input, "하나보지", "")
	input = strings.ReplaceAll(input, "켜보지", "")
	input = strings.ReplaceAll(input, "보지맙", "")
	input = strings.ReplaceAll(input, "초보지", "")
	input = strings.ReplaceAll(input, "로보지", "")
	input = strings.ReplaceAll(input, "가보지", "")
	input = strings.ReplaceAll(input, "홍보지", "")
	input = strings.ReplaceAll(input, "서보지", "")
	input = strings.ReplaceAll(input, "보지금", "")
	input = strings.ReplaceAll(input, "보지못", "")
	input = strings.ReplaceAll(input, "정지금", "")
	input = strings.ReplaceAll(input, "걸보지", "")
	input = strings.ReplaceAll(input, "보지는", "")
	input = strings.ReplaceAll(input, "보지지", "")
	input = strings.ReplaceAll(input, "보지않", "")
	input = strings.ReplaceAll(input, "해보지", "")
	input = strings.ReplaceAll(input, "보지마", "")
	input = strings.ReplaceAll(input, "보지말", "")
	input = strings.ReplaceAll(input, "안보지만", "")
	input = strings.ReplaceAll(input, "정보", "")
	input = strings.ReplaceAll(input, "지팡이", "")
	input = strings.ReplaceAll(input, "행보", "")
	input = strings.ReplaceAll(input, "바보지", "")
	input = strings.ReplaceAll(input, "바보짓", "")
	input = strings.ReplaceAll(input, "물어보지", "")
	input = strings.ReplaceAll(input, "하시나보지", "")
	input = strings.ReplaceAll(input, "오", "")
	input = strings.ReplaceAll(input, "언제 자지", "")
	input = strings.ReplaceAll(input, "언제자지", "")
	input = strings.ReplaceAll(input, "잠 자지", "")
	input = strings.ReplaceAll(input, "자지 말자고", "")
	input = strings.ReplaceAll(input, " 지급", "")
	input = strings.ReplaceAll(input, "남자지", "")
	input = strings.ReplaceAll(input, "여자지", "")
	input = strings.ReplaceAll(input, "감자지", "")
	input = strings.ReplaceAll(input, "왁자지", "")
	input = strings.ReplaceAll(input, "자지러", "")
	input = strings.ReplaceAll(input, "개발자", "")
	input = strings.ReplaceAll(input, "관리자", "")
	input = strings.ReplaceAll(input, "약탈자", "")
	input = strings.ReplaceAll(input, "타자지", "")
	input = strings.ReplaceAll(input, "혼자", "")
	input = strings.ReplaceAll(input, "자지원", "")
	input = strings.ReplaceAll(input, "사용자", "")
	input = strings.ReplaceAll(input, "경력자", "")
	input = strings.ReplaceAll(input, "지식", "")
	input = strings.ReplaceAll(input, "자지마", "")
	input = strings.ReplaceAll(input, "자지말", "")
	input = strings.ReplaceAll(input, "지원자", "")
	input = strings.ReplaceAll(input, "부자지", "")
	input = strings.ReplaceAll(input, "혜자지", "")
	input = strings.ReplaceAll(input, "잘 자지", "")
	input = strings.ReplaceAll(input, "일자지", "")
	input = strings.ReplaceAll(input, "일찍 자지", "")
	input = strings.ReplaceAll(input, "지원", "")
	input = strings.ReplaceAll(input, "자지금", "")
	input = strings.ReplaceAll(input, "자지않", "")
	input = strings.ReplaceAll(input, "어케자지", "")
	input = strings.ReplaceAll(input, "자지도마", "")
	input = strings.ReplaceAll(input, "자지는않", "")
	input = strings.ReplaceAll(input, "자지좀마", "")
	input = strings.ReplaceAll(input, "안자지", "")
	input = strings.ReplaceAll(input, "못자지", "")
	input = strings.ReplaceAll(input, "자지금", "")
	input = strings.ReplaceAll(input, "지건", "")
	input = strings.ReplaceAll(input, "감자", "")
	input = strings.ReplaceAll(input, "아", "")
	input = strings.ReplaceAll(input, "cess", "")
	input = strings.ReplaceAll(input, "```css", "")
	input = strings.ReplaceAll(input, "ex)", "")
	input = strings.ReplaceAll(input, "exit", "")
	input = strings.ReplaceAll(input, "ext", "")
	input = strings.ReplaceAll(input, "images", "")
	input = strings.ReplaceAll(input, "https", "")
	input = strings.ReplaceAll(input, "(ex", "")
	input = strings.ReplaceAll(input, ".ex", "")
	input = strings.ReplaceAll(input, "physics", "")
	input = strings.ReplaceAll(input, "features", "")
	input = strings.ReplaceAll(input, "exam", "")
	input = strings.ReplaceAll(input, "phase", "")
	input = strings.ReplaceAll(input, "except", "")
	input = strings.ReplaceAll(input, "sexual", "")
	input = strings.ReplaceAll(input, "sexy", "")
	input = strings.ReplaceAll(input, "엑섹스", "")
	input = strings.ReplaceAll(input, "엑", "")
	input = strings.ReplaceAll(input, "0ㅑ", "야")
	input = strings.ReplaceAll(input, "야스오", "")
	input = strings.ReplaceAll(input, "크시야", "")
	input = strings.ReplaceAll(input, "카구야", "")
	input = strings.ReplaceAll(input, "스파이", "")
	input = strings.ReplaceAll(input, "말이야", "")
	input = strings.ReplaceAll(input, "스티브", "")
	input = strings.ReplaceAll(input, "스쿼드", "")
	input = strings.ReplaceAll(input, "파랑색", "")
	input = strings.ReplaceAll(input, "오야스미", "")
	input = strings.ReplaceAll(input, "노란색", "")
	input = strings.ReplaceAll(input, "빨간색", "")
	input = strings.ReplaceAll(input, "초록색", "")
	input = strings.ReplaceAll(input, "보라색", "")
	input = strings.ReplaceAll(input, "청색", "")
	input = strings.ReplaceAll(input, "보라색", "")
	input = strings.ReplaceAll(input, "핑크색", "")
	input = strings.ReplaceAll(input, "남색", "")
	input = strings.ReplaceAll(input, "검은색", "")
	input = strings.ReplaceAll(input, "하양색", "")
	input = strings.ReplaceAll(input, "주황색", "")
	input = strings.ReplaceAll(input, "연두색", "")
	input = strings.ReplaceAll(input, "스공", "")
	input = strings.ReplaceAll(input, "스시", "")
	input = strings.ReplaceAll(input, "스키장", "")
	input = strings.ReplaceAll(input, "스킨", "")
	input = strings.ReplaceAll(input, "스킬", "")
	input = strings.ReplaceAll(input, "스틸", "")
	input = strings.ReplaceAll(input, "스탑", "")
	input = strings.ReplaceAll(input, "스트레스", "")
	input = strings.ReplaceAll(input, "해야", "")
	input = strings.ReplaceAll(input, "카시야스", "")
	input = strings.ReplaceAll(input, "야스톤", "")
	input = strings.ReplaceAll(input, "유니섹스", "")
	input = strings.ReplaceAll(input, "스튜디오", "")
	input = strings.ReplaceAll(input, "스티커", "")
	input = strings.ReplaceAll(input, "위대한", "")
	input = strings.ReplaceAll(input, "소유자", "")
	input = strings.ReplaceAll(input, "작업자", "")
	input = strings.ReplaceAll(input, "자기위로", "자위")
	input = strings.ReplaceAll(input, "위대하지", "")
	input = strings.ReplaceAll(input, "암살자", "")
	input = strings.ReplaceAll(input, "학자", "")
	input = strings.ReplaceAll(input, "개발자", "")
	input = strings.ReplaceAll(input, "관리자", "")
	input = strings.ReplaceAll(input, "약탈자", "")
	input = strings.ReplaceAll(input, "사용자", "")
	input = strings.ReplaceAll(input, "지원자", "")
	input = strings.ReplaceAll(input, "위대한", "")
	input = strings.ReplaceAll(input, "소유자", "")
	input = strings.ReplaceAll(input, "작업자", "")
	input = strings.ReplaceAll(input, "자기위로", "자위")
	for _, item := range cache.Sexual {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Belittle(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "뇬", "련")
	input = strings.ReplaceAll(input, "놈", "련")
	input = strings.ReplaceAll(input, "넘", "련")
	input = strings.ReplaceAll(input, "련", "년")
	input = strings.ReplaceAll(input, "뇬", "년")
	input = strings.ReplaceAll(input, "놈", "년")
	input = strings.ReplaceAll(input, "넘", "년")
	input = strings.ReplaceAll(input, "러운지", "")
	input = strings.ReplaceAll(input, "지킬 앤 하이드", "")
	input = strings.ReplaceAll(input, "지킬앤하이드", "")
	input = strings.ReplaceAll(input, "지킬 엔 하이드", "")
	input = strings.ReplaceAll(input, "지킬엔하이드", "")
	input = strings.ReplaceAll(input, "려운지", "")
	input = strings.ReplaceAll(input, "무서운지", "")
	input = strings.ReplaceAll(input, "라운지", "")
	input = strings.ReplaceAll(input, "운지법", "")
	input = strings.ReplaceAll(input, "싸운지", "")
	input = strings.ReplaceAll(input, "운지버섯", "")
	input = strings.ReplaceAll(input, "운 지린다", "")
	input = strings.ReplaceAll(input, "깔보다", "")
	input = strings.ReplaceAll(input, "깔보시", "")
	input = strings.ReplaceAll(input, "1년", "")
	input = strings.ReplaceAll(input, "2년", "")
	input = strings.ReplaceAll(input, "3년", "")
	input = strings.ReplaceAll(input, "4년", "")
	input = strings.ReplaceAll(input, "5년", "")
	input = strings.ReplaceAll(input, "6년", "")
	input = strings.ReplaceAll(input, "7년", "")
	input = strings.ReplaceAll(input, "8년", "")
	input = strings.ReplaceAll(input, "9년", "")
	input = strings.ReplaceAll(input, "0년", "")
	input = strings.ReplaceAll(input, "더운지역", "")
	input = strings.ReplaceAll(input, "나따까리", "")
	for _, item := range cache.Belittle {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Race(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "흑형님", "")
	for _, item := range cache.Race {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Parent(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "ㄴㄴ", "")
	input = strings.ReplaceAll(input, "미국", "")
	input = strings.ReplaceAll(input, "엄창못", "")
	input = strings.ReplaceAll(input, "l", "ㅣ")
	input = strings.ReplaceAll(input, "1", "ㅣ")
	input = strings.ReplaceAll(input, "ㄴㅣ", "니")
	input = strings.ReplaceAll(input, "ㅇㅣ-ㅣ", "애")
	input = strings.ReplaceAll(input, "ㅁㅣ", "미")
	input = strings.ReplaceAll(input, "도", "")
	for _, item := range cache.Parent {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Politics(input string) (bool, string) {
	input = ChangeUnicode(input)
	input = strings.ReplaceAll(input, "카카오톡", "")
	input = strings.ReplaceAll(input, "카톡", "")
	input = strings.ReplaceAll(input, "카페", "")
	input = strings.ReplaceAll(input, "하다가", "")
	input = strings.ReplaceAll(input, "먹다가", "")
	input = strings.ReplaceAll(input, "카와이", "")
	input = strings.ReplaceAll(input, "카츠", "")
	input = strings.ReplaceAll(input, "카레", "")
	input = strings.ReplaceAll(input, "니가", "")
	input = strings.ReplaceAll(input, "내가", "")
	input = strings.ReplaceAll(input, "너가", "")
	input = strings.ReplaceAll(input, "우리가", "")
	input = strings.ReplaceAll(input, "너희가", "")
	input = strings.ReplaceAll(input, "카카오", "")
	input = strings.ReplaceAll(input, "너희가", "")
	input = strings.ReplaceAll(input, "카세트", "")
	input = strings.ReplaceAll(input, "카플레이어", "")
	input = strings.ReplaceAll(input, "카운터", "")
	input = strings.ReplaceAll(input, "카정", "")
	input = strings.ReplaceAll(input, "카드", "")
	for _, item := range cache.Politics {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}
	return false, ""
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
func Japanese(input string) (bool, string) {
	input = ChangeUnicode(input)
	for _, item := range cache.Japanese {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Chinese(input string) (bool, string) {
	input = ChangeUnicode(input)
	for _, item := range cache.Chinese {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
func Special(input string) (bool, string) {
	input = ChangeUnicode(input)
	for _, item := range cache.Emoji {
		if strings.Contains(input, item) {
			return true, item
		}
		//if Levenshtein(input, item) <= 3 {
		//	return true, item
		//}
	}

	return false, ""
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
	var swear string

	detect, swear = General(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DGeneral,
		}
	}

	detect, swear = Minor(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DMinor,
		}
	}

	detect, swear = Sexual(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DSexual,
		}
	}

	detect, swear = Belittle(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DBelittle,
		}
	}

	detect, swear = Race(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DRace,
		}
	}

	detect, swear = Parent(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DParent,
		}
	}

	detect, swear = Politics(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DPolitics,
		}
	}

	detect, swear = English(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DEnglish,
		}
	}

	detect, swear = Japanese(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DJapanese,
		}
	}

	detect, swear = Chinese(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DChinese,
		}
	}

	detect, swear = Special(input)
	if detect {
		return CheckInfo{
			Detect: true,
			Swear:  swear,
			Type:   DSpecial,
		}
	}

	return CheckInfo{
		Detect: false,
		Swear:  "",
		Type:   DNone,
	}
}
