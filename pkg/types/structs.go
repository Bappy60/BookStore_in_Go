package types

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
