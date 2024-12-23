package auth

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type ValidTokenResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
	Exp int64 `json:"exp"`
}