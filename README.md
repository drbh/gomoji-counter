# Get Emjoi Counts!



```golang
package main

import (
	"github.com/davecgh/go-spew/spew"
	gomojicount "github.com/drbh/gomoji-counter"
)

func main() {
	spew.Dump(gomojicount.GetEmojiFrequencyCount(
		"drbh ✅ I hope this ✅ works! ✅🏰✌️ 🤞 🖖 🤘 🤙 🖐️ ✋ 👌 "))
}
```



```bash
([]gomojicount.EmojiFreq) (len=10 cap=16) {
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=3) "✅",
  count: (int) 3
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "🏰",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=3) "✌",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "🤞",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "🖖",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "🤘",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "🤙",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "🖐",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=3) "✋",
  count: (int) 1
 },
 (gomojicount.EmojiFreq) {
  emoji: (string) (len=4) "👌",
  count: (int) 1
 }
}
```