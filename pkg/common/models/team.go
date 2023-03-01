package models

import "gorm.io/gorm"

type Team struct {
		gorm.Model
		Name 	string `json:"name"`
		Qtd_titles string `json:"qtd_titles"`
}

