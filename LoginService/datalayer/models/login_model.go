package models

type Remove_LoginModel struct {
	Name      string   `json:"name" bson:"name"`
	Role      string   `json:"role" bson:"role"`
	Token     string   `json:"token" bson:"token"`
	NickNames []string `json:"nick_name" bson:"nick_name"`
}

type Remove_ProfileModel struct {
	UserName   string `json:"user_name" bson:"user_name"`
	SaltPasswd string `json:"salt_passwd" bson:"salt_passwd"`
}

type AuthModel struct {
	UserName string `json:"user_name" bson:"user_name"`
	JwtToken string `json:"jwt_token" bson:"jwt_token"`
}
