<div align="center">
  <h1>Korcen.go</h1>

  ![Go Version](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/go_version.svg)
  ![Module Version](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/module_version.svg)
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

Copyright© All rights reserved.

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