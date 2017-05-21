package atbash

import (
	"strings"
)

const testVersion = 2

var plain = "abcdefghijklmnopqrstuvwxyz"
var cipher = "zyxwvutsrqponmlkjihgfedcba"
var groupFixedLength = 5
var punctuation = ' '

func Atbash(input string) string {
	output := make([]rune, 0)
	input = strings.ToLower(input)
	groupLength := 0
	newCipher := rune(' ')
	for _, c := range input {
		if (c < 'a' || c > 'z') && (c < '1' || c > '9') {
			continue
		}
		newCipher = rune(' ')
		if c >= '1' && c <= '9' {
			newCipher = c
		} else {
			for j, p := range plain {
				if p == c {
					newCipher = rune(cipher[j])
					break
				}
			}
		}
		if newCipher != ' ' {
			groupLength++
			if groupLength == groupFixedLength+1 {
				output = append(output, punctuation)
				groupLength = 1
			}
			output = append(output, newCipher)
		}
	}
	return string(output)
}
