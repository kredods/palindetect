package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

func GetSession() (session *mgo.Session, err error) {
	session, err = mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		fmt.Print("error connecting to db")
		return
	}
	return
}

func GetDBName() string {
	return "palindetect"
}
