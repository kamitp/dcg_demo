package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTokenStatusFromServer(token string) error {

	if len(token) == 0 {
		return fmt.Errorf("Invalid Token")
	}

	fmt.Print(token)

	jsonData := map[string]string{}
	jsonData["jwt_token"] = token
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("http://172.17.0.1:8080/validate", "application/json", bytes.NewBuffer(jsonValue))
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var respData validateResponse
	json.Unmarshal(body, &respData)
	if respData.Status != "success" {
		return fmt.Errorf("Invalid Token")
	}
	return nil
}

type validateResponse struct {
	Status string `json:"status"`
}

// func (rc Recipe) verifyWithAuthService(token string) error {
// 	jsonData := map[string]string{}
// 	jsonValue, _ := json.Marshal(jsonData)
// 	response, err := http.Post("http://localhost:8080/validate", "application/json", bytes.NewBuffer(jsonValue))
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Print(body)
// 	//json.Unmarshal(body, &serviceData)
// 	return nil
// }
