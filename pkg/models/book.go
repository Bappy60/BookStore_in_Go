package models

type Book struct {
	ID              uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name            string `gorm:"" json:"name"`
	PublicationYear int    `json:"publication_year"`
	NumberOfPages   int    `json:"number_of_pages"`
	AuthorID        uint   `gorm:"index" json:"author_id"` // foreign key
	Author          Author `gorm:"foreignKey:AuthorID"`
	Publication     string `json:"publication"`
}
