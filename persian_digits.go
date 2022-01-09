package persiandigits

import (
	"unicode/utf8"
)

func ToEnglish(text string) (converted string, n int) {
	result := make([]rune, 0, utf8.RuneCountInString(text))
	for _, c := range text {
		if int(c) >= '۰' && int(c) <= '۹' {
			result = append(result, rune(int(c)-'۰'+'0'))
			n += 1
		} else {
			result = append(result, c)
		}
	}
	converted = string(result)

	return converted, n
}

func ToPersian(text string) (converted string, n int) {
	result := make([]rune, 0, utf8.RuneCountInString(text))
	for _, c := range text {
		if int(c) >= '0' && int(c) <= '9' {
			result = append(result, rune(int(c)-'0'+'۰'))
			n += 1
		} else {
			result = append(result, c)
		}
	}
	converted = string(result)

	return converted, n
}
