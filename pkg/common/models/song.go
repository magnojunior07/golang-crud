package models

import "gorm.io/gorm"

type CharlieBrownJrSongs struct {
   gorm.Model
   SongName	string `json:"song_name"`
   Album	string `json:"album"`
   Duration	string `json:"duration"`
   LinkSpotify	string `json:"link_spotify"`
   LinkYouTube	string `json:"link_you_tube"`
}

