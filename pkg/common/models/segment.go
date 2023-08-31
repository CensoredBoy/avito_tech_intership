package models

import(
    "time"
)

type Segment struct {
    ID         uint    `gorm:"primarykey"`
    Slug       string  `gorm:"string"`
    Users      []*User `gorm:"many2many:user_segment;"`
}

type SegmentRequestBody struct {
    Slug       string `json:"slug"`
}

type SegmentAssociation struct {
    SegmentId uint
    Expires   *time.Time `gorm:"datetime:timestamp"`
}