package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Music struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Section    string             `json:"section,omitempty" bson:"section,omitempty" validate:"required"`
	Titulo     string             `json:"titulo,omitempty" bson:"titulo,omitempty"`
	IdExternal string             `json:"id_external,omitempty" bson:"id_external,omitempty" validate:"required"`
	Thumbnail  string             `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
	Url        string             `json:"url,omitempty" bson:"url,omitempty"`
	IsSinged   bool               `json:"isSinged" bson:"isSinged"`
}
