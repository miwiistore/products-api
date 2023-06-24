package domain

type Picture struct {
	PictureID int64  `json:"picture_id,omitempty"`
	Url       string `json:"url"`
	Order     int16  `json:"order"`
}
