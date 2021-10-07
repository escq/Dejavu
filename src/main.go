package main

import (
  "fmt"
  "os/exec"
  "os"
)

func main() {
  const words = "Hello world\nHello world\nHello world"
  clear()
  fmt.Println(words)
}

func clear() {
  cmdName := "clear"
  cmd := exec.Command(cmdName)
  cmd.Stdout = os.Stdout
  cmd.Run()

}
