package dto

type LoginData struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type ExposedUserData struct {
	UserName string `json:"userName"`
	Id       string `json:"id"`
}

type LoginResponseData struct {
	User  ExposedUserData `json:"user"`
	Token string          `json:"token"`
}
