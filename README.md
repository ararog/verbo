# verbo

[![Build Status](https://travis-ci.org/ararog/verbo.svg?branch=master)](https://travis-ci.org/ararog/verbo)

A collection of string utilities for GO. This project is pretty much inspired on [underscore.string](https://github.com/epeli/underscore.string)
project.

## Example

```golang
import (
  "fmt"
  verbo "github.com/ararog/verbo"
)

str := verbo.LeftPad("1", 8, "0")
fmt.Printf("Output: %s\n", str)
```

## Contributing

You can contribute by creating new testcases and fixing code of tests that are
currently failing, the functions Chop, Humanize, Levenstein, Pad, Slugify,
Underscored and Words are not fully functional and requires some additional fixes.

Feel free to fork and create a PR!

## Questions

Feel free to ask questions on stackoverflow using tag "verbo".
