package model

import (
	"strings"
	"time"
)

type Course struct {
	ID        string    `bson:"_id" json:"_id" form:"_id"`
	Name      string    `bson:"name" json:"name" form:"name"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt" form:"createdAt"`
}

func (c *Course) CorrectCourseName() {
	c.Name = strings.ToUpper(c.Name)
	c.Name = strings.ReplaceAll(c.Name, " ", "")
	c.Name = strings.ReplaceAll(c.Name, "-", "")
	c.Name = strings.ReplaceAll(c.Name, ".", "")
	c.Name = strings.ReplaceAll(c.Name, ",", "")
	c.Name = strings.ReplaceAll(c.Name, "_", "")
}
