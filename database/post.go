package database

type Post struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Body   string
	UserID uint
}
