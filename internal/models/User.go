package model

type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
