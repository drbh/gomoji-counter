package main

import (
	"github.com/davecgh/go-spew/spew"
	gomojicount "github.com/drbh/gomoji-counter"
)

func main() {
	spew.Dump(gomojicount.GetEmojiFrequencyCount(
		"drbh ✅ I hope this ✅ works! ✅🏰✌️ 🤞 🖖 🤘 🤙 🖐️ ✋ 👌 "))
}
