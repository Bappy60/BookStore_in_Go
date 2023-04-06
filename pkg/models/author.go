package models

type Author struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"author_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Age   int    `json:"author_age"`
	Books []Book `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

