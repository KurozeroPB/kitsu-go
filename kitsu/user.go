package kitsu

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
)

// User struct with all the user's data from kitsu
type User struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt           string   `json:"createdAt"`
		UpdatedAt           string   `json:"updatedAt"`
		Name                string   `json:"name"`
		PastNames           []string `json:"pastNames"`
		Slug                string   `json:"slug"`
		About               string   `json:"about"`
		Location            string   `json:"location"`
		WaifuOrHusbando     string   `json:"waifuOrHusbando"`
		FollowersCount      int      `json:"followersCount"`
		FollowingCount      int      `json:"followingCount"`
		LifeSpentOnAnime    int      `json:"lifeSpentOnAnime"`
		Birthday            string   `json:"birthday"`
		Gender              string   `json:"gender"`
		CommentsCount       int      `json:"commentsCount"`
		FavoritesCount      int      `json:"favoritesCount"`
		LikesGivenCount     int      `json:"likesGivenCount"`
		ReviewsCount        int      `json:"reviewsCount"`
		LikesReceivedCount  int      `json:"likesReceivedCount"`
		PostsCount          int      `json:"postsCount"`
		RatingsCount        int      `json:"ratingsCount"`
		MediaReactionsCount int      `json:"mediaReactionsCount"`
		ProExpiresAt        int      `json:"proExpiresAt"`
		Title               string   `json:"title"`
		ProfileCompleted    bool     `json:"profileCompleted"`
		FeedCompleted       bool     `json:"feedCompleted"`
		Website             string   `json:"website"`
		Avatar              struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Medium   string `json:"medium"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"avatar"`
		CoverImage struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"coverImage"`
		RatingSystem string `json:"ratingSystem"`
		Theme        string `json:"theme"`
		FacebookID   string `json:"facebookId"`
	} `json:"attributes"`
	Relationships struct{} `json:"relationships"` // Relationships can contain different stuff for everyone
}

// SearchUser search for a user on kitsu.io
// query being the user to search for
func SearchUser(query string) (*User, error) {
	uri := fmt.Sprintf("%s/users?filter[name]=%s", baseURL, query)
	parJSON, e := gabs.ParseJSON(get(uri))
	if e != nil {
		return nil, e
	}
	user := parJSON.Path("data").Data().([]interface{})
	resJSON, er := json.Marshal(user[0])
	if er != nil {
		return nil, er
	}
	u := new(User)
	err := json.Unmarshal(resJSON, &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
