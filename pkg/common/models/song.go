package models

import (
	"time"

	"gorm.io/gorm"
)

type CharlieBrownJrSongs struct {
   Id int `json:"id" gorm:"primaryKey"`
   CreatedAt time.Time `json:"created_at"`
   UpdatedAt time.Time `json:"updated_at"`
   DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
   SongName	string `json:"song_name"`
   Album	string `json:"album"`
   Duration	string `json:"duration"`
   LinkSpotify	string `json:"link_spotify"`
   LinkYouTube	string `json:"link_you_tube"`
}

