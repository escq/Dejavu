package main

import (
  "fmt"
  "os/exec"
  "os"
  "strings"
  "bufio"

  term "github.com/nsf/termbox-go"
)

type cell struct {
  char rune
  fg  term.Attribute
  bg  term.Attribute
}

func main() {

  words :=[4]string{"easy","peasy","lemon","squeezy"}
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("  " + words[0])

  for {
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare(words[0], text) == 0 {
      fmt.Println("Correct word");
      break;
    } else {
      fmt.Println("Incorrect word")
    }
  }
}
