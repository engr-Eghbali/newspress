package magic

import (
	"log"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
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

func Like(user string, newsID string) (flag bool) {

	var temp News
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {

		log.Print("\n!!!!-- DB connection error:")
		log.Print(err)
		log.Print("\n")
		return false
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("news").C("news")
	err = c.FindId(bson.ObjectIdHex(newsID)).One(&temp)
	for i, _ := range temp.Likes {

		if temp.Likes[i] == user {
			return false
		}

	}

	idQuerier := bson.M{"_id": bson.ObjectIdHex(newsID)}
	change := bson.M{"$push": bson.M{"likes": user}}
	err = c.Update(idQuerier, change)
	if err != nil {
		log.Print("failed to like:")
		log.Print(err)
		return false
	} else {

		return true
	}

}

func DisLike(user string, newsID string) (flag bool) {
	var temp News
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {

		log.Print("\n!!!!-- DB connection error:")
		log.Print(err)
		log.Print("\n")
		return false
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("news").C("news")

	err = c.FindId(bson.ObjectIdHex(newsID)).One(&temp)
	for i, _ := range temp.Likes {

		if temp.Likes[i] == user {

			idQuerier := bson.M{"_id": bson.ObjectIdHex(newsID)}
			change := bson.M{"$pull": bson.M{"likes": user}}
			err = c.Update(idQuerier, change)
			if err != nil {
				log.Print("failed to like:")
				log.Print(err)
				return false
			} else {

				return true
			}
		}

	}
	return false

}
