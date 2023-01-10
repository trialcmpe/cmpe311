package model

import "time"

type Note struct {
	ID          string    `bson:"_id" json:"_id" form:"_id"`
	Description string    `bson:"description" json:"description" form:"description"`
	Course      string    `bson:"course" json:"course" form:"course"`
	Creator     string    `bson:"creator" json:"creator" form:"creator"`
	URL         string    `bson:"url" json:"url" form:"url"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt" form:"createdAt"`
}
