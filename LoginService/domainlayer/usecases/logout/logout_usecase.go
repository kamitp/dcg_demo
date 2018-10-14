package logout

import (
	repo "dcg_demo/LoginService/domainlayer/interfaces"
)

type Logout struct {
	repo repo.AuthRepoInterface
}

func (lu Logout) Logout(userName string) error {
	lu.repo.SetUserLogout(userName)
	return nil
}
