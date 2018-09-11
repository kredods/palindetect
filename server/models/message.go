package models

import(
	"gopkg.in/mgo.v2/bson"
)
type Message struct {
	Id  bson.ObjectId `json:"id"`
	Body  string `json:"body"`
	IsPalindrome bool `json:"isPalindrome"`
}