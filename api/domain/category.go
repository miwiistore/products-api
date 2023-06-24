package domain

type Category struct {
	CategoryID int16  `json:"category_id,omitempty"`
	Name       string `json:"name"`
}
