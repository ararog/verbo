package verbo

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"math"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Camelize(str string, decapitalize bool) string {
	str = strings.TrimSpace(str)
	re := regexp.MustCompile(`[-_\s]+(.)?`)
	str = re.ReplaceAllStringFunc(str, func(c string) string {
		if c != "" {
			out := re.FindStringSubmatch(c)
			return strings.ToUpper(out[1])
		} else {
			return ""
		}
	})

	if decapitalize {
		return Decapitalize(str)
	} else {
		return str
	}
}

func Capitalize(str string, lowercaseRest bool) string {
	remainingChars := ""
	if !lowercaseRest {
		remainingChars = str[1:]
	} else {
		remainingChars = strings.ToLower(str[1:])
	}
	return strings.ToUpper(string(str[0])) + remainingChars
}

func Chop(str string, step int) []string {
	if str == "" {
		var a []string
		return a
	}

	step = int(math.Floor(float64(step)))
	if step > 0 {
		re := regexp.MustCompile(".{1," + string(step) + "}")
		return re.FindStringSubmatch(str)
	} else {
		a := []string{str}
		return a
	}
}

func Classify(str string) string {
	re := regexp.MustCompile(`[\W_]`)
	str = re.ReplaceAllString(str, " ")
	return Capitalize(Camelize(str, false), false)
}

func Clean(str string) string {
	str = strings.TrimSpace(str)
	re := regexp.MustCompile(`\s\s+`)
	return re.ReplaceAllString(str, " ")
}

