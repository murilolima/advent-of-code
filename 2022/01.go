package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

var max1, max2, max3 = 0, 0, 0

func includeElf(elf int) {
  if (elf > max1) {
    max3 = max2
    max2 = max1
    max1 = elf
  } else if (elf > max2) {
    max3 = max2
    max2 = elf
  } else if (elf > max3) {
    max3 = elf
  }
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  current_elf := 0
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      includeElf(current_elf)
      current_elf = 0
    } else {
      val, _ := strconv.Atoi(line)
      current_elf += val
    }
  }
  includeElf(current_elf)
  fmt.Println(max1 + max2 + max3)
}
