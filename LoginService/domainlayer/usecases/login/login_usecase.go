package login

import (
	"fmt"

	std "dcg_demo/LoginService/domainlayer/constants"
	authRepo "dcg_demo/LoginService/domainlayer/interfaces"
	"dcg_demo/LoginService/utilities"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Repo authRepo.AuthRepoInterface
}

func (li Login) Login(userName string, passwd string) (string, error) {

	if li.Repo == nil {
		return "", fmt.Errorf("repository is nil")
	}

	dbSaltPasswd, err := li.Repo.GetSaltPassword(userName)
	err = bcrypt.CompareHashAndPassword([]byte(dbSaltPasswd), []byte(passwd))
	if err != nil {
		return "", err
	}
	token, err := utilities.CreateJWT(userName, std.JwtSecrete)
	if err != nil {
		return "", err
	}
	li.Repo.StoreJWTToken(userName, token)
	return token, nil
}
