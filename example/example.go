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
