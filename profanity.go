package korcen

import "github.com/fluffy-melli/korcen-go/cache"

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
func English(input string) (bool, string) {
	newtext := ChangeUnicode(input)
	prof := NewProfanity(cache.English)
	return prof.Censor(newtext)
}
