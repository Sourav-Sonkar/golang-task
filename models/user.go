package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	DOB         string             `json:"dob,omitempty"`
	Address     string             `json:"address,omitempty"`
	Description string             `json:"description,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
}
