package model

type User struct {
	GUID      string `json:"guid" bson:"guid" form:"guid"`
	Name      string `json:"name" bson:"name" form:"name"`
	Password  string `json:"password,omitempty" bson:"password" form:"password"`
	Email     string `json:"email" bson:"email" form:"email" `
	Group     string `json:"group" bson:"group" form:"group"`
	Picture   string `json:"picture" bson:"picture"`
	Address   string `json:"address,omitempty" bson:"address" form:"address" `
	IsSuspend bool   `json:"is_suspend" bson:"is_suspend" form:"is_suspend"`
	CreatedAt int64  `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at" form:updated_at""`
}
