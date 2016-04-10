package verbo

import (
	"regexp"
	"unicode/utf8"
)

func adjacent(str string, direction int) string {
	if utf8.RuneCountInString(str) == 0 {
		return ""
	}
	return string(int(str[utf8.RuneCountInString(str)-1]) + direction)
}

func defaultToWhiteSpace(characters string) string {
  if characters == "" {
    return "\\s"
	} else {
    return "[" + escapeRegExp(characters) + "]"
	}
}

func escapeRegExp(str string) string {
	re := regexp.MustCompile(`([.*+?^=!:${}()|[\]\/\\])`)
  return re.ReplaceAllString(str, `\\$1`)
}

func sizeOfLastRunes(str string, count int) int {
	return len(str[len(str) - count:])
}

func stringLengthToRunesSize(str string, length int) int {

	count := utf8.RuneCountInString(str)
	if count < length {
		length = count
	}

	size := 0
	i := 0
	for i < length {
		_, s := utf8.DecodeRuneInString(str)
		str = str[s:]
		size += s
		i++
	}

	return size
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
	  return b
	}
}

func strRepeat(str string, qty int) string {
	if qty < 1 {
		return ""
	}
	result := ""
	for qty > 0 {
		if qty&1 == 1 {
			result += str
		}
		qty >>= 1
		str += str
	}
	return result
}

func toPositive(number int) int {
	if number < 0 {
		return 0
	} else {
		return +number
	}
}
