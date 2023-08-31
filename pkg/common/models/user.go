package models

type User struct{
    ID      uint       `gorm:"primarykey"`
    UserId  uint       `gorm:"integer"`
    Segments []*Segment `gorm:"many2many:user_segment;"`
}


type GetUserSegmentsRequestBody struct {
	UserId uint `json:"user_id"`
}

type GetUserSegmentsResponseBody struct {
	Type    string  `json:"type"`
	UserId  uint    `json:"user_id"`
	Slugs  []string `json:"slugs"`
}

type ChangeUserSegmentsRequestBody struct {
	SlugsToAdd    []string `json:"slugs_to_add"`
	SlugsToDelete []string `json:"slugs_to_delete"`
	UserId        uint     `json:"user_id"`
	Expires       *string  `json:"expires"`
}