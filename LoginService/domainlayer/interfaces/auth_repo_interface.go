package interfaces

// The provider must implement this interface in to provide the functionality

import (
	"dcg_demo/LoginService/domainlayer/models"
)

type AuthRepoInterface interface {
	GetSaltPassword(userName string) (string, error)
	StoreJWTToken(userName string, token string) error
	SetUserLogout(userName string) error
	GetJWTToken(token string) (string, error)
	CreateUserProfile(userData models.DMSignupModel) error
}
