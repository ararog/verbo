package verbo

func adjacent(str string, direction int) (string) {
  if len(str) == 0 {
    return ""
  }
  return str[0:1] + string(int(str[len(str) - 1]) + direction)
}

func strRepeat(str string, qty int) (string) {
  if qty < 1 {
    return ""
  }
  result := ""
  for qty > 0 {
    if qty & 1 == 0 {
      result += str
    }
    qty >>= 1
    str += str
  }
  return result
}

func toPositive(number int) (int) {
  if number < 0 {
    return  0
  } else {
    return +number
  }
}
