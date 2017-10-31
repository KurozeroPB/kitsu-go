package kitsu

// https://kitsu.io/api/edge/drama?filter[text]=

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
)

/* There are currently no dramas on kitsu so this will return with nothing until they add dramas to the website. */

// Drama struct with all the drama data from kitsu
type Drama struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		Slug      string `json:"slug"`
		Synopsis  string `json:"synopsis"`
		Titles    struct {
			En   string `json:"en"`
			EnJp string `json:"en_jp"`
			JaJp string `json:"ja_jp"`
		} `json:"titles"`
		CanonicalTitle    string   `json:"canonicalTitle"`
		AbbreviatedTitles []string `json:"abbreviatedTitles"`
		AverageRating     string   `json:"averageRating"`
		UserCount         int      `json:"userCount"`
		FavoritesCount    int      `json:"favoritesCount"`
		StartDate         string   `json:"startDate"`
		EndDate           string   `json:"endDate"`
		PopularityRank    int      `json:"popularityRank"`
		RatingRank        int      `json:"ratingRank"`
		AgeRating         string   `json:"ageRating"`
		AgeRatingGuide    string   `json:"ageRatingGuide"`
		SubType           string   `json:"subType"`
		Status            string   `json:"status"`
		PosterImage       struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Medium   string `json:"medium"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"posterImage"`
		CoverImage struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"coverImage"`
		EpisodeCount   int    `json:"episodeCount"`
		EpisodeLength  int    `json:"episodeLength"`
		YoutubeVideoID string `json:"youtubeVideoId"`
		NSFW           bool   `json:"nsfw"`
	} `json:"attributes"`
	Relationships struct{} `json:"relationships"`
}

// SearchDrama search for a drama on kitsu.io
// query being the producer to search for
func SearchDrama(query string) (*Drama, error) {
	uri := fmt.Sprintf("%s/drama?filter[text]=%s", baseURL, query)
	parJSON, e := gabs.ParseJSON(get(uri))
	if e != nil {
		return nil, e
	}
	drama := parJSON.Path("data").Data().([]interface{})
	resJSON, er := json.Marshal(drama[0])
	if er != nil {
		return nil, er
	}
	dra := new(Drama)
	err := json.Unmarshal(resJSON, &dra)
	if err != nil {
		return nil, err
	}
	return dra, nil
}
