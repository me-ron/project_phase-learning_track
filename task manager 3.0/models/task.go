package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	User 		User `json:"user"`
   }