package magic

import (
	"log"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID       bson.ObjectId `json:"_id"`
	Author   string        `json:"author"`
	Likes    []string      `json:"likes"`
	Date     string        `json:"date"`
	Time     string        `json:"time"`
	Text     string        `json:"text"`
	Cat      string        `json:"cat"`
	Keyword  string        `json:"keyword"`
	Comments []string      `json:"comments"`
	File     string        `json:"file"`
}

func PreviewProcess(user string) (resp []News, flag bool) {

	var result []News
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {

		log.Print("\n!!!!-- DB connection error:")
		log.Print(err)
		log.Print("\n")
		return result, false
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("news").C("news")

	err = c.Find(bson.M{"cat": "freshnews"}).All(&result)

	if err != nil {

		log.Print(err)
		log.Print("\n")
		return result, false
	} else {

		return result, true
	}
}
