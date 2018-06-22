package magic

import (
	"log"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Author   string        `json:"author"`
	Likes    []string      `json:"likes"`
	Date     string        `json:"date"`
	Title    string        `json:"title"`
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

func More(newsID string) (resp News, flag bool) {

	var result News
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

	err = c.FindId(bson.ObjectIdHex(newsID)).One(&result)

	if err != nil {

		log.Print(err)
		log.Print("\n")
		return result, false
	} else {

		return result, true
	}
}

func Cm(newsID string, user string, cm string) (flag bool) {

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
	idQuerier := bson.M{"_id": bson.ObjectIdHex(newsID)}
	change := bson.M{"$push": bson.M{"comments": user + "@" + cm}}
	err = c.Update(idQuerier, change)
	if err != nil {
		log.Print("failed to follow:")
		log.Print(err)
		return false
	} else {
		log.Print("user commented")
		return true
	}
}

func SendNews(user string, title string, keywords string, time string, text string) (flag bool, fileid string) {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {

		log.Print("\n!!!!-- DB connection error:")
		log.Print(err)
		log.Print("\n")
		return false, ""
	} else {

		id := bson.NewObjectId()
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("news").C("news")

		err = c.Insert(&News{ID: id, Author: user, Date: time, Title: title, Text: text, Cat: "freshnews", Keyword: keywords, File: "S:\\newsPress\\server\\files\\" + id.Hex() + ".jpg"})

		if err != nil {

			log.Print(err)
			return false, ""

		} else {
			fileid = id.Hex()
			log.Print("\n news Inserted:")
			return true, id.Hex()
		}

	}
}

func SearchNews(category string, query string) (respose []News, flag bool) {

	var tmp []News
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

	if category == "all" {

		err = c.Find(bson.M{"$or": []bson.M{bson.M{"title": bson.RegEx{".*" + query + ".*", ""}}, bson.M{"text": bson.RegEx{".*" + query + ".*", ""}}}}).All(&result)
	} else {

		log.Print(category)
		err = c.Find(bson.M{"keyword": category}).All(&tmp)
		for _, res := range tmp {
			log.Print(res.Keyword)
			if strings.Contains(res.Text, query) || strings.Contains(res.Title, query) {
				result = append(result, res)
			}
		}

	}

	if err != nil {

		log.Print(err)
		log.Print("\n")
		return result, false
	} else {

		return result, true
	}

}
