package repositories

import (
	net "dcg_demo/RecipeService/datalayer/datasources/network"
	auth "dcg_demo/RecipeService/domainlayer/interfaces"
)

type TokenAuthRepo struct {
	auth.TokenAuthInterface
}

func (rc TokenAuthRepo) GetTokenStatusFromServer(token string) error {
	return net.GetTokenStatusFromServer(token)
}
