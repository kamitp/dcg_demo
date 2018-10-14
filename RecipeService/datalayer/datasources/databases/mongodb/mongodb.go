package mongodb

import (
	config "dcg_demo/RecipeService/config"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Session

func init() {
	session, err := mgo.Dial(config.DBIP + ":27017") // uncomment this code when dockerizing
	//session, err := mgo.Dial("127.0.0.1:27017") // comment this code when dockerizing

	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db = session
}

func GetCollectionFromDB(dbName string, collection string) *mgo.Collection {
	if db != nil {
		col := db.DB(dbName).C(collection)
		return col
	}
	return nil
}
