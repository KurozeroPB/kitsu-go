package kitsu

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
)

// Manga struct with all the manga data from kitsu
type Manga struct {
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
		Subtype           string   `json:"subtype"`
		Status            string   `json:"status"`
		TBA               string   `json:"tba"`
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

		ChapterCount  int    `json:"chapterCount"`
		VolumeCount   int    `json:"volumeCount"`
		Serialization string `json:"serialization"`
		MangaType     string `json:"mangaType"`
	} `json:"attributes"`
}

// SearchManga search for a manga on kitsu.io
// query being the manga to search for
// offset being the page offset
func SearchManga(query string, offset int) (*Manga, error) {
	uri := fmt.Sprintf("%s/manga?filter[text]=%s&page[offset]=%v", baseURL, query, offset)
	parJSON, e := gabs.ParseJSON(get(uri))
	if e != nil {
		return nil, e
	}
	manga := parJSON.Path("data").Data().([]interface{})
	resJSON, er := json.Marshal(manga[0])
	if er != nil {
		return nil, er
	}
	man := new(Manga)
	err := json.Unmarshal(resJSON, &man)
	if err != nil {
		return nil, err
	}
	return man, nil
}

// MangaByID holds the data from searching a manga by the id
type MangaByID struct {
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
		Subtype           string   `json:"subtype"`
		Status            string   `json:"status"`
		TBA               string   `json:"tba"`
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

		ChapterCount  int    `json:"chapterCount"`
		VolumeCount   int    `json:"volumeCount"`
		Serialization string `json:"serialization"`
		MangaType     string `json:"mangaType"`
	} `json:"attributes"`
}

// GetManga will fetch a manga with the given id from kitsu.io
// id of course being the id
func GetManga(id int) (*MangaByID, error) {
	uri := fmt.Sprintf("%s/manga/%v", baseURL, id)
	parJSON, e := gabs.ParseJSON(get(uri))
	if e != nil {
		return nil, e
	}
	manga := parJSON.Path("data").Data().(map[string]interface{})
	resJSON, er := json.Marshal(manga)
	if er != nil {
		return nil, er
	}
	man := new(MangaByID)
	err := json.Unmarshal(resJSON, &man)
	if err != nil {
		return nil, err
	}
	return man, nil
}
