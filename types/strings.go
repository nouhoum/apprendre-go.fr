package main

import "fmt"

func main() {
  fmt.Println("Hello Go strings")

  fmt.Println("Go " + "learn " + " Go language.")

  fmt.Println(len("Hello"))
  fmt.Println(len("Hello Go"))

  fmt.Println("Hello Go strings"[2]) // Affiche le code correspondant au caractère l
  fmt.Println("Hello Go strings"[4]) // Affiche le code correspondantau caractère o
}
