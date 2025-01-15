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

해당 프로젝트의 라이센스 또한 Apache-2.0 을 따르고 있습니다
```
![Apache-2.0](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/Apache-2.0.png)

`korcen-go`는 원본 프로젝트의 `Apache-2.0` 라이선스를 따르고 있습니다. 코드를 사용할 경우 라이선스 내용을 준수해주세요. 

- 라이선스 고지 및 저작권 고지 필수(일반인이 접근 가능한 부분에 표시)

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
```

>golang
```go
package main

import (
	"fmt"

	"github.com/fluffy-melli/korcen-go"
)

func main() {
	fmt.Println(korcen.Check(""))
}
```