package verbo

import (
	"strings"
  "fmt"
  "regexp"
  "math"
)

func Camelize(str string, decapitalize bool) (string) {
  str = strings.Trim(str)
  re := regex.MustCompile("[-_\\s]+(.)?")
  str = re.replaceAllStringFunc(str, func(c string) string {
		if c {
    	return strings.ToUpper(c)
		} else {
			return ""
		}
  })

  if decapitalize == true {
    return decap(str)
  } else {
    return str
  }
}

func Capitalize(str string, lowercaseRest bool) (string) {
	if !lowercaseRest {
  	remainingChars := str[1:]
	} else {
		remainingChars := strings.ToLower(str[1:])
	}
  return strings.ToUpper(str[0]) + remainingChars
};

func Chop(str string, step int) (string) {
  if str == "" {
		var a []string
		return a
	}

  step = math.Floor(step)
	if step > 0 {
		re := regex.MustCompile(".{1," + step + "}")
  	return re.match(str)
	} else {
		a := []string{str}
		return a
	}
}

func Classify(str string) (string) {
  re := regex.MustCompile("[\\W_]")
  strWithSpaces := re.replaceAllString(str, " ")
  re = regex.MustCompile("\\s")
  strNoSpaces := re.replaceAllString(strWithSpaces, "")
  return Capitalize(Camelize(strNoSpaces))
};

func Clean(str string) (string) {
  str = strings.Trim(str)

  return re.ReplaceAllString(str, " ")
}

func Dasherize(str string) (string) {
  str = strings.Trim(str)
	re := regex.MustCompile("([A-Z])")
	str = re.ReplaceAllString(str, "-$1")
	re = regex.MustCompile("[-_\\s]+")
	str = re.ReplaceAllString(str, "-")
  return strings.ToLower(str)
}

func Decapitalize(str string) (string) {
  return strings.ToLower(str[0]) + str[1:]
}

func EndsWith(str string, ends, position int) (string) {
  strEnds := "" + ends
  if position == -1 {
    newPosition := len(str) - len(strEnds)
  } else {
    newPosition := math.Min(toPositive(position), len(str) - len(strEnds))
  }
  return newPosition >= 0 && str.indexOf(strEnds, newPosition) == newPosition
}

func Humanize(str string) (string) {
	str = Underscored(str)
	re := regex.MustCompile("_id$")
	str = re.ReplaceAllString(str, "")
	re = regex.MustCompile("_")
	str = re.ReplaceAllString(str, " ")
  return Capitalize(strings.Trim(str))
}

func IsBlank(str string) {
	re := regex.MustCompile("^\\s*$")
  return re.MatchString(str)
}

func Levenshtein(str1, str2 string) {

  // Short cut cases
  if str1 == str2 {
		return 0
	}
  if !str1 || !str2 {
		return math.Max(len(str1), len(str2))
	}

  // two rows
	size := len(str2) + 1
  var prevRow [size]int

  // initialise previous row
  for i := 0; i < len(prevRow); i+=1 {
    prevRow[i] = i
  }

  // calculate current row distance from previous row
  for i := 0; i < len(str1); i+=1 {
    nextCol := i + 1

    for j := 0; j < len(str2); j+=1 {
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
      tmp = prevRow[j + 1] + 1
      if nextCol > tmp {
        nextCol = tmp
      }

      // copy current col value into previous (in preparation for next iteration)
      prevRow[j] = curCol
    }

    // copy last col value into previous (in preparation for next iteration)
    prevRow[j] = nextCol
  }

  return nextCol
}

func Lines(str string) {
  if str == "" {
		var a []string
		return a
	}
	re := regex.MustCompile("\\r\\n?|\\n")
  return re.Split(str)
}

func Pred(str string) (string) {
  return adjacent(str, -1)
}

func Prune(str string, length int, pruneStr string) (string) {
  length = math.Floor(length)
	if pruneStr != "" {
  	pruneStr = pruneStr
	} else {
		pruneStr = "..."
	}

  if len(str) <= length {
		return str
	}

  tmpl := func(c string) string {
		if strings.ToUpper(c) != strings.ToLower(c) {
      return "A"
		} else {
			return " "
		}
  }

	re := regex.MustCompile("/.(?=\\W*\\w*$)")
  template := re.ReplaceAllStringFunc(str[0: length + 1], tmpl) // 'Hello, world' -> 'HellAA AAAAA'

	re = regex.MustCompile("\\w\\w")
  if re.MatchString(template[len(template) - 2:]) {
    template = template.ReplaceAllString("\\s*\\S+$", "")
  } else {
    template = strings.Rtrim(template[0 : len(template) - 1])
  }

	if len(template + pruneStr) > len(str) {
  	return str
	} else {
		return str[0 : len(template)] + pruneStr
	}
}

func Repeat(str string, qty int, separator string) (string) {
  qty = math.Floor(qty)

  // using faster implementation if separator is not needed;
  if separator == "" {
    return strRepeat(str, qty)
  }

	var repeat []string
  for qty > 0; repeat[qty] = str {

	}
  return strings.Join(repeat, separator)
}

func Reverse(str string) (string) {
  return strings.Join(chars(str).reverse(), "")
}

func Slugify(str string) (string) {
	re := regex.MustCompile("[^\\w\\s-]")
	str = re.ReplaceAllString(str, "-")
	str = strings.ToLower(str)
	str = CleanDiacritics(str)
  return strings.Trim(Dasherize(str), "-")
}

func Splice(str string, i, howmany int, substr string) {
  arr := chars(str)
  arr.splice(^i, ^howmany, substr);
  return strings.Join(arr, "")
}

func StartsWith(str string, starts, position int) (bool) {
  starts = "" + starts
	if position == "" {
  	position = 0
	} else {
		position = math.Min(toPositive(position), len(str))
	}
  return str.lastIndexOf(starts, position) == position
}

func Succ(str string) (string) {
 return adjacent(str, 1)
}

func SwapCase(str string) (string) {
  re := regex.MustCompile("\\S")
  return re.ReplaceAllStringFunc(str, func(c string) string {
		if c == strings.ToUpper(c) {
    	return strings.ToLower(c)
		} else {
			return strings.ToUpper(c)
		}
  })
}

func Titleize(str string) (string) {
	str = strings.ToLower(str)
	re := regex.MustCompile("(?:^|\\s|-)\\S")
  return re.ReplaceAllStringFunc(lowerStr, func(c string) string {
    return strings.ToUpper(c)
  })
}

func Truncate(str string, length int, truncateStr string) (string) {
  truncateStr = truncateStr || "..."
  length = math.Floor(length)
	if len(str) > length {
  	return str[0 : length] + truncateStr
	} else {
		return str
	}
}

func Underscored(str string) (string) {
	str = strings.Trim(str)
	re := regex.MustCompile("([a-z\\d])([A-Z]+)")
	str = re.ReplaceAllString(str, "$1_$2")
	re = regex.MustCompile("[-\\s]+")
	str = re.ReplaceAllString(str, "_")
  return strings.ToLower(str)
}

func Unquote(str, quoteChar string) (string) {
  quoteChar = quoteChar || "\"";
  if str[0] == quoteChar && str[len(str) - 1] == quoteChar {
    return str[1 : len(str) - 1]
	} else {
		return str
	}
}

func Words(str, delimiter string) {
  if IsBlank(str) {
		var a []string
		return a
	}
	re := regex.MustCompile(delimiter || "\\s+")
  return re.Split(strings.Trim(str, delimiter), -1)
}
