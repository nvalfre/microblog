package models

type User struct {
	ID        string   `json:"id" bson:"id"`
	Username  string   `json:"username" bson:"username"`
	Followers []string `json:"followers" bson:"followers"`
	Following []string `json:"following" bson:"following"`
}
