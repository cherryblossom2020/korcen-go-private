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

# 🛠 Maker

>[Tanat05](https://github.com/Tanat05) / [korcen](https://github.com/Tanat05/korcen)
```
https://github.com/Tanat05/korcen
---------------------------------
This project is a modified version of the original `korcen` project.
The original project can be found at `https://github.com/Tanat05/korcen`.

This project also follows the Apache-2.0 license.
```
![Apache-2.0](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/Apache-2.0.png)

The korcen-go project follows the Apache-2.0 license. If you use the code, please comply with the terms of the license.

- License and copyright notices must be displayed (in areas accessible to the general public).

Copyright© All rights reserved.

---

# ❓ Main Features

```diff
[Inappropriate Language Detection]
- General inappropriate language
- Minor inappropriate language
- Sexual language or expressions
- Derogatory expressions
- Racially discriminatory inappropriate expressions
- Language that belittles or attacks parents
- Politically inappropriate language or expressions
- Inappropriate expressions in English
- Inappropriate expressions in Japan
- Inappropriate expressions in China
- Other special inappropriate expressions
```

---

# ⬇️ Installation

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