package signup

import (
	"fmt"
	repo "dcg_demo/LoginService/domainlayer/interfaces"
	dm "dcg_demo/LoginService/domainlayer/models"

	"golang.org/x/crypto/bcrypt"
)

type Signup struct {
	Repo repo.AuthRepoInterface
}

func (su Signup) Signup(userData dm.DMSignupModel) error {

	if su.Repo == nil {
		return fmt.Errorf("repository is nil")
	}

	saltPasswd, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userData.Password = string(saltPasswd)
	return su.Repo.CreateUserProfile(userData)
}
