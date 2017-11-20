package kitsu

import (
	"encoding/json"
	"fmt"
	"net/url"

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
		ProExpiresAt        string   `json:"proExpiresAt"`
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
	newQuery := url.QueryEscape(query)
	uri := fmt.Sprintf("%s/users?filter[name]=%s", baseURL, newQuery)
	byt, er := get(uri)
	if er != nil {
		return nil, er
	}
	parJSON, e := gabs.ParseJSON(byt)
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

// UserByID struct with all the user's data from kitsu
type UserByID struct {
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
		ProExpiresAt        string   `json:"proExpiresAt"`
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
	Relationships struct{} `json:"relationships"`
}

// GetUser get a user by his/her id from kitsu.io
// id of course being the id
func GetUser(id int) (*UserByID, error) {
	newQuery := url.QueryEscape(fmt.Sprintf("%d", id))
	uri := fmt.Sprintf("%s/users/%s", baseURL, newQuery)
	byt, er := get(uri)
	if er != nil {
		return nil, er
	}
	parJSON, e := gabs.ParseJSON(byt)
	if e != nil {
		return nil, e
	}
	user := parJSON.Path("data").Data().(map[string]interface{})
	resJSON, er := json.Marshal(user)
	if er != nil {
		return nil, er
	}
	u := new(UserByID)
	err := json.Unmarshal(resJSON, &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Stats struct with all the stats from the user
type Stats struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		Kind      string `json:"kind"`
		StatsData struct {
			Total         int      `json:"total"`
			TotalMedia    int      `json:"total_media"`
			Activity      []string `json:"activity"`
			AllCategories struct {
				War   int `json:"war"`
				Asia  int `json:"asia"`
				Cops  int `json:"cops"`
				Idol  int `json:"idol"`
				Mars  int `json:"mars"`
				Navy  int `json:"navy"`
				Past  int `json:"past"`
				Alien int `json:"alien"`
				Angel int `json:"angel"`
				Angst int `json:"angst"`
				China int `json:"china"`
				Crime int `json:"crime"`
				Deity int `json:"deity"`
				Demon int `json:"demon"`
			} `json:"all_categories"`
			AllTime struct {
				TotalTime     int `json:"total_time"`
				TotalMedia    int `json:"total_media"`
				TotalProgress int `json:"total_progress"`
			} `json:"all_time"`
			AllYears struct{} `json:"all_years"`
		} `json:"statsData"`
	} `json:"attributes"`
	Relationships struct{} `json:"relationships"`
}

// GetStats get the stats of a user by his/her id from kitsu.io
// id of course being the id
func GetStats(id int) (*Stats, error) {
	newQuery := url.QueryEscape(fmt.Sprintf("%d", id))
	uri := fmt.Sprintf("%s/stats/%v", baseURL, newQuery)
	byt, er := get(uri)
	if er != nil {
		return nil, er
	}
	parJSON, e := gabs.ParseJSON(byt)
	if e != nil {
		return nil, e
	}
	stats := parJSON.Path("data").Data().(map[string]interface{})
	resJSON, er := json.Marshal(stats)
	if er != nil {
		return nil, er
	}
	sts := new(Stats)
	err := json.Unmarshal(resJSON, &sts)
	if err != nil {
		return nil, err
	}
	return sts, nil
}
