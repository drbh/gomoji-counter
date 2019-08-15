package main

// package gomojicount

import (
	"fmt"

	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/desc"
)

type GoMoji struct {
	Emojis *sortedmap.SortedMap
}

func (gmji *GoMoji) process(longList string) {
	sm := sortedmap.New(4, desc.Int)

	runes_array := []rune(longList)

	for _, v := range runes_array {
		if v > 127 && v != 65039 && v != 8205 {
			if val, ok := sm.Get(string(v)); ok {
				newcount := val.(int) + 1
				sm.Replace(string(v), newcount)
			} else {
				sm.Insert(string(v), 1)
			}
		}
	}
	gmji.Emojis = sm
}

type EmojiFreq struct {
	emoji string
	count int
}

func (gmji *GoMoji) values() []EmojiFreq {
	var vals = []EmojiFreq{}
	iterCh, err := gmji.Emojis.IterCh()
	if err != nil {
		fmt.Println(err)
	} else {
		defer iterCh.Close()

		for rec := range iterCh.Records() {
			// fmt.Printf("%+v\n", rec)
			vals = append(vals, EmojiFreq{
				rec.Key.(string),
				rec.Val.(int),
			})
		}
	}
	return vals
}

func GetEmojiFrequencyCount(text string) []EmojiFreq {
	gm := GoMoji{}
	gm.process(text)
	return gm.values()
}

func main() {

	longList := "david went to ✅to ✅to ✅to ✅to ✅to ✅ theh 🎆"

	gm := GoMoji{}
	gm.process(longList)
	fmt.Println(gm.values())

	longList = "💪 arms!"

	gm.process(longList)
	fmt.Println(gm.values())
}
