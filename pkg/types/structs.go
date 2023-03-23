package types

import validation "github.com/go-ozzo/ozzo-validation"

type ResponseStruc struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type FilterStruc struct {
	ID          uint
	Name        *string
	Author      *string
	Publication *string
}
type UpdateStruc struct {
	ID          uint
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (updateStruc UpdateStruc) Validate() error {
	return validation.ValidateStruct(&updateStruc,
		validation.Field(&updateStruc.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&updateStruc.Author, validation.Required, validation.Length(3, 50)),
		validation.Field(&updateStruc.Publication),
	)
}