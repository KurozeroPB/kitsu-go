package kitsu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jeffail/gabs"
)

const (
	baseURL   = "https://kitsu.io/api/edge"
	userAgent = "kitsu.go/v0.0.1 - (github.com/KurozeroPB/kitsu.go)"
)

func executeRequest(request *http.Request, expectedStatus int) []byte {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != expectedStatus {
		panic(fmt.Errorf(
			"Expected status %d; Got %d \nResponse: %#v",
			expectedStatus,
			response.StatusCode,
			buf.String(),
		))
	}
	return buf.Bytes()
}
func newRequest(method string, url string) *http.Request {
	return newUARequest(method, url, userAgent)
}
func newUARequest(method string, url string, ua string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", ua)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")
	return request
}
func safeGET(url string, expectedStatus int) []byte {
	return executeRequest(
		newRequest("GET", url),
		expectedStatus,
	)
}
func get(url string) []byte {
	return safeGET(url, 200)
}

type titleStruct struct {
	EnJp string `json:"en_jp"`
	JaJp string `json:"ja_jp"`
}

type coverImageStruct struct {
	Original string `json:"original"`
	Tiny     string `json:"tiny"`
	Small    string `json:"small"`
	Large    string `json:"large"`
}

type posterImageStruct struct {
	Tiny     string `json:"tiny"`
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Original string `json:"original"`
}

type attStruct struct {
	TBA               string             `json:"tba"`
	AbbreviatedTitles []string           `json:"abbreviatedTitles"`
	AverageRating     string             `json:"averageRating"`
	Status            string             `json:"status"`
	AgeRating         string             `json:"ageRating"`
	Subtype           string             `json:"subtype"`
	CanonicalTitle    string             `json:"canonicalTitle"`
	EpisodeLength     int                `json:"episodeLength"`
	CoverImage        *coverImageStruct  `json:"coverImage"`
	Slug              string             `json:"slug"`
	Titles            *titleStruct       `json:"titles"`
	AgeRatingGuide    string             `json:"ageRatingGuide"`
	StartDate         string             `json:"startDate"`
	EpisodeCount      int                `json:"episodeCount"`
	FavoritesCount    int                `json:"favoritesCount"`
	NSFW              bool               `json:"nsfw"`
	EndDate           string             `json:"endDate"`
	RatingRank        int                `json:"ratingRank"`
	PosterImage       *posterImageStruct `json:"posterImage"`
	Synopsis          string             `json:"synopsis"`
	ShowType          string             `json:"showType"`
	UserCount         int                `json:"userCount"`
	PopularityRank    int                `json:"popularityRank"`
}

type linksStruct struct {
	Self string `json:"self"`
}

// Anime struct with all the anime data from kitsu
type Anime struct {
	ID             string       `json:"id"`
	Type           string       `json:"type"`
	Links          *linksStruct `json:"links"`
	Attributes     *attStruct   `json:"attributes"`
	YoutubeVideoID string       `json:"youtubeVideoId"`
}

// SearchAnime search for an anime on kitsu.io
// query being the anime to search for
// offset being the page offset
func SearchAnime(query string, offset int) (*Anime, error) {
	uri := fmt.Sprintf("%s/anime?filter[text]=%s&page[offset]=%v", baseURL, query, offset)

	parJSON, e := gabs.ParseJSON(get(uri))
	if e != nil {
		return nil, e
	}

	anime := parJSON.Path("data").Data().([]interface{})

	ani := new(Anime)

	resJSON, er := json.Marshal(anime[0]) // Right now I'm doing anime[0] because I have no idea how to handle it when it would return more than 1 result.
	if er != nil {
		return nil, er
	}
	err := json.Unmarshal(resJSON, &ani)
	if err != nil {
		return nil, err
	}
	return ani, nil
}
