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

 TODO:
 -
1. Manga search
2. Character search 
3. Producers search
4. Drama search (?)

<!--
Character search:
https://kitsu.io/api/edge/characters?filter[name]=
http://docs.kitsu.apiary.io/#reference/characters-&-people/characters/fetch-collection

Producers search:
https://kitsu.io/api/edge/producers?filter[slug]=
http://docs.kitsu.apiary.io/#reference/characters-&-people/producers/fetch-collection

Drama search:
https://kitsu.io/api/edge/drama?filter[text]=
http://docs.kitsu.apiary.io/#reference/media/drama/fetch-collection
-->