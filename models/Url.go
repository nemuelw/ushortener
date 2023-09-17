package models

type Url struct {
	URLHash string `json:"url_hash" gorm:"primary_key"`
	LongURL string `json:"long_url"`
}
