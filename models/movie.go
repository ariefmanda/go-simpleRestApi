package models

import (
	"log"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

const (
	DB= "hello"
	COLLECTION = "movies"
)

type (
	Movie struct {
		ID          bson.ObjectId `bson:"_id" json:"id"`
		Name        string        `bson:"name" json:"name"`
		CoverImage  string        `bson:"cover_image" json:"cover_image"`
		Description string        `bson:"description" json:"description"`
	}
)

func Connect() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(DB)
	return
}

func FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}
 
func FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}
 
func Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}
 
func Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}
 
func Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}