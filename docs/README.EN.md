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

# 🛠 Maker

>[Tanat05](https://github.com/Tanat05) / [korcen](https://github.com/Tanat05/korcen)
```
https://github.com/Tanat05/korcen
---------------------------------
This project is a modified version of the original `korcen`
The original project can be found at `https://github.com/Tanat05/korcen`

This project also follows the Apache-2.0 license.
```

>[gyarang](https://github.com/gyarang) / [gohangul](https://github.com/gyarang/gohangul)
```
https://github.com/gyarang/gohangul
-----------------------------------
This project is distributed using `gohangul`
You can find the project at `https://github.com/gyarang/gohangul`

It also follows the MIT License, which is the license for the respective project.
```

![Apache-2.0](https://github.com/fluffy-melli/korcen-go/blob/main/docs/asset/Apache-2.0.png)

`korcen-go `follows both the `Apache-2.0` and `MIT License` licenses.
Please comply with the license terms when using the code.

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
L Since version 0.7.0, the function output has changed, so please be cautious when using it.
```

>golang code example
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