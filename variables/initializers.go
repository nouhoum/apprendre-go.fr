package main

import "fmt"

func main() {
  var (
    likes int = 12
    name string = "John Doe"
    ok bool = true
  )

  fmt.Println(likes)
  fmt.Println(name)
  fmt.Println(ok)

  friends, city := 12, "Paris"
  var age = 23
  fmt.Println(friends)
  fmt.Println(city)
  fmt.Println(age)

  artist := "Otis"
  fmt.Println(artist)
}
