package gl

import (
	"github.com/easonlin404/gl/util"
	"strings"
)

var Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var Base = len(Alphabet)

func Encode(i int) string {
	if i == 0 {
		return string(Alphabet[0])
	}

	s := ""

	for i > 0 {
		s += string(Alphabet[i%Base])
		i = i / Base
	}

	return util.Reverse(s)
}

func Base62(id int) []int {
	var value []int
	if id == 0 {
		return append(value,0)
	}
	for id > 0 {
		value = append(value, id % 62)
		id = id / 62
	}
	return value
}

func Decode(s string) int {
	i := 0

	for _, elem := range s {
		i = (i * Base) + strings.Index(Alphabet, string(elem))
	}

	return i
}
