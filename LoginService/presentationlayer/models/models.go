package models

type PLSignupModel struct {
	UserName string `json:"user_name"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Gender   string `json:"gender"`
}
