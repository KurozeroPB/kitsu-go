package kitsu

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Jeffail/gabs"
)

// Character struct with all the character data from kitsu
type Character struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		MalID       int    `json:"malId"`
		Description string `json:"description"`
		Image       struct {
			Original string `json:"original"`
		}
	} `json:"attributes"`
}

// SearchCharacter search for a character on kitsu.io
// query being the character to search for
func SearchCharacter(query string) (*Character, error) {
	newQuery := url.QueryEscape(query)
	uri := fmt.Sprintf("%s/characters?filter[name]=%s", baseURL, newQuery)
	byt, er := get(uri)
	if er != nil {
		return nil, er
	}
	parJSON, e := gabs.ParseJSON(byt)
	if e != nil {
		return nil, e
	}
	character := parJSON.Path("data").Data().([]interface{})
	resJSON, er := json.Marshal(character[0])
	if er != nil {
		return nil, er
	}
	char := new(Character)
	err := json.Unmarshal(resJSON, &char)
	if err != nil {
		return nil, err
	}
	return char, nil
}
