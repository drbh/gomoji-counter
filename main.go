package main

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

func (gmji *GoMoji) values() {
	iterCh, err := gmji.Emojis.IterCh()
	if err != nil {
		fmt.Println(err)
	} else {
		defer iterCh.Close()

		for rec := range iterCh.Records() {
			fmt.Printf("%+v\n", rec)
		}
	}
}

func main() {

	longList := "david went to ✅to ✅to ✅to ✅to ✅to ✅ theh 🎆"

	gm := GoMoji{}
	gm.process(longList)
	gm.values()

	longList = "💪 arms!"

	gm.process(longList)
	gm.values()
}
