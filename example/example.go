package main

import (
	"fmt"

	"github.com/fluffy-melli/korcen-go"
)

func main() {
	korcen.InitProfanityData()

	sampleInput := "MESSAGE"

	ck := korcen.Check(sampleInput)

	if ck.Detect {
		for _, dtif := range ck.Swear {
			message := ck.NewText[:dtif.Start] +
				"\x1b[1m\x1b[41m\"" + ck.NewText[dtif.Start:dtif.End] + "\"\x1b[0m" +
				ck.NewText[dtif.End:]

			discord_ansi := "```ansi\n" + message + "\n```"
			fmt.Println(discord_ansi)
		}
	} else {
		fmt.Println("비속어가 감지되지 않았습니다.")
	}

}
