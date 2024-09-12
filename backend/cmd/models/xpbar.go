package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Xpbar struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	UserName   string             `bson:"user_name" json:"user_name"`
	Name       string             `bson:"name" json:"name"`
	CurrentXp  int                `bson:"currenp_xp" json:"current_xp"`
	RequaredXp int                `bson:"requared_xp" json:"requared_xp"`
}
