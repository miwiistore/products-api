package domain

type Picture struct {
	PictureID int64  `json:"picture_id"`
	ProductID int64  `json:"product_id"`
	Base64    string `json:"base_64"`
	Order     int16  `json:"order"`
	Url       string `json:"url"`
}
