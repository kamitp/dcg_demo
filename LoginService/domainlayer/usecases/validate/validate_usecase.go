package validate

import (
	"fmt"
	std "dcg_demo/LoginService/domainlayer/constants"
	repo "dcg_demo/LoginService/domainlayer/interfaces"
	"dcg_demo/LoginService/utilities"
)

type Validate struct {
	Repo repo.AuthRepoInterface
}

func (vl Validate) Validate(jwtToken string) error {

	if vl.Repo == nil {
		return fmt.Errorf("internal error")
	}
	token, err := vl.Repo.GetJWTToken(jwtToken)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	_, err = utilities.ValidateJWTToken(token, std.JwtSecrete)
	return err
}
