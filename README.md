# kitsu.go
__Interact with the kitsu.io api using Go__</br>

## Install
`go get github.com/KurozeroPB/kitsu.go`

## Usage
Simple example usage:
```go
package main

import (
  "fmt"

  "github.com/KurozeroPB/kitsu.go"
)

func main() {
  // Example to search for an anime
  anime, e := kitsu.SearchAnime("fate/apocrypha", 0)
  if e != nil {
    fmt.Println(e)
    return
  }
  fmt.Println(anime.Attributes.PosterImage.Original)
}
```

## Docs

#### SearchAnime(query, offset)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| query     | string        | The anime you want to search
| offset    | int        | Page offset

#### SearchManga(query, offset)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| query     | string        | The manga you want to search
| offset    | int        | Page offset

#### SearchCharacter(query)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| query     | string        | The name of the character you want to search for

#### SearchProducer(query)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| query     | string        | The name of the producer you want to search for

#### SearchUser(query)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| query     | string        | The name of the user you want to search for

#### SearchDrama(query)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| query     | string        | The drama you want to search
> There are currently no dramas on kitsu so this will return an error until they add dramas to the website.

#### GetAnime(id)
| Parameter | Type          | Description |
|-----------|:-------------:|-------------|
| id        | int           | The id of the anime you want to get

## TODO
1. ~~Manga search~~
2. ~~Character search~~
3. ~~Producers search~~
4. ~~Users search~~
5. ~~Drama search~~
6. ~~Get anime by id~~
7. Get manga by id
8. Get user by id
9. Get user stats by id

<!--
Get manga by id:
https://kitsu.io/api/edge/manga/id
http://docs.kitsu.apiary.io/#reference/media/manga/fetch-resource

Get user by id:
https://kitsu.io/api/edge/users/id
http://docs.kitsu.apiary.io/#reference/users/users/fetch-resource

Get user stats by id:
https://kitsu.io/api/edge/stats/id
http://docs.kitsu.apiary.io/#reference/users/stats/fetch-resource
-->
