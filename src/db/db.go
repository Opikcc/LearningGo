package db

import (
	mgo "gopkg.in/mgo.v2"
)

var Session *mgo.Session //Database Session object

func init() {
	// initialization code here
	Session, err := mgo.Dial("localhost")
}
func get() {
	//logic for get that uses Session object
}
func add() {
	//logic for add that uses Session object
}
func update() {
	//logic for update that uses Session object
}
func delete() {
	//logic for delete that uses Session object
}
