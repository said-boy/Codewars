package code6Kyu

// Link Codewars : https://www.codewars.com/kata/54da539698b8a2ad76000228

func IsValidWalk(walk []rune) bool {
  x := 0
  y := 0
  
  if len(walk) != 10 {
    return false
  }
  
  for _, v := range walk {
    switch string(v) {
      case "n" :
        y++
      case "s" :
        y--
      case "e" :
        x++
      case "w" :
        x--
    }
  }
  
  if x == 0 && y == 0 {
    return true
  }
  return false
  
}