package helpers

import "os"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func GetInput(f string) string {
  dat, err := os.ReadFile(f)
  check(err)
  return string(dat)
}

func Sum(a []int) (total int){
  total = 0
  for _, v := range a {
    total += v
  }
  return
}
