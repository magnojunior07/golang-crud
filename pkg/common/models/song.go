package models

import "gorm.io/gorm"

type CharlieBrownJrSongs struct {
		gorm.Model
		SongName 			string `json:"songName"`
		Album 				string `json:"qtd_titles"`
		Duration 			string `json:"duration"`
		LinkSpotify 	string `json:"linkSpotify"`
		LinkYouTube 	string `json:"linkYouTube"`
}

