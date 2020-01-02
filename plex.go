package main

type PlexMessage struct {
	Event    string       `json:"event"`
	IsUser   bool         `json:"user"`
	IsOwner  bool         `json:"owner"`
	Account  PlexAccount  `json:"account"`
	Server   PlexServer   `json:"server"`
	Player   PlexPlayer   `json:"player"`
	Metadata PlexMetadata `json:"metadata"`
}

type PlexAccount struct {
	ID    int    `json:"id"`
	Thumb string `json:"thumb"`
	Title string `json:"title"`
}
type PlexServer struct {
	Title string `json:"title"`
	UUID  string `json:"uuid"`
}
type PlexPlayer struct {
	IsLocal       bool   `json:"local"`
	PublicAddress string `json:"publicAddress"`
	Title         string `json:"title"`
	UUID          string `json:"uuid"`
}

type PlexMetadata struct {
	LibrarySectionType   string `json:"librarySectionType"`
	LibrarySectionID     int    `json:"librarySectionID"`

	RatingKey            string `json:"ratingKey"`
	ParentRatingKey      string `json:"parentRatingKey"`
	GrandparentRatingKey string `json:"grandparentRatingKey"`

	Key                  string `json:"key"`
	ParentKey            string `json:"parentKey"`
	GrandparentKey       string `json:"grandparentKey"`

	Title                string `json:"title"`
	ParentTitle          string `json:"parentTitle"`
	GrandparentTitle     string `json:"grandparentTitle"`

	Thumb                string `json:"thumb"`
	ParentThumb          string `json:"parentThumb"`
	GrandparentThumb     string `json:"grandparentThumb"`

	Art                  string `json:"art"`
	ParentArt            string `json:"parentArt"`
	GrandparentArt       string `json:"grandparentArt"`

	GUID                 string `json:"guid"`
	Type                 string `json:"type"`

	Summary              string `json:"summary"`
	Index                int    `json:"index"`
	ParentIndex          int    `json:"parentIndex"`
	RatingCount          int    `json:"ratingCount"`

	AddedAt              int    `json:"addedAt"`
	UpdatedAt            int    `json:"updatedAt"`
}
