<div align="center">
  <h1>Korcen.go</h1>

  [![Go Version](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/go_version.svg)](https://go.dev/)
  [![Module Version](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/module_version.svg)](https://pkg.go.dev/github.com/fluffy-melli/korcen-go)
</div>
<div align="center">
  <h3>
    <a href="https://github.com/fluffy-melli/korcen-go">KR</a> /
    <a href="https://github.com/fluffy-melli/korcen-go/blob/main/docs/README.EN.md">EN</a>
  </h3>
</div>

![131_20220604170616](https://user-images.githubusercontent.com/85154556/171998341-9a7439c8-122f-4a9f-beb6-0e0b3aad05ed.png)

# 🛠 제작자

>[Tanat05](https://github.com/Tanat05) / [korcen](https://github.com/Tanat05/korcen)
```
https://github.com/Tanat05/korcen
---------------------------------
이 프로젝트는 원본 `korcen` 프로젝트를 수정하여 배포한 것입니다.
원본 프로젝트는 `https://github.com/Tanat05/korcen`에서 확인할 수 있습니다.

이 프로젝트의 라이센스 또한 Apache-2.0 을 따르고 있습니다
```

>[gyarang](https://github.com/gyarang) / [gohangul](https://github.com/gyarang/gohangul)
```
https://github.com/gyarang/gohangul
-----------------------------------
이 프로젝트는 `gohangul` 을 사용해서 배포한 것입니다.
해당 프로젝트는 `https://github.com/gyarang/gohangul`에서 확인할 수 있습니다.

해당 프로젝트의 라이센스인 MIT License 또한 따르고 있습니다
```

![Apache-2.0](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/Apache-2.0.png)

`korcen-go`는 `Apache-2.0` & `MIT License` 라이선스를 `모두` 따르고 있습니다.
코드를 사용할 경우 라이선스 내용을 준수해주세요. 

Copyright© All rights reserved.

---

# ❓ 주요 기능

```diff
[부적절한 언어 감지]
- 일반적인 부적절한 언어
- 경미한 부적절한 언어
- 성적 언어나 표현
- 비하하는 표현
- 인종 차별적인 부적절한 표현
- 부모를 비하하거나 공격하는 언어
- 정치적 부적절한 언어나 표현
- 영어에서 부적절한 표현
- 일본에서 부적절한 표현
- 중국에서 부적절한 표현
- 기타 특수한 부적절한 표현
```

---

# ⬇️ 설치 방법

>mod
```sh
$ go get github.com/fluffy-melli/korcen-go
L v0.7.0 버전 이후로 함수 출력이 달라졌으므로 사용시에 유의해주시기 바랍니다
```

>golang 코드 예제
```go
package main

import (
	"fmt"

	"github.com/fluffy-melli/korcen-go"
)

func main() {
	// Discord Message Event Logic
	ck := korcen.Check("MESSAGE")
	if ck.Detect {
		dtif := ck.Swear[0]
		message := ck.NewText[:dtif.Start] + "\x1b[1m\x1b[41m\"" + ck.NewText[dtif.Start:dtif.End] + "\"\x1b[0m" + ck.NewText[dtif.End:]
		discord_ansi := "```ansi\n" + message + "\n```"
		fmt.Println(discord_ansi)
	}
	// Discord Send Logic
}
```