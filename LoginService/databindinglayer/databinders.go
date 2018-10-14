package databindinglayer

import (
	dlmodel "dcg_demo/LoginService/datalayer/models"
	dmmodels "dcg_demo/LoginService/domainlayer/models"
	plmodel "dcg_demo/LoginService/presentationlayer/models"
)

func BindPLTODMSignup(userDate *plmodel.PLSignupModel) dmmodels.DMSignupModel {

	return dmmodels.DMSignupModel{
		UserName: userDate.UserName,
		Password: userDate.Password,
		Gender:   userDate.Gender,
		Mobile:   userDate.Mobile,
		Email:    userDate.Email}
}

func BindDMToDLSignup(userDate *dmmodels.DMSignupModel) dlmodel.DLSignupModel {

	return dlmodel.DLSignupModel{
		UserName:   userDate.UserName,
		SaltPasswd: userDate.Password,
		Gender:     userDate.Gender,
		Mobile:     userDate.Mobile,
		Email:      userDate.Email}
}
