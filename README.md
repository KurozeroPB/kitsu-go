# kitsu.go
__Interact with the kitsu.io api using Go__</br>
This project is in very early stages, currently you can only search for anime I'm planning on adding manga search next.

### Install
`go get github.com/KurozeroPB/kitsu.go`

### Usage
Simple example usage:
```go
package main

import (
  "fmt"

  "github.com/KurozeroPB/kitsu.go"
)

func main() {
  anime, e := kitsu.SearchAnime("fate/apocrypha", 0)
  if e != nil {
  fmt.Println(e)
    return
  }
  fmt.Println(anime.Attributes.PosterImage.Original)
}
```

