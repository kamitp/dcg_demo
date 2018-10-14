package models

type DLSignupModel struct {
	UserName   string `bson:"user_name"`
	SaltPasswd string `bson:"salt_passwd"`
	Email      string `bson:"email"`
	Mobile     string `bson:"mobile"`
	Gender     string `bson:"gender"`
}
