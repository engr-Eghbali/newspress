package magic

import (
	"log"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	Title string `json:"title"`

	Text string `json:"text"`

	Date string `json:"date"`

	Likes []string `json:"likes"`

	File string `json:"file"`

	Comments []string `json:"off"`

	Author string `json:"author"`

	Keywords string `json:"keywords"`
}

type User struct {
	Username string `json:"userName"`

	Password string `json:"password"`

	Usertype string `json:"userType"`

	Rate int64 `json:"rate"`

	Followers []string `json:"followers"`

	Followings []string `json:"followings"`

	Posted []string `json:"posted"`
}

func LoginProcess(user string, pass string) (resp User, flg bool) {
	var result User
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {

		log.Print("\n!!!!-- DB connection error:")
		log.Print(err)
		log.Print("\n")
		return result, false
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("userInfo").C("users")

	err = c.Find(bson.M{"username": user}).One(&result)

	if err != nil {

		log.Print("\n user login query failed:\n")
		log.Print(err)
		return result, false

	} else {
		if result.Password == pass {

			return result, true

		} else {

			return result, false
		}

	}

}

func SubmitProcess(user string, pass string) (resp User, flg bool) {
	var result User
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {

		log.Print("\n!!!!-- DB connection error:")
		log.Print(err)
		log.Print("\n")
		return result, false
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("userInfo").C("users")

	err = c.Find(bson.M{"username": user}).One(&result)

	if err == nil {

		log.Print("\n User duplicated:\n")
		log.Print(err)
		return result, false

	} else {

		result.Username = user
		result.Password = pass
		result.Usertype = "کاربر جدید"
		result.Rate = 1
		err = c.Insert(result)
		if err != nil {
			log.Print("\n!!!!--submit err:")
			log.Print(err)
			log.Print("\n")
			return result, false
		} else {

			return result, true
		}

	}

}
