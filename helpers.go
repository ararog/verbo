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
