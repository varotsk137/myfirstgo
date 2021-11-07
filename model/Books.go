package model

type Book struct {
	Bid      uint    `json:"bid" gorm:"primaryKey"`
	ISBN     string  `json:"isbn"`
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	Author   *Author `json:"author"`
}
