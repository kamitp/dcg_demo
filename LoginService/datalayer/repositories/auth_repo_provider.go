package repositories

import (
	"fmt"

	authInterface "dcg_demo/LoginService/domainlayer/interfaces"

	binders "dcg_demo/LoginService/databindinglayer"
	mongo "dcg_demo/LoginService/datalayer/datasources/databases/mongodb"
	dlmodel "dcg_demo/LoginService/datalayer/models"
	std "dcg_demo/LoginService/domainlayer/constants"
	dmmodel "dcg_demo/LoginService/domainlayer/models"
	utilities "dcg_demo/LoginService/utilities"

	"gopkg.in/mgo.v2/bson"
)

type AuthRepoProvider struct {
	authInterface.AuthRepoInterface
}

func (auth AuthRepoProvider) GetSaltPassword(userName string) (string, error) {

	var profiles []dlmodel.DLSignupModel
	col := mongo.GetCollectionFromDB(std.LoginServiceDB, std.ProfileCollection)
	if col == nil {
		return "", fmt.Errorf("collection could not be found")
	}
	query := bson.M{"user_name": userName}
	err := col.Find(query).All(&profiles)
	if err != nil {
		return "", err
	}
	if len(profiles) == 0 {
		return "", fmt.Errorf("user not found")
	}
	return profiles[0].SaltPasswd, nil
}

func (auth AuthRepoProvider) StoreJWTToken(userName string, token string) error {

	col := mongo.GetCollectionFromDB(std.LoginServiceDB, std.AuthCollection)

	var authModel dlmodel.AuthModel
	err := col.Find(bson.M{"user_name": userName}).One(&authModel)

	if err != nil && err.Error() == "not found" {
		// Insert document
		return col.Insert(&dlmodel.AuthModel{UserName: userName, JwtToken: token})
	} else {
		//Update document
		authModel.JwtToken = token
		err = col.Update(bson.M{"user_name": userName}, bson.M{"$set": bson.M{"jwt_token": token}})
	}
	return nil
}

func (auth AuthRepoProvider) SetUserLogout(userName string) error {
	return nil
}

func (auth AuthRepoProvider) ValidateJWTToken(token string) error {
	col := mongo.GetCollectionFromDB(std.LoginServiceDB, std.AuthCollection)
	var authModel dlmodel.AuthModel
	err := col.Find(bson.M{"jwt_token": token}).One(&authModel)
	if err != nil && err.Error() == "not found" {
		return err
	} else {
		_, err = utilities.ValidateJWTToken(token, std.JwtSecrete)
	}
	return err
}

func (auth AuthRepoProvider) GetJWTToken(token string) (string, error) {
	col := mongo.GetCollectionFromDB(std.LoginServiceDB, std.AuthCollection)
	var authModel dlmodel.AuthModel
	err := col.Find(bson.M{"jwt_token": token}).One(&authModel)
	return authModel.JwtToken, err
}

func (auth AuthRepoProvider) CreateUserProfile(userData dmmodel.DMSignupModel) error {
	col := mongo.GetCollectionFromDB(std.LoginServiceDB, std.ProfileCollection)
	dlmodel := binders.BindDMToDLSignup(&userData)
	err := col.Insert(&dlmodel)
	if err != nil {
		return err
	}

	return nil
}
