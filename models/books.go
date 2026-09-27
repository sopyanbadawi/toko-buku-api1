package models

type Books struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
	YEAR   int    `json:"year"`
	STOCK  int    `json:"stock"`
}
