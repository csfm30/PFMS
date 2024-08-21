package models

type Account struct {
	Model

	ProfilePicture string `json:"profile_picture" `
	DisplayName    string `json:"display_name"`
	Status         string `json:"status"`
	Username       string `json:"username"`
	FirstNameTh    string `json:"first_name_th"`
	LastNameTh     string `json:"last_name_th"`
	FirstNameEng   string `json:"first_name_eng"`
	LastNameEng    string `json:"last_name_eng"`
	NicknameTh     string `json:"nickname_th"`
	NicknameEng    string `json:"nickname_eng"`
	Password       string `json:"password"`
	Email          string `json:"email"`

	MobileNo string `json:"mobile_no"`

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	Role      string `json:"role"`
	IsSetData bool   `json:"is_set_data"`
	IsLogin   bool   `json:"is_login"`
}
