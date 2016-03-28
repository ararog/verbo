package verbo

import (
	"strings"
  "fmt"
)

func Camelize(str string, decapitalize bool) (string) {
  str := trim(str).replace(/[-_\s]+(.)?/g, function(match, c) {
    return c ? c.toUpperCase() : ""
  })

  if (decapitalize === true) {
    return decap(str)
  } else {
    return str
  }
}

func Capitalize(str string, lowercaseRest bool) (string) {
  str = makeString(str)
  var remainingChars = !lowercaseRest ? str.slice(1) : str.slice(1).toLowerCase()

  return str.charAt(0).toUpperCase() + remainingChars
};

func Chop(str string, step int) (string) {
  if (str == null) return []
  str = String(str)
  step = ~~step
  return step > 0 ? str.match(new RegExp(".{1," + step + "}", "g")) : [str]
}

func Classify(str string) (string) {
  str = makeString(str)
  return capitalize(camelize(str.replace(/[\W_]/g, " ")).replace(/\s/g, ""))
};

func Clean(str string) (string) {
 return trim(str).replace(/\s\s+/g, ' ')
}

func Dasherize(str string) (string) {
  return trim(str).replace(/([A-Z])/g, "-$1").replace(/[-_\s]+/g, "-").toLowerCase()
}

func Decapitalize(str string) (string) {
  str = makeString(str)
  return str.charAt(0).toLowerCase() + str.slice(1)
}

func EndsWith(str string, ends, position int) (string) {
  str = makeString(str)
  ends = "" + ends
  if (typeof position == "undefined") {
    position = str.length - ends.length
  } else {
    position = Math.min(toPositive(position), str.length) - ends.length
  }
  return position >= 0 && str.indexOf(ends, position) === position
}

func Humanize(str string) (string) {
  return capitalize(trim(underscored(str).replace(/_id$/, "").replace(/_/g, " ")))
}

func IsBlank(str string) {
  return (/^\s*$/).test(makeString(str))
}

func Levenshtein(str1 string, str2 string) {

  str1 = makeString(str1)
  str2 = makeString(str2)

  // Short cut cases
  if (str1 === str2) return 0
  if (!str1 || !str2) return Math.max(str1.length, str2.length)

  // two rows
  var prevRow = new Array(str2.length + 1)

  // initialise previous row
  for (var i = 0; i < prevRow.length; ++i) {
    prevRow[i] = i
  }

  // calculate current row distance from previous row
  for (i = 0; i < str1.length; ++i) {
    var nextCol = i + 1

    for (var j = 0; j < str2.length; ++j) {
      var curCol = nextCol

      // substution
      nextCol = prevRow[j] + ( (str1.charAt(i) === str2.charAt(j)) ? 0 : 1 )
      // insertion
      var tmp = curCol + 1
      if (nextCol > tmp) {
        nextCol = tmp
      }
      // deletion
      tmp = prevRow[j + 1] + 1
      if (nextCol > tmp) {
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
  if (str == null) return []
  return String(str).split(/\r\n?|\n/)
}

func Pred(str string) {
  return adjacent(str, -1)
}

func Prune(str string, length int, pruneStr string) (string) {
  str = makeString(str);
  length = ~~length;
  pruneStr = pruneStr != null ? String(pruneStr) : "..."

  if (str.length <= length) return str

  var tmpl = function(c) {
      return c.toUpperCase() !== c.toLowerCase() ? 'A' : ' '
    },
    template = str.slice(0, length + 1).replace(/.(?=\W*\w*$)/g, tmpl) // 'Hello, world' -> 'HellAA AAAAA'

  if (template.slice(template.length - 2).match(/\w\w/)) {
    template = template.replace(/\s*\S+$/, "")
  } else {
    template = rtrim(template.slice(0, template.length - 1))
  }

  return (template + pruneStr).length > str.length ? str : str.slice(0, template.length) + pruneStr;
}

func Repeat(str string, qty int, separator) {
  str = makeString(str)

  qty = ~~qty

  // using faster implementation if separator is not needed;
  if (separator == null) {
    return strRepeat(str, qty)
  }

  // this one is about 300x slower in Google Chrome
  /*eslint no-empty: 0*/
  for (var repeat = []; qty > 0; repeat[--qty] = str) {}
  return repeat.join(separator)
}

func Reverse(str string) {
  return chars(str).reverse().join("")
}

func Slugify(str string) {
  return trim(dasherize(cleanDiacritics(str).replace(/[^\w\s-]/g, '-').toLowerCase()), '-')
}

func Splice(str string, i, howmany, substr) {
  var arr = chars(str)
  arr.splice(~~i, ~~howmany, substr);
  return arr.join("")
}

func StartsWith(str, starts, position) {
  str = makeString(str)
  starts = "" + starts
  position = position == null ? 0 : Math.min(toPositive(position), str.length)
  return str.lastIndexOf(starts, position) === position
}

func Succ(str string) {
 return adjacent(str, 1)
}

func SwapCase(str string) {
  return makeString(str).replace(/\S/g, function(c) {
    return c === c.toUpperCase() ? c.toLowerCase() : c.toUpperCase()
  })
}

func Titleize(str string) (string) {
  return makeString(str).toLowerCase().replace(/(?:^|\s|-)\S/g, function(c) {
    return c.toUpperCase()
  })
}

func Truncate(str string, length int, truncateStr string) (string) {
  str = makeString(str);
  truncateStr = truncateStr || "..."
  length = ~~length
  return str.length > length ? str.slice(0, length) + truncateStr : str
}

func Underscored(str string) (string) {
  return trim(str).replace(/([a-z\d])([A-Z]+)/g, "$1_$2").replace(/[-\s]+/g, "_").toLowerCase()
}

func Unquote(str string, quoteChar) (string) {
  quoteChar = quoteChar || "\"";
  if (str[0] === quoteChar && str[str.length - 1] === quoteChar)
    return str.slice(1, str.length - 1)
  else return str
}

func Words(str string, delimiter) {
  if (isBlank(str)) return []
  return trim(str, delimiter).split(delimiter || /\s+/)
}
