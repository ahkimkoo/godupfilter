package shingle

import (
	// "fmt"
	"github.com/huichen/sego"
	"github.com/zhenjl/cityhash"
	"strings"
)

var segmenter sego.Segmenter

func init() {
	segmenter.LoadDictionary("src/github.com/huichen/sego/data/dictionary.txt")
}

func Shingling(text []byte, shingle_size int) map[uint32]bool {
	segments := segmenter.Segment(text)
	tokens := make(map[uint32]bool)
	for i := 0; i < len(segments)-shingle_size; i++ {
		sub_segments := segments[i : i+shingle_size]
		var s []string
		for _, segment := range sub_segments {
			if segment.Token().Pos() != "x" {
				s = append(s, segment.Token().Text())
			}
		}
		token := strings.Join(s, "#")
		token_byte := []byte(token)
		hash := cityhash.CityHash32(token_byte, uint32(len(token_byte)))
		// fmt.Printf("%v===>%v\n", token, hash)
		_, ok := tokens[hash]
		if !ok {
			tokens[hash] = true
		}
	}
	return tokens
}

func Similarity(a, b map[uint32]bool) float32 {
	var x, y map[uint32]bool
	if len(a) <= len(a) {
		x = a
		y = b
	} else {
		x = b
		y = a
	}
	same_count := 0
	for k := range x {
		_, ok := y[k]
		if ok {
			same_count++
		}
	}
	// fmt.Printf("diffenece: %d, same:%d", (len(a) + len(b) - same_count), same_count)
	return float32(same_count) / float32(len(a)+len(b)-same_count)
}
