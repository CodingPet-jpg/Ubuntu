package intset

import (
	"bytes"
	"fmt"
)

type Intset struct {
	words []uint64
}

func (set *Intset) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(set.words) {
		set.words = append(set.words, 0)
	}
	set.words[word] |= 1 << bit
}

func (set *Intset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(set.words) && set.words[word]&(1<<bit) != 0
}

// 将set和target取并集并放在set中返回
func (set *Intset) UnionWith(target *Intset) {
	for index, item := range target.words {
		if index < len(set.words) {
			set.words[index] |= item
		} else {
			set.words = append(set.words, item)
		}
	}
}

func (set *Intset) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range set.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
