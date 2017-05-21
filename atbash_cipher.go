package atbash

import (
	"strings"
)

const testVersion = 2

var groupFixedLength = 5
var punctuation = " "

func Atbash(input string) string {
	input = strings.ToLower(input)
	var group []string
	s := make([]rune, groupFixedLength)
	i := 0
	for _, c := range input {
		if !IsLetter(c) && !IsNumber(c) {
			continue
		}
		s[i] = Transpose(c)
		i++
		if i == groupFixedLength {
			group = append(group, string(s))
			i = 0
		}
	}
	if i > 0 {
		group = append(group, string(s[0:i]))
	}
	return strings.Join(group, punctuation)
}
func IsLetter(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}
func IsNumber(r rune) bool {
	if r >= '1' && r <= '9' {
		return true
	}
	return false
}
func Transpose(r rune) rune {
	if IsNumber(r) {
		return r
	}
	return 219 - r

}
