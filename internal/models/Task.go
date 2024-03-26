package model

type Task struct {
	Id        string `bson:"_id"`
	Name      string `bson:"name"`
	Completed bool   `bson:"completed"`
}