func CleanDiacritics(str string) string {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

func Dasherize(str string) string {
	str = strings.TrimSpace(str)
	re := regexp.MustCompile("([A-Z])")
	str = re.ReplaceAllString(str, "-$1")
	re = regexp.MustCompile(`[-_\s]+`)
	str = re.ReplaceAllString(str, "-")
	return strings.ToLower(str)
}

func Decapitalize(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

func Humanize(str string) string {
	str = Underscored(str)
	re := regexp.MustCompile(`_id$`)
	str = re.ReplaceAllString(str, "")
	re = regexp.MustCompile(`_`)
	str = re.ReplaceAllString(str, " ")
	return Capitalize(strings.TrimSpace(str), false)
}

func IsBlank(str string) bool {
	re := regexp.MustCompile(`^\s*$`)
	return re.MatchString(str)
}

func LeftPad(str string, length int, padStr string) string {
  return Pad(str, length, padStr, "")
}

func Levenshtein(str1, str2 string) int {

	// Short cut cases
	if str1 == str2 {
		return 0
	}
	if str1 == "" || str2 == "" {
		return int(math.Max(float64(utf8.RuneCountInString(str1)), float64(utf8.RuneCountInString(str2))))
	}

	// two rows
	size := utf8.RuneCountInString(str2) + 1
	prevRow := make([]int, size, size)

	// initialise previous row
	for i := 0; i < len(prevRow); i += 1 {
		prevRow[i] = i
	}

	nextCol := 0

	// calculate current row distance from previous row
	i := 0
	for i < utf8.RuneCountInString(str1) {
		nextCol = i + 1
		j := 0
		for j < utf8.RuneCountInString(str2) {
			curCol := nextCol

			// substution
			val := 0
			if str1[i] == str2[j] {
				val = 0
			} else {
				val = 1
			}

			nextCol := prevRow[j] + val

			// insertion
			tmp := curCol + 1
			if nextCol > tmp {
				nextCol = tmp
			}
			// deletion
			tmp = prevRow[j+1] + 1
			if nextCol > tmp {
				nextCol = tmp
			}

			// copy current col value into previous (in preparation for next iteration)
			prevRow[j] = curCol
			j += 1
		}
		// copy last col value into previous (in preparation for next iteration)
		prevRow[j] = nextCol
		i += 1
	}

	return nextCol
}

func Lines(str string) []string {
	if str == "" {
		var a []string
		return a
	}
	re := regexp.MustCompile(`\r\n?|\n`)
	return re.Split(str, -1)
}

func Pad(str string, length int, padStr, padtype string) string {
  length = int(math.Floor(float64(length)))

  var padlen = 0

  if padStr == "" {
    padStr = " "
  } else if utf8.RuneCountInString(padStr) > 1 {
    padStr = string(padStr[0])
	}

  switch padtype {
  case "right":
    padlen = length - utf8.RuneCountInString(str)
    return str + strRepeat(padStr, padlen)
  case "both":
    padlen = length - utf8.RuneCountInString(str)
    return strRepeat(padStr, int(math.Ceil(float64(padlen) / 2))) + str + strRepeat(padStr, int(math.Floor(float64(padlen) / 2)))
  default: // 'left'
    padlen = length - utf8.RuneCountInString(str)
    return strRepeat(padStr, padlen) + str
  }
}

func Pred(str string) string {
	return adjacent(str, -1)
}

func Prune(str string, length int, pruneStr string) string {
	length = int(math.Floor(float64(length)))
	if pruneStr != "" {
		pruneStr = pruneStr
	} else {
		pruneStr = "..."
	}

	if utf8.RuneCountInString(str) <= length {
		return str
	}

	tmpl := func(c string) string {
		if strings.ToUpper(c) != strings.ToLower(c) {
			return "A"
		} else {
			return " "
		}
	}

	re := regexp.MustCompile(`.(?=\W*\w*$)`)
	template := re.ReplaceAllStringFunc(str[0:length+1], tmpl) // 'Hello, world' -> 'HellAA AAAAA'

	re = regexp.MustCompile(`\w\w`)
	if re.MatchString(template[len(template)-2:]) {
		re = regexp.MustCompile(`\s*\S+$`)
		template = re.ReplaceAllString(template, "")
	} else {
		template = strings.TrimRight(template[0:len(template)-1], " ")
	}

	if utf8.RuneCountInString(template+pruneStr) > utf8.RuneCountInString(str) {
		return str
	} else {
		return str[0:utf8.RuneCountInString(template)] + pruneStr
	}
}

func Repeat(str string, qty int, separator string) string {
	qty = int(math.Floor(float64(qty)))

	// using faster implementation if separator is not needed;
	if separator == "" {
		return strRepeat(str, qty)
	}

	var repeat []string
	for qty > 0 {
		repeat[qty] = str
	}
	return strings.Join(repeat, separator)
}

func Reverse(str string) string {
	n := utf8.RuneCountInString(str)
	runes := make([]rune, n)
	for _, rune := range str {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func RightPad(str string, length int, padStr string) string {
  return Pad(str, length, padStr, "right");
}

func Slugify(str string) string {
	re := regexp.MustCompile(`[^\w\s-]`)
	str = re.ReplaceAllString(str, "-")
	str = strings.ToLower(str)
	str = CleanDiacritics(str)
	return strings.Trim(Dasherize(str), "-")
}

/*
func Splice(str string, i, howmany int, substr string) (string) {
  arr := chars(str)
  arr.splice(^i, ^howmany, substr);
  return strings.Join(arr, "")
}
*/

func Succ(str string) string {
	return adjacent(str, 1)
}

func SwapCase(str string) string {
	re := regexp.MustCompile("\\S")
	return re.ReplaceAllStringFunc(str, func(c string) string {
		if c == strings.ToUpper(c) {
			return strings.ToLower(c)
		} else {
			return strings.ToUpper(c)
		}
	})
}

func Titleize(str string) string {
	str = strings.ToLower(str)
	re := regexp.MustCompile("(?:^|\\s|-)\\S")
	return re.ReplaceAllStringFunc(str, func(c string) string {
		return strings.ToUpper(c)
	})
}

func Trim(str, characters string) string {
  if characters == "" {
		return strings.TrimSpace(str)
	}

	characters = defaultToWhiteSpace(characters)
	re := regexp.MustCompile("^" + characters + "+|" + characters + "+$")
	return re.ReplaceAllString(str, "")
}

func Truncate(str string, length int, truncateStr string) string {
	if truncateStr == "" {
		truncateStr = "..."
	}
	length = int(math.Floor(float64(length)))
	if utf8.RuneCountInString(str) > length {
		return str[0:length] + truncateStr
	} else {
		return str
	}
}

func Underscored(str string) string {
	str = strings.TrimSpace(str)
	re := regexp.MustCompile(`([a-z\d])([A-Z]+)`)
	str = re.ReplaceAllString(str, `$1-$2`)
	re = regexp.MustCompile(`[-\s]+`)
	str = re.ReplaceAllString(str, "_")
	return strings.ToLower(str)
}

func Unquote(str, quoteChar string) string {
	if quoteChar == "" {
		quoteChar = "\""
	}
	if string(str[0]) == quoteChar &&
		string(str[utf8.RuneCountInString(str)-1]) == quoteChar {
		return str[1 : utf8.RuneCountInString(str)-1]
	} else {
		return str
	}
}

func Words(str, delimiter string) []string {
	if IsBlank(str) {
		var a []string
		return a
	}
	if delimiter == "" {
		delimiter = `\s+`
	}
	re := regexp.MustCompile(delimiter)
	return re.Split(Trim(str, delimiter), -1)
}
